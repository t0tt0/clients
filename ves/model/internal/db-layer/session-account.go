package dblayer

import (
	extend_traits "github.com/Myriad-Dreamin/go-ves/lib/extend-traits"
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/database"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

var ()

func wrapToSessionAccount(sessionAccount interface{}, err error) (*SessionAccount, error) {
	if sessionAccount == nil {
		return nil, err
	}
	return sessionAccount.(*SessionAccount), err
}

type SessionAccountDB struct {
	traits abstraction.ORMTraits

	queryID      extend_traits.Where1Func
	_queryID     extend_traits.Where1Func_
	acknowledged extend_traits.Count1Func
	total        extend_traits.Count1Func
}

func (s SessionAccountDB) GetTraits() abstraction.ORMTraits {
	return s.traits
}

func (s SessionAccountDB) Query(opts ...abstraction.SessionAccountQueryOption) (sas []database.SessionAccount, err error) {
	panic("implement me")
}

func (s SessionAccountDB) Scan(obj interface{}, opts ...abstraction.SessionAccountQueryOption) (err error) {
	panic("implement me")
}

func NewSessionAccountDB(cb func(interface{}) abstraction.ORMTraits, _ module.Module) (*SessionAccountDB, error) {
	traits := cb(SessionAccount{})
	return &SessionAccountDB{
		traits:       traits,
		queryID:      traits.Where1("session_id = ?"),
		_queryID:     traits.Where1_("session_id = ?"),
		acknowledged: traits.Count1("session_id = ? and acknowledged = 1"),
		total:        traits.Count1("session_id = ?"),
	}, nil
}

func (s SessionAccountDB) Find(sa *SessionAccount) (bool, error) {
	db := s.traits.GetGormDB().Model(sa).
		Where("session_id = ? and chain_id = ? and address = ?",
			sa.SessionID, sa.ChainID, sa.Address).Find(&sa)
	if db.RecordNotFound() {
		return false, nil
	} else if db.Error != nil {
		return false, db.Error
	}
	return true, nil
}

func (s SessionAccountDB) Migrate() error {
	if err := s.traits.Migrate(); err != nil {
		return err
	}

	return s.traits.GetGormDB().Model(new(SessionAccount)).
		AddUniqueIndex("sa_sca",
			"session_id", "chain_id", "address").Error
}

func (s SessionAccountDB) Create(sa *SessionAccount) (int64, error) {
	return s.traits.Create(sa)
}

func (s SessionAccountDB) UpdateAcknowledged(sa *SessionAccount) (int64, error) {
	db := s.traits.GetGormDB()
	db = db.Model(sa).Where("session_id = ? and chain_id = ? and address = ?", sa.SessionID, sa.ChainID, sa.Address).Update("acknowledged", sa.Acknowledged)
	return db.RowsAffected, db.Error
}

func (s SessionAccountDB) Delete(sa *SessionAccount) (int64, error) {
	return s.traits.Delete(sa)
}

func (s *SessionAccountDB) Filter(f *database.SessionAccountFilter) (user []SessionAccount, err error) {
	err = s.traits.Filter(f, &user)
	return
}

func (s *SessionAccountDB) FilterI(f interface{}) (interface{}, error) {
	return s.Filter(f.(*database.SessionAccountFilter))
}

func (s *SessionAccountDB) GetAcknowledged(guid string) (int64, error) {
	return s.acknowledged(guid)
}

func (s *SessionAccountDB) GetTotal(guid string) (int64, error) {
	return s.total(guid)
}

func (s *SessionAccountDB) ID(id string) (sessionAccount []SessionAccount, err error) {
	// todo wheres1
	return s.QueryChain().HasID(id).Query()
}

func (s *SessionAccountDB) ID_(db *gorm.DB, id string) (sessionAccount []SessionAccount, err error) {
	return (&SessionAccountQuery{db}).HasID(id).Query()
}

type SessionAccountQuery struct {
	db *gorm.DB
}

func (s *SessionAccountDB) QueryChain() *SessionAccountQuery {
	return &SessionAccountQuery{db: s.traits.GetGormDB()}
}

func (sessionAccountDB *SessionAccountQuery) Order(order string, reorder ...bool) *SessionAccountQuery {
	sessionAccountDB.db = sessionAccountDB.db.Order(order, reorder...)
	return sessionAccountDB
}

func (sessionAccountDB *SessionAccountQuery) Page(page, pageSize int) *SessionAccountQuery {
	sessionAccountDB.db = sessionAccountDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return sessionAccountDB
}

func (sessionAccountDB *SessionAccountQuery) HasID(id string) *SessionAccountQuery {
	sessionAccountDB.db = sessionAccountDB.db.Where("session_id = ?", id)
	return sessionAccountDB
}

func (sessionAccountDB *SessionAccountQuery) Preload() *SessionAccountQuery {
	sessionAccountDB.db = sessionAccountDB.db.Set("gorm:auto_preload", true)
	return sessionAccountDB
}

func (sessionAccountDB *SessionAccountQuery) Query() (sessionAccounts []SessionAccount, err error) {
	err = sessionAccountDB.db.Find(&sessionAccounts).Error
	return
}

func (sessionAccountDB *SessionAccountQuery) Scan(desc interface{}) (err error) {
	err = sessionAccountDB.db.Scan(desc).Error
	return
}
