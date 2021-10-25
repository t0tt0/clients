package sessionservice

import (
	ethbni "github.com/HyperService-Consortium/go-ves/lib/bni/eth"
	tenbni "github.com/HyperService-Consortium/go-ves/lib/bni/ten"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

var (
	ethBNPType = reflect.TypeOf(new(ethbni.BN))
	tenBNPType = reflect.TypeOf(new(tenbni.BN))
)

func TestService_getBlockChainInterface(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	dns := MockChainDNS(ctl)
	f := createService(
		dns,
	)

	newMockDNS(f, dns)

	type args struct {
		chainID uint64
	}
	tests := []struct {
		name     string
		fields   *Service
		args     args
		wantType reflect.Type
		wantErr  bool
		wantCode int
	}{
		{name: "getEthereumChainRouter", fields: f, args: args{
			chainID: ethereumChainID,
		}, wantType: ethBNPType},
		{name: "getTendermintChainRouter", fields: f, args: args{
			chainID: tendermintChainID,
		}, wantType: tenBNPType},
		{name: "chainIDNotFound", fields: f, args: args{
			chainID: unknownChainID,
		}, wantErr: true, wantCode: types.CodeChainIDNotFound},
		{name: "chainTypeNotFound", fields: f, args: args{
			chainID: unknownChainTypeID,
		}, wantErr: true, wantCode: types.CodeChainTypeNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &Service{
				cfg:            tt.fields.cfg,
				key:            tt.fields.key,
				accountDB:      tt.fields.accountDB,
				db:             tt.fields.db,
				sesFSet:        tt.fields.sesFSet,
				opInitializer:  tt.fields.opInitializer,
				signer:         tt.fields.signer,
				logger:         tt.fields.logger,
				cVes:           tt.fields.cVes,
				respAccount:    tt.fields.respAccount,
				storage:        tt.fields.storage,
				storageHandler: tt.fields.storageHandler,
				dns:            tt.fields.dns,
				nsbClient:      tt.fields.nsbClient,
			}
			got, err := svc.getBlockChainInterface(tt.args.chainID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBlockChainInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				checkErrorCode(t, err, tt.wantCode)
				return
			}
			if !reflect.DeepEqual(reflect.TypeOf(got), tt.wantType) {
				t.Errorf("getBlockChainInterface() got = %v, want %v", reflect.TypeOf(got), tt.wantType)
			}
		})
	}
}
