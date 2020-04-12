package bni

import (
	"encoding/hex"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/config"
	dep_uip "github.com/HyperService-Consortium/go-ves/dependency/uip"
	"github.com/HyperService-Consortium/go-ves/types"
	"golang.org/x/crypto/ed25519"
	"testing"
)

var _ uip.BlockChainInterface = new(BN)

func TestBN_Translate(t *testing.T) {
	type fields struct {
		dns    types.ChainDNSInterface
		signer uip.Signer
	}
	type args struct {
		intent  uip.TransactionIntent
		storage uip.Storage
	}

	var ten, err = signaturer.NewTendermintNSBSigner(ed25519.NewKeyFromSeed(append(make([]byte, 31), 2)))
	if err != nil {
		t.Errorf("Translate() error = %v", err)
		return
	}

	ten2, err := signaturer.NewTendermintNSBSigner(ed25519.NewKeyFromSeed(append(make([]byte, 31), 13)))
	if err != nil {
		t.Errorf("Translate() error = %v", err)
		return
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		//want    uip.RawTransaction
		wantErr bool
	}{
		{"test_easy", fields{dns: config.ChainDNS, signer: ten}, args{
			intent: &dep_uip.TransactionIntent{
				TransactionIntent: &opintent.TransactionIntent{
					TransType: trans_type.Payment,
					Src:       ten.GetPublicKey(),
					Dst:       ten2.GetPublicKey(),
					Meta:      nil,
					Amt:       "15",
					ChainID:   3,
				},
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
			got, err := bn.Translate(tt.args.intent, tt.args.storage)
			if (err != nil) != tt.wantErr {
				t.Errorf("Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			b, err := got.Serialize()
			if err != nil {
				t.Errorf("Translate() error = %v", err)
				return
			}
			fmt.Println(hex.EncodeToString(b))
		})
	}
}
