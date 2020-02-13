package vesclient

import (
	"bytes"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"testing"
	"time"
)

const (
	nsbBaseKey = "ten1"
)

var (
	accountNSBBasePrivateKey1 = make([]byte, 64)
	vcWithDefaultSigner       = createFields(
		sugar.HandlerError(signaturer.NewTendermintNSBSigner(
			make([]byte, 64))))
	vcWithAccountMockData = createFields(
		withNSBBase(nsbBaseKey),
		mockAccountDB(
			AccountQueryAliasMockData{
				K: nsbBaseKey, V: &Account{
					ID:        0,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					Alias:     nsbBaseKey,
					Address:   encodeAddress(make([]byte, 32)),
					Addition:  encodeAddition(accountNSBBasePrivateKey1),
					ChainType: 0,
					ChainID:   0,
				},
			},
		),
	)
)

func TestVesClient_getNSBSigner(t *testing.T) {
	tests := []struct {
		name      string
		fields    fields
		wantBytes []byte
		wantErr   bool
	}{
		{"withDefault",
			vcWithDefaultSigner,
			vcWithDefaultSigner.nsbSigner.GetPublicKey(),
			false},
		{"withQueryAlias",
			vcWithAccountMockData,
			sugar.HandlerError(
				signaturer.NewTendermintNSBSigner(
					accountNSBBasePrivateKey1)).(uiptypes.Signer).GetPublicKey(),
			false},
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
			got, err := vc.getNSBSigner()
			if (err != nil) != tt.wantErr {
				t.Errorf("getNSBSigner() error = %v, wantErr %v", describer.Describe(err), tt.wantErr)
				return
			}
			if err != nil {
				fmt.Println("good:", describer.Describe(err))
				return
			}
			if !bytes.Equal(got.GetPublicKey(), tt.wantBytes) {
				t.Errorf("getNSBSigner() got = %v, want %v", got.GetPublicKey(), tt.wantBytes)
			}
		})
	}
}
