package sessionservice

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/Myriad-Dreamin/go-ves/config"
	payment_option "github.com/Myriad-Dreamin/go-ves/lib/bni/payment-option"
	"github.com/Myriad-Dreamin/go-ves/lib/encoding"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/tidwall/gjson"

	"golang.org/x/net/context"

	transtype "github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	tx "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	uipbase "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	ethbni "github.com/Myriad-Dreamin/go-ves/lib/bni/eth"
	tenbni "github.com/Myriad-Dreamin/go-ves/lib/bni/ten"
)

var bnis = map[uint64]uiptypes.BlockChainInterface{
	1: ethbni.NewBN(config.ChainDNS),
	2: ethbni.NewBN(config.ChainDNS),
	3: tenbni.NewBN(config.ChainDNS),
	4: tenbni.NewBN(config.ChainDNS),
}

func (svc *Service) RequireRawTransaction(
	ctx context.Context, in *uiprpc.SessionRequireRawTransactRequest) (
	*uiprpc.SessionRequireRawTransactReply, error) {
	// todo errors.New("TODO")

	ses, err := svc.db.QueryGUID(encoding.EncodeBase64(in.GetSessionId()))
	if err != nil {
		return nil, wrapper.Wrap(types.CodeSessionFindError, err)
	} else if ses == nil {
		return nil, wrapper.WrapCode(types.CodeSessionNotFindError)
	}

	var transactionIntent tx.TransactionIntent
	txb, err := svc.sesFSet.FindTransaction(ses.GetGUID(), ses.UnderTransacting)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeTransactionFindError, err)
	}
	err = json.Unmarshal(txb, &transactionIntent)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeDeserializeTransactionError, err)
	}

	//fmt.Println(".......")

	bn := bnis[transactionIntent.ChainID]

	if transactionIntent.TransType == transtype.ContractInvoke {
		var meta uiptypes.ContractInvokeMeta

		err := json.Unmarshal(transactionIntent.Meta, &meta)
		if err != nil {
			return nil, err
		}

		var intDesc uiptypes.TypeID
		for _, param := range meta.Params {
			if intDesc = value_type.FromString(param.Type); intDesc == value_type.Unknown {
				return nil, errors.New("unknown type: " + param.Type)
			}

			result := gjson.ParseBytes(param.Value)
			if !result.Get("constant").Exists() {
				if result.Get("contract").Exists() &&
					result.Get("pos").Exists() &&
					result.Get("field").Exists() {
					ca, err := hex.DecodeString(result.Get("contract").String())
					if err != nil {
						return nil, err
					}
					pos, err := hex.DecodeString(result.Get("pos").String())
					if err != nil {
						return nil, err
					}
					desc := []byte(result.Get("field").String())

					v, err := bn.GetStorageAt(transactionIntent.ChainID, intDesc, ca, pos, desc)
					if err != nil {
						return nil, err
					}
					vv, err := json.Marshal(v)
					if err != nil {
						return nil, err
					}
					err = svc.storage.SetKV(ses.GetGUID(), desc, vv)
					if err != nil {
						return nil, err
					}
				} else {
					return nil, errors.New("no enough info of source description")
				}
			}
		}
	} else {
		n, ok, err := payment_option.NeedInconsistentValueOption(gjson.ParseBytes(transactionIntent.Meta))
		if err != nil {
			svc.logger.Warn("omit need inc-val option")
			return nil, err
		}
		if ok {
			v, err := bnis[2].GetStorageAt(n.ChainID, n.TypeID, n.ContractAddress, n.Pos, n.Description)
			if err != nil {
				svc.logger.Error("get failed")
				return nil, err
			}
			svc.logger.Info("getting state from blockchain", "address", hex.EncodeToString(n.ContractAddress), "value:", v.GetValue(), "type", v.GetType(), "at pos", hex.EncodeToString(n.Pos))
			err = svc.storageHandler.SetStorageOf(n.ChainID, n.TypeID, n.ContractAddress, n.Pos, n.Description, v)
			if err != nil {
				svc.logger.Error("set failed")
				return nil, err
			}
		}
	}

	var b uiptypes.RawTransaction
	b, err = bn.Translate(&transactionIntent, svc.storageHandler)
	if err != nil {
		svc.logger.Error("translate error", "err", err)
		return nil, err
	}

	if transactionIntent.TransType == transtype.Payment {

		svc.logger.Info("return r-tx", "tid", ses.UnderTransacting, "src", transactionIntent.Src, "dst", transactionIntent.Dst)
		x, err := b.Serialize()
		if err != nil {
			return nil, err
		}

		return &uiprpc.SessionRequireRawTransactReply{
			RawTransaction: x,
			Tid:            uint64(ses.UnderTransacting),
			Src: &uipbase.Account{
				Address: transactionIntent.Src,
				ChainId: transactionIntent.ChainID,
			},
			Dst: &uipbase.Account{
				Address: transactionIntent.Dst,
				ChainId: transactionIntent.ChainID,
			},
		}, nil
	} else {
		x, err := b.Serialize()
		if err != nil {
			return nil, err
		}

		svc.logger.Info("return r-tx", "tid", ses.UnderTransacting, "src", transactionIntent.Src, "dst", svc.respAccount.GetAddress())
		return &uiprpc.SessionRequireRawTransactReply{
			RawTransaction: x,
			Tid:            uint64(ses.UnderTransacting),
			Src: &uipbase.Account{
				Address: transactionIntent.Src,
				ChainId: transactionIntent.ChainID,
			},
			Dst: &uipbase.Account{
				Address: svc.respAccount.GetAddress(),
				ChainId: svc.respAccount.GetChainId(),
			},
		}, nil
	}

}
