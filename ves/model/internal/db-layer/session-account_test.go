package dblayer_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/ves/config"
	dblayer "github.com/HyperService-Consortium/go-ves/ves/model/internal/db-layer"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"testing"
)


func TestSessionAccount_UpdateAcknowledged(t *testing.T) {
	m := dep.Require(config.ModulePath.Minimum.Global.SQLMock).(sqlmock.Sqlmock)
	//a := []byte("1234")

	m.ExpectBegin()
	m.ExpectExec(`UPDATE "session_account"`).
		WithArgs(
			false, // acknowledged
			"",    // where session_id =
			0,     // where chain_id =
			"",    // where address =
		).WillReturnResult(sqlmock.NewResult(1, 1))
	m.ExpectCommit()

	m.ExpectBegin()
	m.ExpectExec(`UPDATE "session_account"`).
		WithArgs(
			true, // acknowledged
			"",   // where session_id =
			1,    // where chain_id =
			"",   // where address =
		).WillReturnResult(sqlmock.NewResult(1, 1))
	m.ExpectCommit()

	type fields struct {
		SessionID    string
		ChainID      uip.ChainIDUnderlyingType
		Address      string
		Acknowledged bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		{name: "", fields: fields{
			SessionID:    "",
			ChainID:      0,
			Address:      "",
			Acknowledged: false,
		}, want: 1},
		{name: "", fields: fields{
			SessionID:    "",
			ChainID:      1,
			Address:      "",
			Acknowledged: true,
		}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := &dblayer.SessionAccount{
				SessionID:    tt.fields.SessionID,
				ChainID:      tt.fields.ChainID,
				Address:      tt.fields.Address,
				Acknowledged: tt.fields.Acknowledged,
			}
			got, err := sugar.HandlerError(p.NewSessionAccountDB(dep)).(*dblayer.SessionAccountDB).UpdateAcknowledged(sa)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAcknowledged() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateAcknowledged() got = %v, want %v", got, tt.want)
			}
		})
	}
}
