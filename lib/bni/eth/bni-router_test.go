package bni

import (
	"encoding/hex"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/go-ves/config"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"testing"
)

func TestBN_Route(t *testing.T) {
	type fields struct {
		dns    types.ChainDNSInterface
		signer uip.Signer
	}
	type args struct {
		intent  *uip.TransactionIntent
		storage uip.Storage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"test_easy", fields{
				dns: config.ChainDNS,
				signer: passwordSigner{
					pb: sugar.HandlerError(hex.DecodeString("0a8f483d32e20a7b17b598235489570b92f67e31")).([]byte),
					ps: "123456"},
			}, args{
				intent: &uip.TransactionIntent{
					TransType: trans_type.Payment,
					Src:       sugar.HandlerError(hex.DecodeString("0a8f483d32e20a7b17b598235489570b92f67e31")).([]byte),
					Dst:       sugar.HandlerError(hex.DecodeString("0a8f483d32e20a7b17b598235489570b92f67e31")).([]byte),
					Meta:      nil,
					Amt:       "03e8",
					ChainID:   6,
				},
				storage: nil,
			}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bn := &BN{
				dns:    tt.fields.dns,
				signer: tt.fields.signer,
			}
			got, err := bn.Route(tt.args.intent, tt.args.storage)
			if (err != nil) != tt.wantErr {
				t.Errorf("Route() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			fmt.Println(got)
			fmt.Println(hex.EncodeToString(got))
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Route() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
