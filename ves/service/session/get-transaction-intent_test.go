package sessionservice

import (
	"errors"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestService_getTransactionIntent(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sesFSet := MockSessionFSet(ctl)

	f := createService(
		sesFSet,
	)

	//FindTransactionError
	sesFSet.EXPECT().FindTransaction(sessionIDFindTransactionError, int64(0)).
		Return(nil, errors.New("find transaction error"))

	//DeserializeTransactionError
	sesFSet.EXPECT().FindTransaction(sessionIDDeserializeTransactionError, int64(0)).
		Return([]byte(""), nil)

	//Ok
	_, ti, b := dataGoodTransactionIntent(t)
	sesFSet.EXPECT().FindTransaction(sessionIDOk, int64(0)).
		Return(b, nil)

	type args struct {
		sessionID     []byte
		transactionID int64
	}
	tests := []struct {
		name     string
		fields   *Service
		args     args
		want     *opintent.TransactionIntent
		wantErr  bool
		wantCode int
	}{
		{name: "FindTransactionError", fields: f, args: args{
			sessionID:     sessionIDFindTransactionError,
			transactionID: 0,
		}, wantErr: true, wantCode: types.CodeTransactionFindError},
		{name: "DeserializeTransactionError", fields: f, args: args{
			sessionID:     sessionIDDeserializeTransactionError,
			transactionID: 0,
		}, wantErr: true, wantCode: types.CodeDeserializeTransactionError},
		{name: "Ok", fields: f, args: args{
			sessionID:     sessionIDOk,
			transactionID: 0,
		}, wantErr: false, want: ti},
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
			got, err := svc.getTransactionIntent(tt.args.sessionID, tt.args.transactionID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTransactionIntent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				checkErrorCode(t, err, tt.wantCode)
				return
			}

			if !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("getTransactionIntent() got = %v, want %v", got, tt.want)
			}
		})
	}
}
