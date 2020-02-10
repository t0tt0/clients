package dblayer

import (
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/lib/base64"
	extend_traits "github.com/Myriad-Dreamin/go-ves/lib/extend-traits"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

var (
	sessionAccountTraits       Traits
	sessionAccountQueryID      extend_traits.Where1Func
	_sessionAccountQueryID     extend_traits.Where1Func_
	sessionAccountAcknowledged extend_traits.Count1Func
	sessionAccountTotal        extend_traits.Count1Func
)

func injectSessionAccountTraits() error {
	sessionAccountTraits = NewTraits(SessionAccount{})

	sessionAccountQueryID = sessionAccountTraits.Where1("session_id = ?")
	_sessionAccountQueryID = sessionAccountTraits.Where1_("session_id = ?")
	sessionAccountAcknowledged = sessionAccountTraits.Count1("session_id = ? and acknowledged = 1")
	sessionAccountTotal = sessionAccountTraits.Count1("session_id = ?")
	return nil
}

func wrapToSessionAccount(sessionAccount interface{}, err error) (*SessionAccount, error) {
	if sessionAccount == nil {
		return nil, err
	}
	return sessionAccount.(*SessionAccount), err
}

type SessionAccount struct {
	SessionID    string                         `dorm:"session_id" gorm:"column:session_id;not_null" json:"-"`
	ChainID      uiptypes.ChainIDUnderlyingType `dorm:"chain_id" gorm:"column:chain_id;not_null" json:"chain_id"`
	Address      string                         `dorm:"address" gorm:"column:address;not_null" json:"address"`
	Acknowledged bool                           `dorm:"acknowledged" gorm:"column:acknowledged;not_null" json:"acknowledged"`

	decodedAddress []byte `gorm:"-" json:"-"`
}

func NewSessionAccount(chainID uiptypes.ChainIDUnderlyingType, address []byte) *SessionAccount {
	return &SessionAccount{
		ChainID:        chainID,
		Address:        base64.EncodeBase64(address),
		decodedAddress: address,
	}
}

func (sa SessionAccount) Find() (bool, error) {
	db := p.GormDB.Find(&sa)
	if db.RecordNotFound() {
		return false, nil
	} else if db.Error != nil {
		return false, db.Error
	}
	return true, nil
}

func (sa SessionAccount) GetChainId() uiptypes.ChainID {
	return sa.ChainID
}

func (sa SessionAccount) GetAddress() uiptypes.Address {
	if sa.decodedAddress == nil {
		sa.decodedAddress = decodeBase64(sa.Address)
	}
	return sa.decodedAddress
}

// TableName specification
func (SessionAccount) TableName() string {
	return "session_account"
}

func (sa SessionAccount) GetID() uint {
	panic("aborted")
}

func (sa SessionAccount) migrate() error {
	if err := sessionAccountTraits.Migrate(); err != nil {
		return err
	}

	return p.GormDB.Model(&sa).
		AddUniqueIndex("sa_sca",
			"session_id", "chain_id", "address").Error
}

func (sa *SessionAccount) Create() (int64, error) {
	return sessionAccountTraits.Create(sa)
}

func (sa *SessionAccount) Update() (int64, error) {
	return sessionAccountTraits.Update(sa)
}

func (sa *SessionAccount) Delete() (int64, error) {
	return sessionAccountTraits.Delete(sa)
}

type SessionAccountDB struct{}

func NewSessionAccountDB(_ module.Module) (*SessionAccountDB, error) {
	return new(SessionAccountDB), nil
}

func GetSessionAccountDB(_ module.Module) (*SessionAccountDB, error) {
	return new(SessionAccountDB), nil
}

func (sessionAccountDB *SessionAccountDB) Filter(f *Filter) (user []SessionAccount, err error) {
	err = sessionAccountTraits.Filter(f, &user)
	return
}

func (sessionAccountDB *SessionAccountDB) FilterI(f interface{}) (interface{}, error) {
	return sessionAccountDB.Filter(f.(*Filter))
}

func (sessionAccountDB *SessionAccountDB) GetAcknowledged(guid string) (int64, error) {
	return sessionAccountAcknowledged(guid)
}

func (sessionAccountDB *SessionAccountDB) GetTotal(guid string) (int64, error) {
	return sessionAccountTotal(guid)
}

func (sessionAccountDB *SessionAccountDB) ID(id string) (sessionAccount []SessionAccount, err error) {
	// todo wheres1
	return sessionAccountDB.QueryChain().HasID(id).Query()
}

func (sessionAccountDB *SessionAccountDB) ID_(db *gorm.DB, id string) (goods *SessionAccount, err error) {
	return wrapToSessionAccount(_sessionAccountQueryID(db, id))
}

type SessionAccountQuery struct {
	db *gorm.DB
}

func (sessionAccountDB *SessionAccountDB) QueryChain() *SessionAccountQuery {
	return &SessionAccountQuery{db: p.GormDB}
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
