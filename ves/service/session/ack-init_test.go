package sessionservice

import (
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
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

	var (
		sessionIDNotFound                            = []byte("abc")
		sessionIDFindError                           = []byte("abe")
		sessionIDFindSessionWithAcknowledgeError     = []byte("x")
		sessionIDFindSessionWithGetAcknowledgedError = []byte("y")
	)
	sesDB := MockSessionDB(ctl)
	sesFSet := MockSessionFSet(ctl)
	sesAccountDB := MockSessionAccountDB(ctl)

	f := createField(
		sesDB,
		sesFSet,
		sesAccountDB,
	)

	sesDB.EXPECT().
		QueryGUIDByBytes(sessionIDFindError).Return(nil, errors.New("find error"))

	sesDB.EXPECT().
		QueryGUIDByBytes(sessionIDNotFound).Return(nil, nil)

	var ses = &model.Session{
		//ID:               0,
		//CreatedAt:        time.Time{},
		//UpdatedAt:        time.Time{},
		//ISCAddress:       "",
		//UnderTransacting: 0,
		//Status:           0,
		//Content:          "",
		//AccountsCount:    0,
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

	//sessionIDFindSessionWithGetAcknowledgedError

	ses = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDFindSessionWithGetAcknowledgedError),
		//ID:               0,
		//CreatedAt:        time.Time{},
		//UpdatedAt:        time.Time{},
		//ISCAddress:       "",
		//UnderTransacting: 0,
		//Status:           0,
		//Content:          "",
		//AccountsCount:    0,
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

	type args struct {
		ctx context.Context
		in  *uiprpc.SessionAckForInitRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     bool
		wantErr  bool
		wantCode int
	}{
		{"queryGUIDFindError", f, args{
			context.Background(),
			&uiprpc.SessionAckForInitRequest{
				SessionId:     sessionIDFindError,
				User:          nil,
				UserSignature: nil,
			},
		}, false, true, types.CodeSessionFindError},
		{"queryGUIDNotFind", f, args{
			context.Background(),
			&uiprpc.SessionAckForInitRequest{
				SessionId:     sessionIDNotFound,
				User:          nil,
				UserSignature: nil,
			},
		}, false, true, types.CodeSessionNotFind},
		{"queryFindSessionWithAcknowledgeError", f, args{
			context.Background(),
			inFindSessionWithAcknowledgeError,
		}, false, true, types.CodeSessionAcknowledgeError},
		{"queryFindSessionWithGetAcknowledgedError", f, args{
			context.Background(),
			inFindSessionWithGetAcknowledgedError,
		}, false, true, types.CodeSessionAccountGetAcknowledgedError},
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
				if tt.wantCode != types.CodeOK {
					if f, ok := wrapper.FromError(err); ok {
						if f.GetCode() != tt.wantCode {
							t.Errorf("not expected code, error code %v, wantCode %v", f.GetCode(), tt.wantCode)
						} else {
							fmt.Println("good", err)
						}
					} else {
						t.Error("not frame error wrapped")
					}
				}
				return
			}
			if got.GetOk() != tt.want {
				t.Errorf("SessionAckForInit() got = %v, want %v", got, tt.want)
			}
		})
	}
}
