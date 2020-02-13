package vesclient

import (
	"fmt"
	ChainType "github.com/HyperService-Consortium/go-uip/const/chain_type"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	ethbni "github.com/Myriad-Dreamin/go-ves/lib/bni/eth"
	tenbni "github.com/Myriad-Dreamin/go-ves/lib/bni/ten"
	"reflect"
	"testing"
)

const (
	ethereumChainID = iota
	tendermintChainID
	unknownChainTypeID
	unknownChainID
)

var (
	vcWithMockChainDNS = createFields(mockChainDNS(
		ChainDNSMockData{
			K: ethereumChainID,
			V: ChainInfo{
				ChainType: ChainType.Ethereum,
			},
		},
		ChainDNSMockData{
			K: tendermintChainID,
			V: ChainInfo{
				ChainType: ChainType.TendermintNSB,
			},
		},
		ChainDNSMockData{
			K: unknownChainTypeID,
			V: ChainInfo{
				ChainType: ChainType.Unassigned,
			},
		},
	))
	ethBNPType = reflect.TypeOf(new(ethbni.BN))
	tenBNPType = reflect.TypeOf(new(tenbni.BN))
)

func TestVesClient_getRouter(t *testing.T) {

	type args struct {
		chainID uint64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantType reflect.Type
		wantErr  bool
	}{
		{"getEthereumChainRouter", vcWithMockChainDNS, args{
			chainID: ethereumChainID,
		}, ethBNPType, false},
		{"getTendermintChainRouter", vcWithMockChainDNS, args{
			chainID: tendermintChainID,
		}, tenBNPType, false},
		{"chainIDNotFound", vcWithMockChainDNS, args{
			chainID: unknownChainID,
		}, nil, true},
		{"chainTypeNotFound", vcWithMockChainDNS, args{
			chainID: unknownChainTypeID,
		}, nil, true},
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
				nsbip:                  tt.fields.nsbip,
				grpcip:                 tt.fields.grpcip,
				nsbBase:                tt.fields.nsbBase,
			}
			got, err := vc.getRouter(tt.args.chainID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				fmt.Printf("good: %v", describer.Describe(err))
				return
			}
			if !reflect.DeepEqual(reflect.TypeOf(got), tt.wantType) {
				t.Errorf("getRouter() got = %v, want %v", reflect.TypeOf(got), tt.wantType)
			}
		})
	}
}

//BlockStorage

func TestVesClient_getBlockStorage(t *testing.T) {
	type args struct {
		chainID uint64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantType reflect.Type
		wantErr  bool
	}{
		{"getEthereumChainBlockStorage", vcWithMockChainDNS, args{
			chainID: ethereumChainID,
		}, ethBNPType, false},
		{"getTendermintChainBlockStorage", vcWithMockChainDNS, args{
			chainID: tendermintChainID,
		}, tenBNPType, false},
		{"chainIDNotFound", vcWithMockChainDNS, args{
			chainID: unknownChainID,
		}, nil, true},
		{"chainTypeNotFound", vcWithMockChainDNS, args{
			chainID: unknownChainTypeID,
		}, nil, true},
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
				nsbip:                  tt.fields.nsbip,
				grpcip:                 tt.fields.grpcip,
				nsbBase:                tt.fields.nsbBase,
			}
			got, err := vc.getBlockStorage(tt.args.chainID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBlockStorage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				fmt.Printf("good: %v", describer.Describe(err))
				return
			}
			if !reflect.DeepEqual(reflect.TypeOf(got), tt.wantType) {
				t.Errorf("getBlockStorage() got = %v, want %v", reflect.TypeOf(got), tt.wantType)
			}
		})
	}
}

func TestVesClient_ensureRouter(t *testing.T) {
	type args struct {
		chainID uint64
		router  uiptypes.Router
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantType reflect.Type
		want     bool
	}{
		{"ensureEthereumChainRouter", vcWithMockChainDNS, args{
			chainID: ethereumChainID,
		}, ethBNPType, true},
		{"ensureTendermintChainRouter", vcWithMockChainDNS, args{
			chainID: tendermintChainID,
		}, tenBNPType, true},
		{"chainIDNotFound", vcWithMockChainDNS, args{
			chainID: unknownChainID,
		}, nil, false},
		{"chainTypeNotFound", vcWithMockChainDNS, args{
			chainID: unknownChainTypeID,
		}, nil, false},
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
				nsbip:                  tt.fields.nsbip,
				grpcip:                 tt.fields.grpcip,
				nsbBase:                tt.fields.nsbBase,
			}
			if got := vc.ensureRouter(tt.args.chainID, &tt.args.router); got != tt.want {
				t.Errorf("ensureRouter() = %v, want %v", got, tt.want)
			} else if got {
				if !reflect.DeepEqual(reflect.TypeOf(tt.args.router), tt.wantType) {
					t.Errorf("getRouter() got = %v, want %v", reflect.TypeOf(tt.args.router), tt.wantType)
				}
			}
		})
	}
}

func TestVesClient_ensureBlockStorage(t *testing.T) {
	type args struct {
		chainID uint64
		router  uiptypes.Storage
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantType reflect.Type
		want     bool
	}{
		{"ensureEthereumChainBlockStorage", vcWithMockChainDNS, args{
			chainID: ethereumChainID,
		}, ethBNPType, true},
		{"ensureTendermintChainBlockStorage", vcWithMockChainDNS, args{
			chainID: tendermintChainID,
		}, tenBNPType, true},
		{"chainIDNotFound", vcWithMockChainDNS, args{
			chainID: unknownChainID,
		}, nil, false},
		{"chainTypeNotFound", vcWithMockChainDNS, args{
			chainID: unknownChainTypeID,
		}, nil, false},
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
				nsbip:                  tt.fields.nsbip,
				grpcip:                 tt.fields.grpcip,
				nsbBase:                tt.fields.nsbBase,
			}
			if got := vc.ensureBlockStorage(tt.args.chainID, &tt.args.router); got != tt.want {
				t.Errorf("ensureBlockStorage() = %v, want %v", got, tt.want)
			} else if got {
				if !reflect.DeepEqual(reflect.TypeOf(tt.args.router), tt.wantType) {
					t.Errorf("getBlockStorage() got = %v, want %v", reflect.TypeOf(tt.args.router), tt.wantType)
				}
			}
		})
	}
}
