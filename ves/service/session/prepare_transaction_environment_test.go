package sessionservice

import (
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/control"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
	"testing"
)

func createTranslateEnvField(options ...interface{}) *prepareTranslateEnvironment {
	t := &prepareTranslateEnvironment{
		Service: createService(options),
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
		case control.BlockChainInterfaceI:
			t.bn = o
		}
	}
	return t
}

//Type: "uint256",
//Value: marshal(map[string]interface{}{
//	"contract": "0000000000000000000000000000000000000000",
//	"pos":      "00",
//	"field":    "staking",
//}),
//	uiptypes.RawParams{
//								{
//									Type: "uint256",
//									Value: marshal(h{
//										"constant": 1001,
//									}),
//								},
//							}

func Test_prepareTranslateEnvironment_ensureValue(t *testing.T) {
	s := createService()
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
		{name: "valueTypeNotFound", env: createTranslateEnvField(s), args: args{
			param: newRawMeta(value_type.Unknown, ""),
		}, wantErr: true, wantCode: types.CodeValueTypeNotFound},
		{name: "constantOk", env: createTranslateEnvField(s), args: args{
			param: newRawMeta(value_type.Uint256, ""),
		}},
		{name: "notEnoughParamInformation", env: createTranslateEnvField(s), args: args{
			param: uiptypes.RawParams{
				Type: valueTypeToString(value_type.Uint256),
				Value: marshal(map[string]interface{}{
					"contract": "xx",
					"pos":      "yy",
				}),
			},
		}, wantErr:true, wantCode:types.CodeNotEnoughParamInformation},
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
