package sessionservice

import (
	"errors"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/mock"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
	"testing"
)

func TestService_SessionAckForInit(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sesDB := MockSessionDB(ctl)
	sesFSet := MockSessionFSet(ctl)
	sesAccountDB := MockSessionAccountDB(ctl)
	cVes := MockCentralVESClient(ctl)

	f := createService(
		sesDB,
		sesFSet,
		sesAccountDB,
		cVes,
	)

	// mock queryGUIDFindError
	sesDB.EXPECT().
		QueryGUIDByBytes(sessionIDFindError).Return(nil, errors.New("find error"))

	// mock queryGUIDNotFind
	sesDB.EXPECT().
		QueryGUIDByBytes(sessionIDNotFound).Return(nil, nil)

	// mock queryFindSessionWithAcknowledgeError
	var ses = &model.Session{
	}
	var inFindSessionWithAcknowledgeError = &uiprpc.SessionAckForInitRequest{
		SessionId: sessionIDFindSessionWithAcknowledgeError,
		User:      nil,
		UserSignature: &uiprpc_base.Signature{
			SignatureType: 0,
			Content:       nil,
		},
	}
	sesDB.EXPECT().
		QueryGUIDByBytes(sessionIDFindSessionWithAcknowledgeError).Return(ses, nil)
	sesFSet.EXPECT().
		AckForInit(
			ses, inFindSessionWithAcknowledgeError.GetUser(),
			mock.MatchSignature(
				signaturer.FromRaw(inFindSessionWithAcknowledgeError.GetUserSignature().Content,
					inFindSessionWithAcknowledgeError.GetUserSignature().SignatureType))).
		Return(errors.New("ack error"))

	//queryFindSessionWithGetAcknowledgedError
	ses = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDFindSessionWithGetAcknowledgedError),
	}
	var inFindSessionWithGetAcknowledgedError = &uiprpc.SessionAckForInitRequest{
		SessionId: sessionIDFindSessionWithGetAcknowledgedError,
		User:      nil,
		UserSignature: &uiprpc_base.Signature{
			SignatureType: 0,
			Content:       nil,
		},
	}
	sesDB.EXPECT().
		QueryGUIDByBytes(sessionIDFindSessionWithGetAcknowledgedError).Return(ses, nil)
	sesFSet.EXPECT().
		AckForInit(
			ses, inFindSessionWithGetAcknowledgedError.GetUser(),
			mock.MatchSignature(
				signaturer.FromRaw(inFindSessionWithGetAcknowledgedError.GetUserSignature().Content,
					inFindSessionWithGetAcknowledgedError.GetUserSignature().SignatureType))).
		Return(nil)
	sesAccountDB.EXPECT().
		GetAcknowledged(ses.ISCAddress).
		Return(int64(0), errors.New("get acknowledged error"))

	ses = &model.Session{
		ISCAddress:       model.EncodeAddress(sessionIDOk),
		AccountsCount: 2,
	}
	var inOk = &uiprpc.SessionAckForInitRequest{
		SessionId: ses.GetGUID(),
		User:      nil,
		UserSignature: &uiprpc_base.Signature{
			SignatureType: 0,
			Content:       nil,
		},
	}
	sesDB.EXPECT().
		QueryGUIDByBytes(ses.GetGUID()).Return(ses, nil)
	sesFSet.EXPECT().
		AckForInit(
			ses, inOk.GetUser(),
			mock.MatchSignature(
				signaturer.FromRaw(inOk.GetUserSignature().Content,
					inOk.GetUserSignature().SignatureType))).
		Return(nil)
	sesAccountDB.EXPECT().
		GetAcknowledged(ses.ISCAddress).
		Return(int64(1), nil)

	ses = &model.Session{
		ISCAddress:       model.EncodeAddress(sessionIDOk2),
		AccountsCount: 1,
	}
	newMockGoodInternalPushTransaction(t, f, sessionIDOk2, sesFSet, cVes)
	var inOk2 = &uiprpc.SessionAckForInitRequest{
		SessionId: ses.GetGUID(),
		User:      nil,
		UserSignature: &uiprpc_base.Signature{
			SignatureType: 0,
			Content:       nil,
		},
	}
	sesDB.EXPECT().
		QueryGUIDByBytes(ses.GetGUID()).Return(ses, nil)
	sesFSet.EXPECT().
		AckForInit(
			ses, inOk2.GetUser(),
			mock.MatchSignature(
				signaturer.FromRaw(inOk2.GetUserSignature().Content,
					inOk2.GetUserSignature().SignatureType))).
		Return(nil)
	sesAccountDB.EXPECT().
		GetAcknowledged(ses.ISCAddress).
		Return(int64(1), nil)

	type args struct {
		ctx context.Context
		in  *uiprpc.SessionAckForInitRequest
	}
	tests := []struct {
		name     string
		fields   *Service
		args     args
		want     bool
		wantErr  bool
		wantCode int
	}{
		{name: "queryGUIDFindError", fields: f, args: args{
			ctx: context.Background(),
			in: &uiprpc.SessionAckForInitRequest{
				SessionId:     sessionIDFindError,
				User:          nil,
				UserSignature: nil,
			},
		}, wantErr: true, wantCode: types.CodeSessionFindError},
		{name: "queryGUIDNotFind", fields: f, args: args{
			ctx: context.Background(),
			in: &uiprpc.SessionAckForInitRequest{
				SessionId:     sessionIDNotFound,
				User:          nil,
				UserSignature: nil,
			},
		}, wantErr: true, wantCode: types.CodeSessionNotFind},
		{name: "queryFindSessionWithAcknowledgeError", fields: f, args: args{
			ctx: context.Background(),
			in:  inFindSessionWithAcknowledgeError,
		}, wantErr: true, wantCode: types.CodeSessionAcknowledgeError},
		{name: "queryFindSessionWithGetAcknowledgedError", fields: f, args: args{
			ctx: context.Background(),
			in:  inFindSessionWithGetAcknowledgedError,
		}, wantErr: true, wantCode: types.CodeSessionAccountGetAcknowledgedError},
		//{name: "callbackError", fields: f, args: args{
		//	ctx: context.Background(),
		//	in:  inOk2,
		//}, want: true},
		{name: "Ok", fields: f, args: args{
			ctx: context.Background(),
			in:  inOk,
		}, want: true},
		{name: "OkWithCallback", fields: f, args: args{
			ctx: context.Background(),
			in:  inOk2,
		}, want: true},
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
			got, err := svc.SessionAckForInit(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionAckForInit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				checkErrorCode(t, err, tt.wantCode)
				return
			}
			if got.GetOk() != tt.want {
				t.Errorf("SessionAckForInit() got = %v, want %v", got, tt.want)
			}
		})
	}
}
