package bni

import (
	"encoding/hex"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/config"
	"github.com/HyperService-Consortium/go-ves/lib/upstream"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"math/big"
	"testing"
)

func TestBN_Translate(t *testing.T) {
	type fields struct {
		dns    types.ChainDNSInterface
		signer uip.Signer
	}
	type args struct {
		intent  uip.TransactionIntent
		storage uip.Storage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		//want    uip.RawTransaction
		wantErr bool
		assert  upstream.GJSONAssertion
	}{
		{"test_easy", fields{
			dns:    config.ChainDNS,
			signer: nil,
		}, args{
			intent: &opintent.TransactionIntent{
				TransType: trans_type.ContractInvoke,
				Src:       sugar.HandlerError(hex.DecodeString("93334ae4b2d42ebba8cc7c797bfeb02bfb3349d6")).([]byte),
				Dst:       sugar.HandlerError(hex.DecodeString("263fef3fe76fd4075ac16271d5115d01206d3674")).([]byte),
				Meta: sugar.HandlerError(
					opintent.Serializer.Meta.Contract.Marshal(
						&opintent.ContractInvokeMeta{
							Code:     []byte("A"),
							FuncName: "updateStake",
							Params: []uip.VTok{
								(*opintent.Uint256)(big.NewInt(1001)),
							},
						})).([]byte),
				Amt:     "00",
				ChainID: 2,
			},
			storage: nil,
		}, false, upstream.GJSONWant(
			upstream.Kv{K: "method", V: "eth_sendTransaction"},
			upstream.Kv{K: "params.0.data", V: "0x7c1f751f00000000000000000000000000000000000000000000000000000000000003e9"},
			upstream.Kv{K: "params.0.from", V: "0x93334ae4b2d42ebba8cc7c797bfeb02bfb3349d6"},
			upstream.Kv{K: "params.0.to", V: "0x263fef3fe76fd4075ac16271d5115d01206d3674"},
			upstream.Kv{K: "params.0.value", V: nil},
		)},
		{"test_payment", fields{
			dns:    config.ChainDNS,
			signer: nil,
		}, args{
			intent: &opintent.TransactionIntent{
				TransType: trans_type.Payment,
				Src:       sugar.HandlerError(hex.DecodeString("ce4871f094b30ed5bed4aa19d28cf654c6e8b3f3")).([]byte),
				Dst:       sugar.HandlerError(hex.DecodeString("d977c0b967631f5bcc1f112fcb926ae53a1432c4")).([]byte),
				Meta:      nil,
				Amt:       "03e8",
				ChainID:   2,
			},
			storage: nil,
		}, false, upstream.GJSONWant(
			upstream.Kv{K: "method", V: "eth_sendTransaction"},
			upstream.Kv{K: "params.0.data", V: nil},
			upstream.Kv{K: "params.0.from", V: "0xce4871f094b30ed5bed4aa19d28cf654c6e8b3f3"},
			upstream.Kv{K: "params.0.to", V: "0xd977c0b967631f5bcc1f112fcb926ae53a1432c4"},
			upstream.Kv{K: "params.0.value", V: "0x3e8"},
		)},
		{"test_with_storage_var", fields{
			dns:    config.ChainDNS,
			signer: nil,
		}, args{
			intent: &opintent.TransactionIntent{
				TransType: trans_type.ContractInvoke,
				Src:       sugar.HandlerError(hex.DecodeString("93334ae4b2d42ebba8cc7c797bfeb02bfb3349d6")).([]byte),
				Dst:       sugar.HandlerError(hex.DecodeString("263fef3fe76fd4075ac16271d5115d01206d3674")).([]byte),
				Meta: sugar.HandlerError(
					opintent.Serializer.Meta.Contract.Marshal(
						&opintent.ContractInvokeMeta{
							FuncName: "updateStake",
							Params: []uip.VTok{
								&opintent.StateVariable{
									Type: value_type.Uint256,
									Contract: opintent.NamespacedRawAccount{
										Address: make([]byte, 20),
										ChainID: 7,
									},
									Pos:   []byte{0},
									Field: []byte("staking"),
								},
							},
						})).([]byte),
				Amt:     "00",
				ChainID: 2,
			},
			storage: upstream.MockBNIStorage{Data: []upstream.MockData{
				{
					ChainID:         7,
					TypeID:          value_type.Uint256,
					ContractAddress: make([]byte, 20),
					Pos:             []byte{0},
					Description:     []byte("staking"),
					V:               upstream.MockValue{T: value_type.Uint256, V: big.NewInt(0x0233)},
				},
			}},
		}, false, upstream.GJSONWant(
			upstream.Kv{K: "method", V: "eth_sendTransaction"},
			upstream.Kv{K: "params.0.data", V: "0x7c1f751f0000000000000000000000000000000000000000000000000000000000000233"},
			upstream.Kv{K: "params.0.from", V: "0x93334ae4b2d42ebba8cc7c797bfeb02bfb3349d6"},
			upstream.Kv{K: "params.0.to", V: "0x263fef3fe76fd4075ac16271d5115d01206d3674"},
			upstream.Kv{K: "params.0.value", V: nil},
		)},
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
			if err = tt.assert.AssertBytes(sugar.HandlerError(got.Bytes()).([]byte)); err != nil {
				t.Error(err)
			}
		})
	}
}

func Test_decoratePrefix(t *testing.T) {
	fmt.Println(decoratePrefix("041a"))
}
