package bni

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-ethabi"
	"github.com/HyperService-Consortium/go-uip/const/token_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/lib/serial"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"golang.org/x/crypto/sha3"
	"reflect"
	"strings"
)

//var contractMetaEncoder = opintent.Serializer.Meta.Contract
type _contractMetaEncoder struct{}

func (*_contractMetaEncoder) Marshal(meta *opintent.ContractInvokeMeta) (_ []byte, err error) {
	var w = bytes.NewBuffer(nil)
	serial.Write(w, meta.Code, &err)
	serial.Write(w, meta.Meta, &err)
	serial.Write(w, meta.FuncName, &err)
	serial.Write(w, uint64(len(meta.Params)), &err)
	for i := range meta.Params {
		opintent.EncodeVTok(w, meta.Params[i], &err)
	}
	return w.Bytes(), err
}
func (*_contractMetaEncoder) Unmarshal(b []byte, meta *opintent.ContractInvokeMeta) (err error) {
	var r = bytes.NewReader(b)
	serial.Read(r, &meta.Code, &err)
	serial.Read(r, &meta.Meta, &err)
	serial.Read(r, &meta.FuncName, &err)
	var paramsLength uint64
	serial.Read(r, &paramsLength, &err)
	if err != nil {
		return
	}
	meta.Params = make([]uip.VTok, paramsLength)
	for i := range meta.Params {
		opintent.DecodeVTok(r, &meta.Params[i], &err)
	}
	return
}

var contractMetaEncoder *_contractMetaEncoder

func decoratePrefix(hexs string) string {
	if !strings.HasPrefix(hexs, "0x") {
		hexs = "0x" + hexs
	}
	return hexs
}

func decorateValuePrefix(hexs string) string {
	if !strings.HasPrefix(hexs, "0x") {
		hexs = "0x" + hexs
	}
	for strings.HasPrefix(hexs, "0x0") && len(hexs) > 3 {
		hexs = "0x" + hexs[3:]
	}
	return hexs
}

func convertVariableToEthVariable(variable uip.Variable) interface{} {
	return variable.GetValue()
}

func convertConstantToEthVariable(constant uip.VTok) interface{} {
	return constant.(gvm.Ref).Unwrap()
}

func convertToEthType(vt value_type.Type) (string, error) {
	switch vt {
	case value_type.Uint8:
		return "uint8", nil
	case value_type.Uint16:
		return "uint16", nil
	case value_type.Uint32:
		return "uint32", nil
	case value_type.Uint64:
		return "uint64", nil
	case value_type.Uint128:
		return "uint128", nil
	case value_type.Uint256:
		return "uint256", nil
	case value_type.Int8:
		return "int8", nil
	case value_type.Int16:
		return "int16", nil
	case value_type.Int32:
		return "int32", nil
	case value_type.Int64:
		return "int64", nil
	case value_type.Int128:
		return "int128", nil
	case value_type.Int256:
		return "int256", nil
	case value_type.Bytes:
		return "bytes", nil
	case value_type.Bool:
		return "bool", nil
	case value_type.String:
		return "string", nil
	default:
		return "", fmt.Errorf("invalid value_type: %v", vt)
	}
}

func ContractInvocationDataABI(_ uip.ChainID, meta *opintent.ContractInvokeMeta, storage uip.Storage) ([]byte, error) {

	abiEncoder := ethabi.NewEncoder()
	keccak := sha3.NewLegacyKeccak256()
	var descSlice []string
	var valSlice = make([]interface{}, 0)
	var funcSig = meta.FuncName + "("
	//var err error
	for id, param := range meta.Params {
		t, err := convertToEthType(value_type.Type(param.GetGVMType()))
		if err != nil {
			return nil, err
		}
		if t == "address payable" || t == "contract" {
			t = "address"
		}
		funcSig += t
		descSlice = append(descSlice, t)
		if param.GetGVMTok() != token_type.Constant {
			if param.GetGVMTok() != token_type.StateVariable {
				return nil, errors.New("only support token_type.{Constant,StateVariable} now")
			}

			// todo remove assertion
			param := param.(*opintent.StateVariable)

			var contract uip.Account
			var ok bool
			if contract, ok = param.Contract.(uip.Account); !ok {
				return nil, fmt.Errorf("assuming contract is uip.Account ,but got %v", reflect.TypeOf(contract))
			}
			v, err := storage.GetStorageAt(contract.GetChainId(), value_type.FromString(t), contract.GetAddress(),
				param.Pos, param.Field)
			if err != nil {
				return nil, err
			}
			valSlice = append(valSlice, convertVariableToEthVariable(v))
		} else {
			valSlice = append(valSlice, convertConstantToEthVariable(param))
		}

		if id == len(meta.Params)-1 {
			funcSig += ")"
		} else {
			funcSig += ","
		}
	}

	// fixed bug: should close parentheses
	if len(meta.Params) == 0 {
		funcSig += ")"
	}
	keccak.Write([]byte(funcSig))
	result := keccak.Sum([]byte(""))[0:4]
	if len(descSlice) > 0 {
		result2, err := abiEncoder.Encodes(descSlice, valSlice)
		if err != nil {
			return nil, err
		}
		result = append(result, result2...)
	}
	return result, nil
}
