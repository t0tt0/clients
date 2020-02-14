package vesclient

import (
	"bytes"
	"fmt"
	ChainType "github.com/HyperService-Consortium/go-uip/const/chain_type"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"testing"
	"time"
)

var (
	accountInvertFindPublicKey1 = func() []byte {
		b := make([]byte, 32)
		b[0] = 1
		b[1] = 233
		return b
	}()
	vcWithInvertFindMockData = createFields(
		withNSBBase(nsbBaseKey),
		mockAccountDB(
			AccountInvertFindMockData{
				K: &accountKey{
					ChainId: 1,
					Address: accountInvertFindPublicKey1},
				V: &Account{
					ID:        0,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					Alias:     nsbBaseKey,
					Address:   encodeAddress(accountInvertFindPublicKey1),
					Addition:  "",
					ChainType: uiptypes.ChainTypeUnderlyingType(ChainType.Ethereum),
					ChainID:   1,
				},
			},
		),
	)
)

func TestVesClient_getRespSigner(t *testing.T) {
	type args struct {
		acc uiptypes.Account
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantBytes []byte
		wantErr   bool
	}{
		{"base", vcWithInvertFindMockData, args{
			&accountKey{
				ChainId: 1,
				Address: accountInvertFindPublicKey1,
			},
		}, accountInvertFindPublicKey1, false},
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
				quit:                   tt.fields.quit,
				nsbHost:                tt.fields.nsbip,
				vesHost:                tt.fields.grpcip,
				nsbBase:                tt.fields.nsbBase,
			}
			got, err := vc.getRespSigner(tt.args.acc)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRespSigner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				fmt.Println("good:", err)
				return
			}
			fmt.Println(got.GetPublicKey())
			if !bytes.Equal(got.GetPublicKey(), tt.wantBytes) {
				t.Errorf("getRespSigner() got = %v, want %v", got.GetPublicKey(), tt.wantBytes)
			}
		})
	}
}
