package sessionservice

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/go-ves/lib/backend/wrapper"
	"github.com/Myriad-Dreamin/go-ves/lib/basic/encoding"
	payment_option "github.com/Myriad-Dreamin/go-ves/lib/bni/payment-option"
	"github.com/Myriad-Dreamin/go-ves/lib/upstream"
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
	ti *opintent.TransactionIntent, bn uip.BlockChainInterface) *prepareTranslateEnvironment {
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
	var meta uip.ContractInvokeMeta
	err := json.Unmarshal(svc.ti.Meta, &meta)
	if err != nil {
		return wrapper.Wrap(types.CodeDeserializeTransactionError, err)
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
		if err = svc.ensureStorage(
			svc.bn, n.ChainID, n.TypeID, n.ContractAddress, n.Pos, n.Description); err != nil {
			return err
		}
	}
	return nil
}

func (svc *prepareTranslateEnvironment) ensureValue(param uip.RawParam) error {
	var intDesc uip.TypeID
	if intDesc = value_type.FromString(param.Type); intDesc == value_type.Unknown {
		return wrapper.WrapString(types.CodeValueTypeNotFound, strconv.Itoa(int(intDesc)))
	}

	result := gjson.ParseBytes(param.Value)
	if !result.Get("constant").Exists() {
		if result.Get("contract").Exists() &&
			result.Get("pos").Exists() &&
			result.Get("field").Exists() {
			// todo move to uip
			ca, err := encoding.DecodeHex(result.Get("contract").String())
			if err != nil {
				return wrapper.Wrap(types.CodeBadContractField, err)
			}
			pos, err := encoding.DecodeHex(result.Get("pos").String())
			if err != nil {
				return wrapper.Wrap(types.CodeBadPosField, err)
			}
			desc := []byte(result.Get("field").String())
			if err = svc.ensureStorage(
				svc.bn, svc.ti.ChainID, intDesc, ca, pos, desc); err != nil {
				return err
			}
		} else {
			return wrapper.WrapCode(types.CodeNotEnoughParamInformation)
		}
	}

	return nil
}

func (svc *Service) ensureStorage(
	// todo: uip-types.Storage
	source control.BlockChainInterfaceI,
	chainID uip.ChainIDUnderlyingType, typeID uip.TypeID,
	contractAddress []byte, pos []byte, description []byte) error {

	v, err := source.GetStorageAt(chainID, typeID, contractAddress, pos, description)
	if err != nil {
		return wrapper.Wrap(types.CodeGetStorageError, err)
	}

	if v.GetType() != typeID {
		return wrapper.WrapString(types.CodeGetStorageTypeError, "type not expected")
	}

	if upstream.StorageValueIsValid(v.GetValue(), v.GetType()) == false {
		return wrapper.WrapString(types.CodeGetStorageTypeError, "type not expected")
	}

	svc.logger.Info("getting state from blockchain",
		"address", hex.EncodeToString(contractAddress),
		"value:", v.GetValue(), "type", v.GetType(), "at pos", hex.EncodeToString(pos))

	err = svc.storageHandler.SetStorageOf(chainID, typeID, contractAddress, pos, description, v)
	if err != nil {
		return wrapper.Wrap(types.CodeSetStorageError, err)
	}
	return nil
}
