package sessionservice

import (
	"context"
	"errors"
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/HyperService-Consortium/go-ves/grpc/uiprpc-base"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/ves/model"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestService_pushTransaction(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sesFSet := MockSessionFSet(ctl)
	cVes := MockCentralVESClient(ctl)

	f := createService(
		sesFSet,
		cVes,
	)

	//FindTransactionError
	var sesFindTransactionError = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDFindTransactionError),
	}
	sesFSet.EXPECT().FindTransaction(sessionIDFindTransactionError, int64(0)).
		Return(nil, errors.New("find transaction error"))

	//DeserializeTransactionError
	var sesDeserializeTransactionError = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDDeserializeTransactionError),
	}
	sesFSet.EXPECT().FindTransaction(sessionIDDeserializeTransactionError, int64(0)).
		Return([]byte(""), nil)

	srcAcc, _, b := dataGoodTransactionIntent(t)
	//AttestationSendError
	var sesAttestationSendError = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDAttestationSendError),
	}
	sesFSet.EXPECT().FindTransaction(sessionIDAttestationSendError, int64(0)).
		Return(b, nil)
	cVes.EXPECT().InternalAttestationSending(gomock.Any(), &uiprpc.InternalRequestComingRequest{
		SessionId: sesAttestationSendError.GetGUID(),
		Host:      f.cfg.BaseParametersConfig.ExposeHost,
		Accounts:  []*uiprpc_base.Account{srcAcc},
	}).Return(nil, errors.New("send error"))

	//AttestationSendErrorNotOk
	var sesAttestationSendErrorNotOk = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDAttestationSendErrorNotOk),
	}
	sesFSet.EXPECT().FindTransaction(sessionIDAttestationSendErrorNotOk, int64(0)).
		Return(b, nil)
	cVes.EXPECT().InternalAttestationSending(gomock.Any(), &uiprpc.InternalRequestComingRequest{
		SessionId: sesAttestationSendErrorNotOk.GetGUID(),
		Host:      f.cfg.BaseParametersConfig.ExposeHost,
		Accounts:  []*uiprpc_base.Account{srcAcc},
	}).Return(&uiprpc.InternalRequestComingReply{Ok: false}, nil)

	//AttestationSendErrorNotOk2
	var sesAttestationSendErrorNotOk2 = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDAttestationSendErrorNotOk2),
	}
	sesFSet.EXPECT().FindTransaction(sessionIDAttestationSendErrorNotOk2, int64(0)).
		Return(b, nil)
	cVes.EXPECT().InternalAttestationSending(gomock.Any(), &uiprpc.InternalRequestComingRequest{
		SessionId: sesAttestationSendErrorNotOk2.GetGUID(),
		Host:      f.cfg.BaseParametersConfig.ExposeHost,
		Accounts:  []*uiprpc_base.Account{srcAcc},
	}).Return(nil, nil)

	//Ok
	newMockGoodInternalPushTransaction(t, f, sessionIDPushTransactionNotNil, sesFSet, cVes)
	var sesOk = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDPushTransactionNotNil),
	}

	type args struct {
		ctx           context.Context
		ses           *model.Session
		transactionID int64
	}
	tests := []struct {
		name     string
		fields   *Service
		args     args
		wantErr  bool
		wantCode int
	}{
		{name: "FindTransactionError", fields: f, args: args{
			ctx: context.Background(),
			ses: sesFindTransactionError,
		}, wantErr: true, wantCode: types.CodeTransactionFindError},
		{name: "DeserializeTransactionError", fields: f, args: args{
			ctx: context.Background(),
			ses: sesDeserializeTransactionError,
		}, wantErr: true, wantCode: types.CodeDeserializeTransactionError},
		{name: "AttestationSendError", fields: f, args: args{
			ctx: context.Background(),
			ses: sesAttestationSendError,
		}, wantErr: true, wantCode: types.CodeAttestationSendError},
		{name: "AttestationSendErrorNotOk", fields: f, args: args{
			ctx: context.Background(),
			ses: sesAttestationSendErrorNotOk,
		}, wantErr: true, wantCode: types.CodeAttestationSendError},
		{name: "AttestationSendErrorNotOk2", fields: f, args: args{
			ctx: context.Background(),
			ses: sesAttestationSendErrorNotOk2,
		}, wantErr: true, wantCode: types.CodeAttestationSendError},
		{name: "Ok", fields: f, args: args{
			ctx: context.Background(),
			ses: sesOk,
		}, wantErr: false},
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
			if err := svc.pushTransaction(tt.args.ctx, tt.args.ses, tt.args.transactionID); (err != nil) != tt.wantErr {
				t.Errorf("pushTransaction() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				checkErrorCode(t, err, tt.wantCode)
				return
			}
		})
	}
}
