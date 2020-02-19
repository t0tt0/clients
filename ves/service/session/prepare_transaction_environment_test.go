package sessionservice

import (
	"errors"
	base_variable "github.com/HyperService-Consortium/go-uip/base-variable"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/lib/encoding"
	"github.com/Myriad-Dreamin/go-ves/lib/upstream"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/mock"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/golang/mock/gomock"
	"math/big"
	"testing"
)

func createTranslateEnvField(options ...interface{}) *prepareTranslateEnvironment {
	t := &prepareTranslateEnvironment{
		Service: createService(options...),
		ses:     nil,
		ti:      nil,
		bn:      nil,
	}
	for i := range options {
		switch o := options[i].(type) {
		case *Service:
			t.Service = o
		case *model.Session:
			t.ses = o
		case *opintent.TransactionIntent:
			t.ti = o
		case *mock.BlockChainInterface:
			t.bn = o
		}
	}
	return t
}

func Test_prepareTranslateEnvironment_ensureValue(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	bn := MockBlockChainInterface(ctl)
	storageHandler := MockStorageHandler(ctl)

	_, ti, _ := dataGoodTransactionIntent(t)

	bn.EXPECT().GetStorageAt(
		ti.ChainID, value_type.Uint256,
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		[]byte("bad")).Return(nil, errors.New("get error"))

	bn.EXPECT().GetStorageAt(
		ti.ChainID, value_type.Uint256,
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		[]byte("goodButErrorType")).Return(base_variable.Variable{
		Type: value_type.Uint128, Value: nil}, nil)

	bn.EXPECT().GetStorageAt(
		ti.ChainID, value_type.Uint256,
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		[]byte("goodButErrorType2")).Return(base_variable.Variable{
		Type: value_type.Uint128, Value: 1}, nil)

	v := base_variable.Variable{
		Type: value_type.Uint256, Value: big.NewInt(1)}
	bn.EXPECT().GetStorageAt(
		ti.ChainID, value_type.Uint256,
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		[]byte("goodButSetError")).Return(v, nil)
	storageHandler.EXPECT().SetStorageOf(
		ti.ChainID, value_type.Uint256,
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		[]byte("goodButSetError"), v).Return(errors.New("set error"))

	bn.EXPECT().GetStorageAt(
		ti.ChainID, value_type.Uint256,
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		[]byte("good")).Return(v, nil)
	storageHandler.EXPECT().SetStorageOf(
		ti.ChainID, value_type.Uint256,
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		[]byte("good"), v).Return(nil)

	type args struct {
		param uiptypes.RawParams
	}
	tests := []struct {
		name     string
		env      *prepareTranslateEnvironment
		args     args
		wantErr  bool
		wantCode int
	}{
		{name: "valueTypeNotFound", env: createTranslateEnvField(), args: args{
			param: newRawMeta(value_type.Unknown, ""),
		}, wantErr: true, wantCode: types.CodeValueTypeNotFound},
		{name: "constantOk", env: createTranslateEnvField(), args: args{
			param: newRawMeta(value_type.Uint256, ""),
		}},
		{name: "notEnoughParamInformation", env: createTranslateEnvField(), args: args{
			param: uiptypes.RawParams{
				Type: valueTypeToString(value_type.Uint256),
				Value: marshal(map[string]interface{}{
					"contract": "xx",
					"pos":      "yy",
				}),
			},
		}, wantErr: true, wantCode: types.CodeNotEnoughParamInformation},
		{name: "badContractField", env: createTranslateEnvField(), args: args{
			param: newVarRawMeta(value_type.Uint256, "xx", "00", "bad"),
		}, wantErr: true, wantCode: types.CodeBadContractField},
		{name: "badPosField", env: createTranslateEnvField(), args: args{
			param: newVarRawMeta(value_type.Uint256, "00", "yy", "bad"),
		}, wantErr: true, wantCode: types.CodeBadPosField},
		{name: "getStorageError", env: createTranslateEnvField(
			ti, bn,
		), args: args{
			param: newVarRawMeta(value_type.Uint256, "00", "00", "bad"),
		}, wantErr: true, wantCode: types.CodeGetStorageError},
		{name: "getStorageTypeError", env: createTranslateEnvField(
			ti, bn,
		), args: args{
			param: newVarRawMeta(value_type.Uint256, "00", "00", "goodButErrorType"),
		}, wantErr: true, wantCode: types.CodeGetStorageTypeError},
		{name: "getStorageTypeError", env: createTranslateEnvField(
			ti, bn,
		), args: args{
			param: newVarRawMeta(value_type.Uint256, "00", "00", "goodButErrorType2"),
		}, wantErr: true, wantCode: types.CodeGetStorageTypeError},
		{name: "setStorageError", env: createTranslateEnvField(
			ti, bn, storageHandler,
		), args: args{
			param: newVarRawMeta(value_type.Uint256, "00", "00", "goodButSetError"),
		}, wantErr: true, wantCode: types.CodeSetStorageError},
		{name: "ok", env: createTranslateEnvField(
			ti, bn, storageHandler,
		), args: args{
			param: newVarRawMeta(value_type.Uint256, "00", "00", "good"),
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.env.ensureValue(tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("ensureValue() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				checkErrorCode(t, err, tt.wantCode)
				return
			}
		})
	}
}

func Test_prepareTranslateEnvironment_do(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	_, ti, _ := dataGoodTransactionIntent(t)
	ti0 := *ti
	ti0.TransType = trans_type.ContractInvoke + 23333
	ti1 := *ti
	ti1.TransType = trans_type.ContractInvoke
	ti1.Meta = sugar.HandlerError(upstream.Serializer.Meta.Contract.Marshal(
		&uiptypes.ContractInvokeMeta{
			FuncName: "updateStake",
			Params: []uiptypes.RawParams{
				{
					Type: "uint256",
					Value: marshal(map[string]interface{}{
						"constant": 1001,
					}),
				},
			},
		})).([]byte)
	ti2 := *ti
	ti2.TransType = trans_type.Payment

	tests := []struct {
		name     string
		env      *prepareTranslateEnvironment
		wantErr  bool
		wantCode int
	}{
		{name: "transactionTypeNotFound", env: createTranslateEnvField(
			&ti0,
		), wantErr: true, wantCode: types.CodeTransactionTypeNotFound},
		{name: "okContractInvoke", env: createTranslateEnvField(
			&ti1,
		), wantErr: false},
		{name: "okPayment", env: createTranslateEnvField(
			&ti2,
		), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.env.do(); (err != nil) != tt.wantErr {
				t.Errorf("do() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				checkErrorCode(t, err, tt.wantCode)
				return
			}
		})
	}
}

func Test_prepareTranslateEnvironment_doContractInvoke(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	_, ti, _ := dataGoodTransactionIntent(t)
	ti1, ti0 := *ti, *ti

	ti.TransType = trans_type.ContractInvoke
	ti0.TransType = trans_type.ContractInvoke
	ti1.TransType = trans_type.ContractInvoke

	ti.Meta = nil
	ti0.Meta = sugar.HandlerError(upstream.Serializer.Meta.Contract.Marshal(
		&uiptypes.ContractInvokeMeta{
			FuncName: "updateStake",
			Params: []uiptypes.RawParams{
				{
					Type:  "uint256",
					Value: nil,
				},
			},
		})).([]byte)
	ti1.Meta = sugar.HandlerError(upstream.Serializer.Meta.Contract.Marshal(
		&uiptypes.ContractInvokeMeta{
			FuncName: "updateStake",
			Params: []uiptypes.RawParams{
				{
					Type: "uint256",
					Value: marshal(map[string]interface{}{
						"constant": 1001,
					}),
				},
			},
		})).([]byte)

	tests := []struct {
		name     string
		env      *prepareTranslateEnvironment
		wantErr  bool
		wantCode int
	}{
		{name: "DeserializeTransactionError", env: createTranslateEnvField(
			ti,
		), wantErr: true, wantCode: types.CodeDeserializeTransactionError},
		{name: "EnsureTransactionValueError", env: createTranslateEnvField(
			&ti0,
		), wantErr: true, wantCode: types.CodeEnsureTransactionValueError},
		{name: "ok", env: createTranslateEnvField(
			&ti1,
		), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.env.doContractInvoke(); (err != nil) != tt.wantErr {
				t.Errorf("doContractInvoke() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				checkErrorCode(t, err, tt.wantCode)
				return
			}
		})
	}
}

func Test_prepareTranslateEnvironment_doPayment(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	_, ti, _ := dataGoodTransactionIntent(t)
	ti.TransType = trans_type.Payment
	//todo test option
	tests := []struct {
		name     string
		env      *prepareTranslateEnvironment
		wantErr  bool
		wantCode int
	}{
		{name: "ok", env: createTranslateEnvField(
			ti,
		), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.env.doPayment(); (err != nil) != tt.wantErr {
				t.Errorf("doPayment() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				checkErrorCode(t, err, tt.wantCode)
				return
			}
		})
	}
}
