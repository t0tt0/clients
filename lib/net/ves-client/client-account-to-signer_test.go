package vesclient

import (
	"fmt"
	ChainType "github.com/HyperService-Consortium/go-uip/const/chain_type"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"testing"
)


func TestVesClient_AccountToSigner(t *testing.T) {
	type args struct {
		account *Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"easy", createFields(), args{
			account: &Account{
				Address:   "00",
				Addition:  "00",
				ChainType: 0,
				ChainID:   1,
			},
		}, true},
		{"ethereum", createFields(), args{
			account: &Account{
				Address:   "00",
				Addition:  "00",
				ChainType: uiptypes.ChainTypeUnderlyingType(
					ChainType.Ethereum),
				ChainID:   1,
			},
		}, false},
		{"tendermint", createFields(), args{
			account: &Account{
				Address:   "00",
				Addition:  "00",
				ChainType: uiptypes.ChainTypeUnderlyingType(
					ChainType.TendermintNSB),
				ChainID:   3,
			},
		}, true},
		{"tendermint", createFields(), args{
			account: &Account{
				Address:   "0000000000000000000000000000000000000000000000000000000000000000",
				Addition:  "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
				ChainType: uiptypes.ChainTypeUnderlyingType(
					ChainType.TendermintNSB),
				ChainID:   3,
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := &VesClient{
				p:                      tt.fields.p,
				rwMutex:                tt.fields.rwMutex,
				logger:                 tt.fields.logger,
				module:                 tt.fields.module,
				closeSessionRWMutex:    tt.fields.closeSessionRWMutex,
				closeSessionSubscriber: tt.fields.closeSessionSubscriber,
				name:                   tt.fields.name,
				db:                     tt.fields.db,
				conn:                   tt.fields.conn,
				nsbSigner:              tt.fields.nsbSigner,
				dns:                    tt.fields.dns,
				nsbClient:              tt.fields.nsbClient,
				waitOpt:                tt.fields.waitOpt,
				cb:                     tt.fields.cb,
				quit:                   tt.fields.quit,
				nsbip:                  tt.fields.nsbip,
				grpcip:                 tt.fields.grpcip,
				nsbBase:                tt.fields.nsbBase,
			}
			got, err := vc.AccountToSigner(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountToSigner() error = %v, wantErr %v", describer.Describe(err), tt.wantErr)
				return
			}
			if err != nil {
				fmt.Printf("good: %v", describer.Describe(err))
				return
			}
			fmt.Println(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("AccountToSigner() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
