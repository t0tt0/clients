package sessionservice

import (
	"errors"
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/ves/model"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
	"reflect"
	"testing"
)

func TestService_SessionRequireRawTransact(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sesDB := MockSessionDB(ctl)
	sesFSet := MockSessionFSet(ctl)
	sesAccountDB := MockSessionAccountDB(ctl)
	cVes := MockCentralVESClient(ctl)
	dns := MockChainDNS(ctl)

	f := createService(
		sesDB,
		sesFSet,
		sesAccountDB,
		cVes,
		dns,
	)
	newMockDNS(f, dns)

	// mock queryGUIDFindError
	sesDB.EXPECT().
		QueryGUIDByBytes(sessionIDFindError).Return(nil, errors.New("find error"))

	// mock queryGUIDNotFind
	sesDB.EXPECT().
		QueryGUIDByBytes(sessionIDNotFound).Return(nil, nil)

	//DeserializeTransactionError
	ses := &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDDeserializeTransactionError),
	}
	sesDB.EXPECT().
		QueryGUIDByBytes(sessionIDDeserializeTransactionError).Return(ses, nil)
	sesFSet.EXPECT().FindTransaction(sessionIDDeserializeTransactionError, int64(0)).
		Return([]byte(""), nil)

	ses = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDGetBlockChainError),
		//ID:               0,
		//CreatedAt:        time.Time{},
		//UpdatedAt:        time.Time{},
		//ISCAddress:       "",
		//UnderTransacting: 0,
		//Status:           0,
		//Content:          "",
		//AccountsCount:    0,
	}
	sesDB.EXPECT().
		QueryGUIDByBytes(sessionIDGetBlockChainError).Return(ses, nil)
	_, _, b := dataTransactionIntentWithBadChainID(t)
	newMockGoodGetTransactionIntent(b, sessionIDGetBlockChainError, sesFSet)

	type args struct {
		ctx context.Context
		in  *uiprpc.SessionRequireRawTransactRequest
	}
	tests := []struct {
		name     string
		fields   *Service
		args     args
		want     *uiprpc.SessionRequireRawTransactReply
		wantErr  bool
		wantCode int
	}{
		{name: "queryGUIDFindError", fields: f, args: args{
			ctx: context.Background(),
			in: &uiprpc.SessionRequireRawTransactRequest{
				SessionId: sessionIDFindError,
			},
		}, wantErr: true, wantCode: types.CodeSessionFindError},
		{name: "queryGUIDNotFind", fields: f, args: args{
			ctx: context.Background(),
			in: &uiprpc.SessionRequireRawTransactRequest{
				SessionId: sessionIDNotFound,
			},
		}, wantErr: true, wantCode: types.CodeSessionNotFind},
		{name: "DeserializeTransactionError", fields: f, args: args{
			ctx: context.Background(),
			in: &uiprpc.SessionRequireRawTransactRequest{
				SessionId: sessionIDDeserializeTransactionError,
			},
		}, wantErr: true, wantCode: types.CodeDeserializeTransactionError},
		{name: "GetBlockChainError", fields: f, args: args{
			ctx: context.Background(),
			in: &uiprpc.SessionRequireRawTransactRequest{
				SessionId: sessionIDGetBlockChainError,
			},
		}, wantErr: true, wantCode: types.CodeGetBlockChainInterfaceError},
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
			got, err := svc.SessionRequireRawTransact(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionRequireRawTransact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SessionRequireRawTransact() got = %v, want %v", got, tt.want)
			}
		})
	}
}
