package sessionservice

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	payment_option "github.com/Myriad-Dreamin/go-ves/lib/bni/payment-option"
	"github.com/Myriad-Dreamin/go-ves/lib/encoding"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/control"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
	"github.com/tidwall/gjson"
	"strconv"
)

type prepareTranslateEnvironment struct {
	*Service
	ses *model.Session
	ti  *opintent.TransactionIntent
	bn  control.BlockChainInterfaceI
}

func newPrepareTranslateEnvironment(service *Service, ses *model.Session,
	ti *opintent.TransactionIntent, bn uiptypes.BlockChainInterface) *prepareTranslateEnvironment {
	return &prepareTranslateEnvironment{Service: service, ses: ses, ti: ti, bn: bn}
}

func (svc *prepareTranslateEnvironment) do() error {
	switch svc.ti.TransType {
	case trans_type.ContractInvoke:
		return svc.doContractInvoke()
	case trans_type.Payment:
		return svc.doPayment()
	default:
		return wrapper.WrapCode(types.CodeTransactionTypeNotFound)
	}
}

type wrapPos struct {
	i   int
	err error
}

func (w wrapPos) Error() string {
	return fmt.Sprintf("<%d,%v>", w.i, w.err)
}

func (svc *prepareTranslateEnvironment) doContractInvoke() error {
	var meta uiptypes.ContractInvokeMeta
	err := json.Unmarshal(svc.ti.Meta, &meta)
	if err != nil {
		return err
	}
	for i, param := range meta.Params {

		if err = svc.ensureValue(param); err != nil {
			return wrapper.Wrap(types.CodeEnsureTransactionValueError, wrapPos{i: i, err: err})
		}
	}
	return nil
}

func (svc *prepareTranslateEnvironment) doPayment() error {
	n, ok, err := payment_option.NeedInconsistentValueOption(gjson.ParseBytes(svc.ti.Meta))
	if err != nil {
		return wrapper.Wrap(types.CodeParsePaymentOptionInconsistentValueError, err)
	}
	if ok {
		v, err := svc.bn.GetStorageAt(n.ChainID, n.TypeID, n.ContractAddress, n.Pos, n.Description)
		if err != nil {
			return err
		}
		svc.logger.Info("getting state from blockchain",
			"address", hex.EncodeToString(n.ContractAddress),
			"value:", v.GetValue(), "type", v.GetType(), "at pos", hex.EncodeToString(n.Pos))
		err = svc.storageHandler.SetStorageOf(n.ChainID, n.TypeID, n.ContractAddress, n.Pos, n.Description, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *prepareTranslateEnvironment) ensureValue(param uiptypes.RawParams) error {
	var intDesc uiptypes.TypeID
	if intDesc = value_type.FromString(param.Type); intDesc == value_type.Unknown {
		return wrapper.WrapString(types.CodeValueTypeNotFound, strconv.Itoa(int(intDesc)))
	}

	result := gjson.ParseBytes(param.Value)
	if !result.Get("constant").Exists() {
		if result.Get("contract").Exists() &&
			result.Get("pos").Exists() &&
			result.Get("field").Exists() {
			ca, err := encoding.DecodeHex(result.Get("contract").String())
			if err != nil {
				return err
			}
			pos, err := encoding.DecodeHex(result.Get("pos").String())
			if err != nil {
				return err
			}
			desc := []byte(result.Get("field").String())

			v, err := svc.bn.GetStorageAt(svc.ti.ChainID, intDesc, ca, pos, desc)
			if err != nil {
				return err
			}
			vv, err := json.Marshal(v)
			if err != nil {
				return err
			}
			err = svc.storage.SetKV(svc.ses.GetGUID(), desc, vv)
			if err != nil {
				return err
			}
		} else {
			return wrapper.WrapCode(types.CodeNotEnoughParamInformation)
		}
	}

	return nil
}
