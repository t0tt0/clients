package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/go-ves/lib/encoding"
	extend_traits "github.com/Myriad-Dreamin/go-ves/lib/extend-traits"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	sessionTraits          Traits
	sessionQueryISCAddress extend_traits.Where1Func
)

func injectSessionTraits() error {
	sessionTraits = NewTraits(Session{})

	sessionQueryISCAddress = sessionTraits.Where1("isc_address = ?")
	return nil
}

func wrapToSession(session interface{}, err error) (*Session, error) {
	if session == nil {
		return nil, err
	}
	return session.(*Session), err
}

type Session struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null" json:"id"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`

	ISCAddress       string `dorm:"isc_address" gorm:"column:isc_address;not_null" json:"isc_address"`
	UnderTransacting int64  `dorm:"under_transacting" gorm:"column:under_transacting;not_null" json:"under_transacting"`
	Status           uint8  `dorm:"status" gorm:"column:status;not_null" json:"status"`
	Content          string `dorm:"content" gorm:"column:content;not_null" json:"content"`

	AccountsCount int64 `dorm:"accounts_cnt" gorm:"column:accounts_cnt;not_null" json:"accounts_cnt"`

	//Accounts
	//Transactions
	//Acks

	decodedISCAddress []byte `gorm:"-" json:"-"`
	//	Signer uiptypes.Signer `json:"-" xorm:"-"`
}

func NewSession(iscAddress []byte) *Session {
	return &Session{
		ISCAddress:        encoding.EncodeBase64(iscAddress),
		decodedISCAddress: iscAddress,
		UnderTransacting:  0,
		Status:            0,
	}
}

// TableName specification
func (Session) TableName() string {
	return "session"
}

func (Session) migrate() error {
	return sessionTraits.Migrate()
}

func (s Session) GetID() uint {
	return s.ID
}

func (s Session) GetGUID() []byte {
	if s.decodedISCAddress == nil {
		s.decodedISCAddress = DecodeAddress(s.ISCAddress)
	}
	return s.decodedISCAddress
}

func (s *Session) Create() (int64, error) {
	return sessionTraits.Create(s)
}

func (s *Session) Update() (int64, error) {
	return sessionTraits.Update(s)
}

func (s *Session) UpdateFields(fields []string) (int64, error) {
	return sessionTraits.UpdateFields(s, fields)
}

func (s *Session) UpdateFields_(db *dorm.DB, fields []string) (int64, error) {
	return sessionTraits.UpdateFields_(db, s, fields)
}

func (s *Session) UpdateFields__(db dorm.SQLCommon, fields []string) (int64, error) {
	return sessionTraits.UpdateFields__(db, s, fields)
}

func (s *Session) Delete() (int64, error) {
	return sessionTraits.Delete(s)
}

type SessionDB struct{}

func NewSessionDB(_ module.Module) (*SessionDB, error) {
	return new(SessionDB), nil
}

func GetSessionDB(_ module.Module) (*SessionDB, error) {
	return new(SessionDB), nil
}

func (sessionDB *SessionDB) Filter(f *Filter) (user []Session, err error) {
	err = sessionTraits.Filter(f, &user)
	return
}

func (sessionDB *SessionDB) FilterI(f interface{}) (interface{}, error) {
	return sessionDB.Filter(f.(*Filter))
}

func (sessionDB *SessionDB) ID(id uint) (session *Session, err error) {
	return wrapToSession(sessionTraits.ID(id))
}

func (sessionDB *SessionDB) ID_(db *gorm.DB, id uint) (session *Session, err error) {
	return wrapToSession(sessionTraits.ID_(db, id))
}

func (sessionDB *SessionDB) QueryGUID(iscAddress string) (session *Session, err error) {
	return wrapToSession(sessionQueryISCAddress(iscAddress))
}

func (sessionDB *SessionDB) QueryGUIDByBytes(iscAddress []byte) (session *Session, err error) {
	return wrapToSession(sessionQueryISCAddress(EncodeAddress(iscAddress)))
}

type SessionQuery struct {
	db *gorm.DB
}

func (sessionDB *SessionDB) QueryChain() *SessionQuery {
	return &SessionQuery{db: p.GormDB}
}

func (sessionDB *SessionQuery) Order(order string, reorder ...bool) *SessionQuery {
	sessionDB.db = sessionDB.db.Order(order, reorder...)
	return sessionDB
}

func (sessionDB *SessionQuery) Page(page, pageSize int) *SessionQuery {
	sessionDB.db = sessionDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return sessionDB
}
func (sessionDB *SessionQuery) BeforeID(id uint) *SessionQuery {
	sessionDB.db = sessionDB.db.Where("id <= ?", id)
	return sessionDB
}

func (sessionDB *SessionQuery) Preload() *SessionQuery {
	sessionDB.db = sessionDB.db.Set("gorm:auto_preload", true)
	return sessionDB
}

func (sessionDB *SessionQuery) Query() (sessions []Session, err error) {
	err = sessionDB.db.Find(&sessions).Error
	return
}

func (sessionDB *SessionQuery) Scan(desc interface{}) (err error) {
	err = sessionDB.db.Scan(desc).Error
	return
}
