package sessionservice

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/token_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	payment_option "github.com/HyperService-Consortium/go-ves/lib/bni/payment-option"
	"github.com/HyperService-Consortium/go-ves/lib/upstream"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/ves/control"
	"github.com/HyperService-Consortium/go-ves/ves/model"
	"github.com/tidwall/gjson"
	"reflect"
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

func (svc *prepareTranslateEnvironment) ensure() error {
	switch svc.ti.TransType {
	case trans_type.ContractInvoke:
		return svc.ensureContractInvoke()
	case trans_type.Payment:
		return svc.ensurePayment()
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

func (svc *prepareTranslateEnvironment) ensureContractInvoke() error {
	if svc.ti.Meta != nil {
		var meta opintent.ContractInvokeMeta
		err := opintent.Serializer.Meta.Contract.Unmarshal(svc.ti.Meta, &meta)
		if err != nil {
			return wrapper.Wrap(types.CodeDeserializeTransactionError, err)
		}
		for i, param := range meta.Params {

			if err = svc.ensureValue(param); err != nil {
				return wrapper.Wrap(types.CodeEnsureTransactionValueError, wrapPos{i: i, err: err})
			}
		}
	}
	return nil
}

func (svc *prepareTranslateEnvironment) ensurePayment() error {
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

// to do
//type variableImpl struct {
//	uip.Variable
//}
//
//func (v variableImpl) Unwrap() interface{} {
//	panic("implement me")
//}
//
//func (v variableImpl) Encode() ([]byte, error) {
//	switch value_type.Type(v.GetGVMType()) {
//	case value_type.Uint8:
//		return opintent.Uint8(v.Unwrap().(uint8)).Encode()
//	case value_type.Uint16:
//		return opintent.Uint16(v.Unwrap().(uint16)).Encode()
//	case value_type.Uint32:
//		return opintent.Uint32(v.Unwrap().(uint32)).Encode()
//	case value_type.Uint64:
//		return opintent.Uint64(v.Unwrap().(uint64)).Encode()
//
//	case value_type.Int8:
//		return opintent.Int8(v.Unwrap().(int8)).Encode()
//	case value_type.Int16:
//		return opintent.Int16(v.Unwrap().(int16)).Encode()
//	case value_type.Int32:
//		return opintent.Int32(v.Unwrap().(int32)).Encode()
//	case value_type.Int64:
//		return opintent.Int64(v.Unwrap().(int64)).Encode()
//
//	case value_type.Uint128:
//		return (*opintent.Uint128)(v.Unwrap().(*big.Int)).Encode()
//	case value_type.Uint256:
//		return (*opintent.Uint256)(v.Unwrap().(*big.Int)).Encode()
//	case value_type.Int128:
//		return (*opintent.Int128)(v.Unwrap().(*big.Int)).Encode()
//	case value_type.Int256:
//		return (*opintent.Int256)(v.Unwrap().(*big.Int)).Encode()
//
//	case value_type.String:
//		return opintent.String(v.Unwrap().(string)).Encode()
//	case value_type.Bytes:
//		return opintent.Bytes(v.Unwrap().([]byte)).Encode()
//	case value_type.Bool:
//		return opintent.Bool(v.Unwrap().(bool)).Encode()
//	case value_type.Unknown:
//		return opintent.Undefined.Encode()
//	}
//	panic(fmt.Errorf("unknown reference type: %v", gvm_type.ExplainGVMType(v.GetGVMType())))
//}
//
//func (v variableImpl) Marshal(w io.Writer, err *error) {
//	panic("implement me")
//}
//
//func (v variableImpl) Unmarshal(r io.Reader, i *uip.VTok, err *error) {
//	panic("implement me")
//}
//
//func (v variableImpl) GetGVMTok() token_type.Type {
//	return token_type.Constant
//}
//
//func (v variableImpl) GetGVMType() gvm.RefType {
//	return gvm.RefType(v.Variable.GetType())
//}
//
//func (v variableImpl) Eval(g *interface{}) (gvm.Ref, error) {
//	return v, nil
//}

func (svc *prepareTranslateEnvironment) ensureValue(param uip.VTok) error {

	if param.GetGVMTok() != token_type.Constant {
		if param.GetGVMTok() != token_type.StateVariable {
			return errors.New("only support token_type.{Constant,StateVariable} now")
		}

		param := param.(opintent.StateVariableI)

		var contract uip.Account
		var ok bool
		if contract, ok = param.GetContract().(uip.Account); !ok {
			return fmt.Errorf("assuming contract is uip.Account ,but got %v", reflect.TypeOf(contract))
		}
		err := svc.ensureStorage(svc.bn, contract.GetChainId(), value_type.Type(
			param.GetGVMType()), contract.GetAddress(), param.GetPos(), param.GetField())
		if err != nil {
			return err
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
