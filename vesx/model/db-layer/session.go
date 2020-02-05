package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	sessionTraits Traits
)

func injectSessionTraits() error {
	sessionTraits = NewTraits(Session{})
	return nil
}

func wrapToSession(session interface{}, err error) (*Session, error) {
	if session == nil {
		return nil, err
	}
	return session.(*Session), err
}

type Session struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`
}

// TableName specification
func (Session) TableName() string {
	return "session"
}

func (Session) migrate() error {
	return sessionTraits.Migrate()
}

func (d Session) GetID() uint {
	return d.ID
}

func (d *Session) Create() (int64, error) {
	return sessionTraits.Create(d)
}

func (d *Session) Update() (int64, error) {
	return sessionTraits.Update(d)
}

func (d *Session) UpdateFields(fields []string) (int64, error) {
	return sessionTraits.UpdateFields(d, fields)
}

func (d *Session) UpdateFields_(db *dorm.DB, fields []string) (int64, error) {
	return sessionTraits.UpdateFields_(db, d, fields)
}

func (d *Session) UpdateFields__(db dorm.SQLCommon, fields []string) (int64, error) {
	return sessionTraits.UpdateFields__(db, d, fields)
}

func (d *Session) Delete() (int64, error) {
	return sessionTraits.Delete(d)
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

func (sessionDB *SessionDB) ID_(db *gorm.DB, id uint) (goods *Session, err error) {
	return wrapToSession(sessionTraits.ID_(db, id))
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
