package vesclient

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
	"time"
)

type sessionTraits struct {
	traits
}

func (m *modelModule) injectSessionTraits() error {
	m.sessionTraits.traits = m.newTraits(Session{})
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

	Intents      string `dorm:"intents" gorm:"column:intents;not_null" json:"intents"`
	Dependencies string `dorm:"dependencies" gorm:"column:dependencies;not_null" json:"dependencies"`

	ISCAddress string `dorm:"isc_address" gorm:"column:isc_address;not_null" json:"isc_address"`
	NSBHost    string `dorm:"nsb_host" gorm:"column:nsb_host;not_null" json:"nsb_host"`
	VESHost    string `dorm:"ves_host" gorm:"column:ves_host;not_null" json:"ves_host"`

	UnderTransacting int64 `dorm:"under_transacting" gorm:"column:under_transacting;not_null" json:"under_transacting"`
	Status           uint8 `dorm:"status" gorm:"column:status;not_null" json:"status"`

	decodedISCAddress []byte `gorm:"-" json:"-"`
}

// TableName specification
func (Session) TableName() string {
	return "session"
}

func NewSession() *Session {
	return &Session{}
}

func (ses Session) migrate(dep *modelModule) error {
	return dep.sessionTraits.Migrate()
}

func (ses Session) migration(dep *modelModule) func() error {
	return func() error {
		return ses.migrate(dep)
	}
}

func (ses Session) GetID() uint {
	return ses.ID
}

func (ses Session) GetGUID() ([]byte, error) {
	if ses.decodedISCAddress == nil {
		var err error
		ses.decodedISCAddress, err = decodeAddress(ses.ISCAddress)
		if err != nil {
			return nil, err
		}
	}
	return ses.decodedISCAddress, nil
}

func (sessionDB SessionDB) Create(ses *Session) (int64, error) {
	return sessionDB.module.sessionTraits.Create(ses)
}

func (sessionDB SessionDB) Update(ses *Session) (int64, error) {
	return sessionDB.module.sessionTraits.Update(ses)
}

func (sessionDB SessionDB) UpdateFields(ses *Session, fields []string) (int64, error) {
	return sessionDB.module.sessionTraits.UpdateFields(ses, fields)
}

func (sessionDB SessionDB) UpdateFields_(ses *Session, db *dorm.DB, fields []string) (int64, error) {
	return sessionDB.module.sessionTraits.UpdateFields_(db, ses, fields)
}

func (sessionDB SessionDB) UpdateFields__(ses *Session, db dorm.SQLCommon, fields []string) (int64, error) {
	return sessionDB.module.sessionTraits.UpdateFields__(db, ses, fields)
}

func (sessionDB SessionDB) Delete(ses *Session) (int64, error) {
	return sessionDB.Delete(ses)
}

type SessionDB struct {
	db     *gorm.DB
	module *modelModule
}

func NewSessionDB(m Module) (*SessionDB, error) {
	return &SessionDB{db: m.GormDB(), module: m.ModelModule()}, nil
}

func GetSessionDB(_ module.Module) (*SessionDB, error) {
	return new(SessionDB), nil
}

func (sessionDB *SessionDB) ID(id uint) (session *Session, err error) {
	return wrapToSession(sessionDB.module.sessionTraits.ID(id))
}

func (sessionDB *SessionDB) ID_(db *gorm.DB, id uint) (session *Session, err error) {
	return wrapToSession(sessionDB.module.sessionTraits.ID_(db, id))
}

type SessionQuery struct {
	db *gorm.DB
}

func (sessionDB *SessionDB) QueryChain() *SessionQuery {
	return &SessionQuery{db: sessionDB.db}
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

type SessionDBInterface interface {
	Create(ses *Session) (int64, error)
	Update(ses *Session) (int64, error)
	UpdateFields(ses *Session, fields []string) (int64, error)
	UpdateFields_(ses *Session, db *dorm.DB, fields []string) (int64, error)
	UpdateFields__(ses *Session, db dorm.SQLCommon, fields []string) (int64, error)
	Delete(ses *Session) (int64, error)
	ID(id uint) (session *Session, err error)
	ID_(db *gorm.DB, id uint) (session *Session, err error)
	QueryChain() *SessionQuery
}
