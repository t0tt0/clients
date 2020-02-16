package sessionservice

import (
	"context"
	"encoding/json"
	"errors"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestService_pushTransaction(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sesDB := MockSessionDB(ctl)
	sesFSet := MockSessionFSet(ctl)
	sesAccountDB := MockSessionAccountDB(ctl)
	cVes := MockCentralVESClient(ctl)



	f := createField(
		sesDB,
		sesFSet,
		sesAccountDB,
		cVes,
	)

	var sesFindTransactionError = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDFindTransactionError),
		//ID:               0,
		//CreatedAt:        time.Time{},
		//UpdatedAt:        time.Time{},
		//ISCAddress:       "",
		//UnderTransacting: 0,
		//Status:           0,
		//Content:          "",
		//AccountsCount:    0,
	}
	sesFSet.EXPECT().FindTransaction(sessionIDFindTransactionError, int64(0)).
		Return(nil, errors.New("find transaction error"))

	var sesDeserializeTransactionError = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDDeserializeTransactionError),
		//ID:               0,
		//CreatedAt:        time.Time{},
		//UpdatedAt:        time.Time{},
		//ISCAddress:       "",
		//UnderTransacting: 0,
		//Status:           0,
		//Content:          "",
		//AccountsCount:    0,
	}
	sesFSet.EXPECT().FindTransaction(sessionIDDeserializeTransactionError, int64(0)).
		Return([]byte(""), nil)

	var srcAcc = &uiprpc_base.Account{
		ChainId: 233,
		Address: []byte{2, 3, 3},
	}
	var ti = opintent.TransactionIntent{
		TransType: 0,
		Src:       srcAcc.Address,
		Dst:       nil,
		Meta:      nil,
		Amt:       "3e8",
		ChainID:   srcAcc.ChainId,
	}
	b, err := json.Marshal(&ti)
	if err != nil {
		t.Fatal("ser", err)
	}
	var sesAttestationSendError = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDAttestationSendError),
		//ID:               0,
		//CreatedAt:        time.Time{},
		//UpdatedAt:        time.Time{},
		//ISCAddress:       "",
		//UnderTransacting: 0,
		//Status:           0,
		//Content:          "",
		//AccountsCount:    0,
	}
	sesFSet.EXPECT().FindTransaction(sessionIDAttestationSendError, int64(0)).
		Return(b, nil)
	cVes.EXPECT().InternalAttestationSending(gomock.Any(), &uiprpc.InternalRequestComingRequest{
		SessionId: sesAttestationSendError.GetGUID(),
		Host:      f.cfg.BaseParametersConfig.ExposeHost,
		Accounts:  []*uiprpc_base.Account{srcAcc},
	}).Return(nil, errors.New("send error"))

	var sesAttestationSendErrorNotOk = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDAttestationSendErrorNotOk),
		//ID:               0,
		//CreatedAt:        time.Time{},
		//UpdatedAt:        time.Time{},
		//ISCAddress:       "",
		//UnderTransacting: 0,
		//Status:           0,
		//Content:          "",
		//AccountsCount:    0,
	}
	sesFSet.EXPECT().FindTransaction(sessionIDAttestationSendErrorNotOk, int64(0)).
		Return(b, nil)
	cVes.EXPECT().InternalAttestationSending(gomock.Any(), &uiprpc.InternalRequestComingRequest{
		SessionId: sesAttestationSendErrorNotOk.GetGUID(),
		Host:      f.cfg.BaseParametersConfig.ExposeHost,
		Accounts:  []*uiprpc_base.Account{srcAcc},
	}).Return(&uiprpc.InternalRequestComingReply{Ok:false}, nil)

	var sesAttestationSendErrorNotOk2 = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDAttestationSendErrorNotOk2),
		//ID:               0,
		//CreatedAt:        time.Time{},
		//UpdatedAt:        time.Time{},
		//ISCAddress:       "",
		//UnderTransacting: 0,
		//Status:           0,
		//Content:          "",
		//AccountsCount:    0,
	}
	sesFSet.EXPECT().FindTransaction(sessionIDAttestationSendErrorNotOk2, int64(0)).
		Return(b, nil)
	cVes.EXPECT().InternalAttestationSending(gomock.Any(), &uiprpc.InternalRequestComingRequest{
		SessionId: sesAttestationSendErrorNotOk2.GetGUID(),
		Host:      f.cfg.BaseParametersConfig.ExposeHost,
		Accounts:  []*uiprpc_base.Account{srcAcc},
	}).Return(nil, nil)

	var sesOk = &model.Session{
		ISCAddress: model.EncodeAddress(sessionIDPushTransactionNotNil),
		//ID:               0,
		//CreatedAt:        time.Time{},
		//UpdatedAt:        time.Time{},
		//ISCAddress:       "",
		//UnderTransacting: 0,
		//Status:           0,
		//Content:          "",
		//AccountsCount:    0,
	}
	sesFSet.EXPECT().FindTransaction(sessionIDPushTransactionNotNil, int64(0)).
		Return(b, nil)
	cVes.EXPECT().InternalAttestationSending(gomock.Any(), &uiprpc.InternalRequestComingRequest{
		SessionId: sesOk.GetGUID(),
		Host:      f.cfg.BaseParametersConfig.ExposeHost,
		Accounts:  []*uiprpc_base.Account{srcAcc},
	}).Return(&uiprpc.InternalRequestComingReply{Ok:true}, nil)

	type args struct {
		ctx           context.Context
		ses           *model.Session
		transactionID int64
	}
	tests := []struct {
		name     string
		fields   fields
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
