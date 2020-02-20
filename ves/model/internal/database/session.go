package database

import (
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"time"
)

type Session struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null" json:"id"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`

	ISCAddress       string `dorm:"isc_address" gorm:"column:isc_address;not_null" json:"isc_address"`
	UnderTransacting int64  `dorm:"under_transacting" gorm:"column:under_transacting;not_null" json:"under_transacting"`

	Status  uint8  `dorm:"status" gorm:"column:status;not_null" json:"status"`
	Content string `dorm:"content" gorm:"column:content;not_null" json:"content"`

	AccountsCount    int64 `dorm:"accounts_cnt" gorm:"column:accounts_cnt;not_null" json:"accounts_cnt"`
	TransactionCount int64 `dorm:"transaction_cnt" gorm:"column:transaction_cnt;not_null" json:"transaction_cnt"`

	//Accounts
	//Transactions
	//Acks

	//	Signer uip.Signer `json:"-" xorm:"-"`
}

func NewSession(iscAddress []byte) *Session {
	return &Session{
		ISCAddress:       EncodeAddress(iscAddress),
		UnderTransacting: 0,
		Status:           0,
	}
}

// TableName specification
func (Session) TableName() string {
	return "session"
}

func (s Session) GetID() uint {
	return s.ID
}

//todo
func (s *Session) GetGUID() []byte {
	//if s.decodedISCAddress == nil {
	//	s.decodedISCAddress = DecodeAddress(s.ISCAddress)
	//}
	return sugar.HandlerError(DecodeAddress(s.ISCAddress)).([]byte)
}

type SessionFilter = Filter
