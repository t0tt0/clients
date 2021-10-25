package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	extend_traits "github.com/HyperService-Consortium/go-ves/lib/backend/extend-traits"
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/abstraction"
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/database"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

type SessionDB struct {
	traits          abstraction.ORMTraits
	queryISCAddress extend_traits.Where1Func
}

func (sessionDB SessionDB) Query(opts ...abstraction.SessionQueryOption) (objs []database.Session, err error) {
	panic("implement me")
}

func (sessionDB SessionDB) Scan(desc interface{}, opts ...abstraction.SessionQueryOption) (err error) {
	panic("implement me")
}

func NewSessionDB(cb func(interface{}) abstraction.ORMTraits, _ module.Module) (*SessionDB, error) {
	traits := cb(Session{})
	return &SessionDB{
		traits:          traits,
		queryISCAddress: traits.Where1("isc_address = ?"),
	}, nil
}

func GetSessionDB(_ module.Module) (*SessionDB, error) {
	return new(SessionDB), nil
}

func wrapToSession(session interface{}, err error) (*Session, error) {
	if session == nil {
		return nil, err
	}
	return session.(*Session), err
}

func (sessionDB SessionDB) Migrate() error {
	return sessionDB.traits.Migrate()
}

func (sessionDB SessionDB) Create(s *Session) (int64, error) {
	return sessionDB.traits.Create(s)
}

func (sessionDB SessionDB) Update(s *Session) (int64, error) {
	return sessionDB.traits.Update(s)
}

func (sessionDB SessionDB) UpdateFields(s *Session, fields []string) (int64, error) {
	return sessionDB.traits.UpdateFields(s, fields)
}

func (sessionDB SessionDB) UpdateFields_(s *Session, db *dorm.DB, fields []string) (int64, error) {
	return sessionDB.traits.UpdateFields_(db, s, fields)
}

func (sessionDB SessionDB) UpdateFields__(s *Session, db dorm.SQLCommon, fields []string) (int64, error) {
	return sessionDB.traits.UpdateFields__(db, s, fields)
}

func (sessionDB SessionDB) Delete(s *Session, ) (int64, error) {
	return sessionDB.traits.Delete(s)
}

func (sessionDB *SessionDB) Filter(f *database.SessionFilter) (user []Session, err error) {
	err = sessionDB.traits.Filter(f, &user)
	return
}

func (sessionDB *SessionDB) FilterI(f interface{}) (interface{}, error) {
	return sessionDB.Filter(f.(*database.SessionFilter))
}

func (sessionDB *SessionDB) ID(id uint) (session *Session, err error) {
	return wrapToSession(sessionDB.traits.ID(id))
}

func (sessionDB *SessionDB) ID_(db *gorm.DB, id uint) (session *Session, err error) {
	return wrapToSession(sessionDB.traits.ID_(db, id))
}

func (sessionDB *SessionDB) QueryGUID(iscAddress string) (session *Session, err error) {
	return wrapToSession(sessionDB.queryISCAddress(iscAddress))
}

func (sessionDB *SessionDB) QueryGUIDByBytes(iscAddress []byte) (session *Session, err error) {
	return wrapToSession(sessionDB.queryISCAddress(database.EncodeAddress(iscAddress)))
}

type SessionQuery struct {
	db *gorm.DB
}

func (sessionDB *SessionDB) QueryChain() *SessionQuery {
	return &SessionQuery{db: sessionDB.traits.GetGormDB()}
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
