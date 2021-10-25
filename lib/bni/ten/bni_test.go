package bni

import (
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/config"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
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
			intent: &opintent.TransactionIntent{
				TransType: trans_type.Payment,
				Src:       ten.GetPublicKey(),
				Dst:       ten2.GetPublicKey(),
				Meta:      nil,
				Amt:       "15",
				ChainID:   3,
			},
			storage: nil,
		}, false},
		{"test_contract", fields{dns: config.ChainDNS, signer: ten}, args{
			intent: &opintent.TransactionIntent{
				TransType: trans_type.ContractInvoke,
				Src:       ten.GetPublicKey(),
				Dst:       ten2.GetPublicKey(),
				Meta:      sugar.HandlerError(opintent.Serializer.Meta.Contract.Marshal(&lexer.ContractInvokeMeta{
					FuncName: "Vote",
					Params:   nil,
				})).([]byte),
				Amt:       "0",
				ChainID:   3,
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

			tx, err := bn.Deserialize(b)
			assert.NoError(t, err)

			assert.EqualValues(t, got.(*rawTransaction).Type,
				tx.(*rawTransaction).Type)
			assert.EqualValues(t, got.(*rawTransaction).Header.Src,
				tx.(*rawTransaction).Header.Src)
			assert.EqualValues(t, got.(*rawTransaction).Header.Dst,
				tx.(*rawTransaction).Header.Dst)
			assert.EqualValues(t, got.(*rawTransaction).Header.Nonce,
				tx.(*rawTransaction).Header.Nonce)
			assert.EqualValues(t, got.(*rawTransaction).Header.Value,
				tx.(*rawTransaction).Header.Value)
			assert.EqualValues(t, got.(*rawTransaction).Header.Signature,
				tx.(*rawTransaction).Header.Signature)
		})
	}
}
