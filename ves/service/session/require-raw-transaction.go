package sessionservice

import (
	"encoding/json"
	ChainType "github.com/HyperService-Consortium/go-uip/const/chain_type"
	"github.com/Myriad-Dreamin/go-ves/lib/encoding"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"golang.org/x/net/context"

	transtype "github.com/HyperService-Consortium/go-uip/const/trans_type"
	tx "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	ethbni "github.com/Myriad-Dreamin/go-ves/lib/bni/eth"
	tenbni "github.com/Myriad-Dreamin/go-ves/lib/bni/ten"
)

func (svc *Service) getBlockChainInterface(chainID uint64) (uiptypes.BlockChainInterface, error) {
	if ci, err := svc.dns.GetChainInfo(chainID); err != nil {
		return nil, wrapper.Wrap(types.CodeChainIDNotFound, err)
	} else {
		switch ci.GetChainType() {
		case ChainType.Ethereum:
			return ethbni.NewBN(svc.dns), nil
		case ChainType.TendermintNSB:
			return tenbni.NewBN(svc.dns), nil
		default:
			return nil, wrapper.WrapCode(types.CodeChainTypeNotFound)
		}
	}
}

func (svc *Service) RequireRawTransaction(
	ctx context.Context, in *uiprpc.SessionRequireRawTransactRequest) (
	*uiprpc.SessionRequireRawTransactReply, error) {

	ses, err := svc.db.QueryGUID(encoding.EncodeBase64(in.GetSessionId()))
	if err != nil {
		return nil, wrapper.Wrap(types.CodeSessionFindError, err)
	} else if ses == nil {
		return nil, wrapper.WrapCode(types.CodeSessionNotFind)
	}

	txb, err := svc.sesFSet.FindTransaction(ses.GetGUID(), ses.UnderTransacting)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeTransactionFindError, err)
	}

	var transactionIntent tx.TransactionIntent
	err = json.Unmarshal(txb, &transactionIntent)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeDeserializeTransactionError, err)
	}

	//fmt.Println(".......")

	bn, err := svc.getBlockChainInterface(transactionIntent.ChainID)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeGetBlockChainInterfaceError, err)
	}

	if err = newPrepareTranslateEnvironment(svc, ses, &transactionIntent, bn).do(); err != nil {
		return nil, wrapper.Wrap(types.CodeTransactionPrepareTranslateError, err)
	}

	var b uiptypes.RawTransaction
	b, err = bn.Translate(&transactionIntent, svc.storageHandler)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeTransactionTranslateError, err)
	}

	x, err := b.Serialize()
	if err != nil {
		return nil, wrapper.Wrap(types.CodeTransactionRawSerializeError, err)
	}

	var dest *uiprpc_base.Account
	if transactionIntent.TransType == transtype.Payment {
		dest = &uiprpc_base.Account{
			Address: transactionIntent.Dst,
			ChainId: transactionIntent.ChainID,
		}
	} else {
		dest = &uiprpc_base.Account{
			Address: svc.respAccount.GetAddress(),
			ChainId: svc.respAccount.GetChainId(),
		}
	}

	return &uiprpc.SessionRequireRawTransactReply{
		RawTransaction: x,
		Tid:            uint64(ses.UnderTransacting),
		Src: &uiprpc_base.Account{
			Address: transactionIntent.Src,
			ChainId: transactionIntent.ChainID,
		},
		Dst: dest,
	}, nil

}
