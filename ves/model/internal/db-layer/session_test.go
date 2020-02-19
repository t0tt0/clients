package dblayer

import (
	"bytes"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/database"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestSession_GetGUID(t *testing.T) {
	var a = assert.New(t)
	var ses = new(Session)
	var b0, b1 = []byte("123"), []byte("124")
	ses.ISCAddress = database.EncodeAddress(b0)
	a.True(bytes.Equal(ses.GetGUID(), b0))
	ses.ISCAddress = database.EncodeAddress(b1)
	a.True(bytes.Equal(ses.GetGUID(), b1))
}

func TestSession_Create(t *testing.T) {
	m := dep.Require(config.ModulePath.Minimum.Global.SQLMock).(sqlmock.Sqlmock)

	m.ExpectBegin()
	m.ExpectExec(`INSERT INTO "session"`).WillReturnResult(sqlmock.NewResult(1, 1))
	m.ExpectCommit()

	//sqlmock.ExpectedQuery()
	type fields struct {
		ID               uint
		CreatedAt        time.Time
		UpdatedAt        time.Time
		ISCAddress       string
		UnderTransacting int64
		Status           uint8
		Content          string
		AccountsCount    int64
	}
	tests := []struct {
		name       string
		fields     fields
		want       int64
		wantErr    bool
		expectedID uint
	}{
		{name: "will update id", fields: fields{
			ID: 0,
		}, want: 1, expectedID: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				ID:               tt.fields.ID,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				ISCAddress:       tt.fields.ISCAddress,
				UnderTransacting: tt.fields.UnderTransacting,
				Status:           tt.fields.Status,
				Content:          tt.fields.Content,
				AccountsCount:    tt.fields.AccountsCount,
			}
			got, err := sugar.HandlerError(NewSessionAccountDB(n, dep)).(*SessionDB).Create(s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}

			if s.ID != tt.expectedID {
				t.Errorf("Create() got = %v, want %v", s.ID, tt.expectedID)
			}
		})
	}
}

type scf struct {
	ID uint

	ISCAddress       string `dorm:"isc_address" gorm:"column:isc_address;not_null" json:"isc_address"`
	UnderTransacting int64  `dorm:"under_transacting" gorm:"column:under_transacting;not_null" json:"under_transacting"`
	Status           uint8  `dorm:"status" gorm:"column:status;not_null" json:"status"`
	Content          string `dorm:"content" gorm:"column:content;not_null" json:"content"`

	AccountsCount int64 `dorm:"accounts_cnt" gorm:"column:accounts_cnt;not_null" json:"accounts_cnt"`
}

func filterComparingField(s *Session) scf {
	return scf{
		ID:               s.ID,
		ISCAddress:       s.ISCAddress,
		UnderTransacting: s.UnderTransacting,
		Status:           s.Status,
		Content:          s.Content,
		AccountsCount:    s.AccountsCount,
	}
}

// todo: test scan

func TestSession_Update(t *testing.T) {
	m := dep.Require(config.ModulePath.Minimum.Global.SQLMock).(sqlmock.Sqlmock)
	a := []byte("1234")
	m.ExpectBegin()
	m.ExpectExec(`INSERT INTO "session"`).WillReturnResult(sqlmock.NewResult(1, 1))
	m.ExpectCommit()

	m.ExpectBegin()
	m.ExpectExec(`UPDATE "session"`).
		WithArgs(
			sqlmock.AnyArg(), // updated_at
			"",               // isc_address
			0, 0, "",
			1, 0, // accounts/transaction count
			1, // where session_id = ?
		).WillReturnResult(sqlmock.NewResult(1, 1))
	m.ExpectCommit()

	m.ExpectBegin()
	m.ExpectExec(`UPDATE "session"`).
		WithArgs(
			sqlmock.AnyArg(),          // updated_at
			database.EncodeAddress(a), // isc_address
			0, 0, "",
			0, 0, // accounts/transaction count
			1, // where session_id = ?
		).WillReturnResult(sqlmock.NewResult(1, 1))
	m.ExpectCommit()

	type fields struct {
		ID               uint
		CreatedAt        time.Time
		UpdatedAt        time.Time
		ISCAddress       string
		UnderTransacting int64
		Status           uint8
		Content          string
		AccountsCount    int64
	}
	tests := []struct {
		name    string
		fields  fields
		result  *Session
		want    int64
		wantErr bool
	}{
		{name: "lost-id", fields: fields{
			ID: 0,
		}, result: &Session{
			ID: 1,
		}, want: 1},
		{name: "update-fields-to-non-zero", fields: fields{
			ID:            1,
			AccountsCount: 1,
		}, result: &Session{
			ID:            1,
			AccountsCount: 1,
		}, want: 1},
		{name: "update-fields-to-zero", fields: fields{
			ID:            1,
			ISCAddress:    database.EncodeAddress(a),
			AccountsCount: 0,
		}, result: &Session{
			ID:            1,
			ISCAddress:    database.EncodeAddress(a),
			AccountsCount: 0,
		}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				ID:               tt.fields.ID,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				ISCAddress:       tt.fields.ISCAddress,
				UnderTransacting: tt.fields.UnderTransacting,
				Status:           tt.fields.Status,
				Content:          tt.fields.Content,
				AccountsCount:    tt.fields.AccountsCount,
			}
			got, err := sugar.HandlerError(NewSessionAccountDB(n, dep)).(*SessionDB).Update(s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}

			if err != nil {
				return
			}

			if !reflect.DeepEqual(filterComparingField(tt.result), filterComparingField(s)) {
				t.Errorf("Update() got = %v, want %v", filterComparingField(s), filterComparingField(tt.result))
			}
		})
	}
}
