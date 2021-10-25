package sessionservice

import (
	"errors"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/lib/basic/encoding"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/ves/mock"
	"github.com/HyperService-Consortium/go-ves/ves/model"
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

//type badParam struct {
//
//}
//
//func (b badParam) Unwrap() interface{} {
//	return opintent.Undefined
//}
//
//func (b badParam) Encode() ([]byte, error) {
//	panic("implement me")
//}
//
//func (b badParam) Marshal(w io.Writer, err *error) {
//	panic("implement me")
//}
//func (b badParam) Unmarshal(r io.Reader, i *uip.VTok, err *error) {
//	panic("implement me")
//}
//
//func (b badParam) GetGVMTok() gvm.TokType {
//	return token_type.Constant
//}
//
//func (b badParam) GetGVMType() gvm.RefType {
//	return gvm.RefType(value_type.Length+1)
//}
//
//func (b badParam) Eval(_ *gvm.ExecCtx) (gvm.Ref, error) {
//	return b, nil
//}

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
		[]byte("goodButErrorType")).Return(uip.VariableImpl{
		Type: value_type.Uint128, Value: nil}, nil)

	bn.EXPECT().GetStorageAt(
		ti.ChainID, value_type.Uint256,
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		sugar.HandlerError(encoding.DecodeHex("00")).([]byte),
		[]byte("goodButErrorType2")).Return(uip.VariableImpl{
		Type: value_type.Uint128, Value: big.NewInt(1)}, nil)

	v := uip.VariableImpl{
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
		param uip.VTok
	}
	tests := []struct {
		name     string
		env      *prepareTranslateEnvironment
		args     args
		wantErr  bool
		wantCode int
	}{
		//{name: "valueTypeNotFound", env: createTranslateEnvField(), args: args{
		//	param: badParam{},
		//}, wantErr: true, wantCode: types.CodeValueTypeNotFound},
		{name: "undefinedOk", env: createTranslateEnvField(), args: args{
			param: opintent.Undefined,
		}, wantErr: false},
		{name: "constantOk", env: createTranslateEnvField(), args: args{
			param: (*opintent.Uint256)(big.NewInt(1)),
		}},
		//{name: "notEnoughParamInformation", env: createTranslateEnvField(), args: args{
		//	param: uip.RawParam{
		//		Type: valueTypeToString(value_type.Uint256),
		//		Value: marshal(map[string]interface{}{
		//			"contract": "xx",
		//			"pos":      "yy",
		//		}),
		//	},
		//}, wantErr: true, wantCode: types.CodeNotEnoughParamInformation},
		//{name: "badContractField", env: createTranslateEnvField(), args: args{
		//	param: newVarRawMeta(value_type.Uint256, "xx", "00", "bad"),
		//}, wantErr: true, wantCode: types.CodeBadContractField},
		//{name: "badPosField", env: createTranslateEnvField(), args: args{
		//	param: newVarRawMeta(value_type.Uint256, "00", "yy", "bad"),
		//}, wantErr: true, wantCode: types.CodeBadPosField},
		{name: "getStorageError", env: createTranslateEnvField(
			ti, bn,
		), args: args{
			param: newVarRawMeta(value_type.Uint256, ti.ChainID, []byte{0}, []byte{0}, []byte("bad")),
		}, wantErr: true, wantCode: types.CodeGetStorageError},
		{name: "getStorageTypeError", env: createTranslateEnvField(
			ti, bn,
		), args: args{
			param: newVarRawMeta(value_type.Uint256, ti.ChainID, []byte{0}, []byte{0}, []byte("goodButErrorType")),
		}, wantErr: true, wantCode: types.CodeGetStorageTypeError},
		{name: "getStorageTypeError", env: createTranslateEnvField(
			ti, bn,
		), args: args{
			param: newVarRawMeta(value_type.Uint256, ti.ChainID, []byte{0}, []byte{0}, []byte("goodButErrorType2")),
		}, wantErr: true, wantCode: types.CodeGetStorageTypeError},
		{name: "setStorageError", env: createTranslateEnvField(
			ti, bn, storageHandler,
		), args: args{
			param: newVarRawMeta(value_type.Uint256, ti.ChainID, []byte{0}, []byte{0}, []byte("goodButSetError")),
		}, wantErr: true, wantCode: types.CodeSetStorageError},
		{name: "ok", env: createTranslateEnvField(
			ti, bn, storageHandler,
		), args: args{
			param: newVarRawMeta(value_type.Uint256, ti.ChainID, []byte{0}, []byte{0}, []byte("good")),
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

func Test_prepareTranslateEnvironment_ensure(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	_, ti, _ := dataGoodTransactionIntent(t)
	ti0 := *ti
	ti0.TransType = trans_type.ContractInvoke + 23333
	ti1 := *ti
	ti1.TransType = trans_type.ContractInvoke
	ti1.Meta = sugar.HandlerError(
		opintent.Serializer.Meta.Contract.Marshal(
			&opintent.ContractInvokeMeta{
				FuncName: "updateStake",
				Params: []uip.VTok{
					(*opintent.Uint256)(big.NewInt(1001)),
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
			if err := tt.env.ensure(); (err != nil) != tt.wantErr {
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
	ti0.Meta = sugar.HandlerError(opintent.Serializer.Meta.Contract.Marshal(
		&opintent.ContractInvokeMeta{
			FuncName: "updateStake",
			Params: []uip.VTok{
				opintent.LocalStateVariable{},
			},
		})).([]byte)
	ti1.Meta = sugar.HandlerError(opintent.Serializer.Meta.Contract.Marshal(
		&opintent.ContractInvokeMeta{
			FuncName: "updateStake",
			Params: []uip.VTok{
				(*opintent.Uint256)(big.NewInt(1001)),
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
			if err := tt.env.ensureContractInvoke(); (err != nil) != tt.wantErr {
				t.Errorf("ensureContractInvoke() error = %v, wantErr %v", err, tt.wantErr)
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
			if err := tt.env.ensurePayment(); (err != nil) != tt.wantErr {
				t.Errorf("ensurePayment() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				checkErrorCode(t, err, tt.wantCode)
				return
			}
		})
	}
}
