package vesclient

import (
	base_account "github.com/HyperService-Consortium/go-uip/base-account"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/go-ves/lib/encoding"
	"github.com/Myriad-Dreamin/go-ves/lib/extend-traits"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
	"time"
)

var ()

type accountTraits struct {
	traits
	accountInvertFind extend_traits.Where2Func
	accountQueryAlias extend_traits.Where1Func
}

func (m *modelModule) injectAccountTraits() error {
	m.accountTraits.traits = m.newTraits(Account{})
	m.accountTraits.accountInvertFind = m.accountTraits.Where2("chain_id = ? and address = ?")
	m.accountTraits.accountQueryAlias = m.accountTraits.Where1("alias = ?")
	return nil
}

func wrapToAccount(account interface{}, err error) (*Account, error) {
	if account == nil {
		return nil, err
	}
	return account.(*Account), err
}

type Account struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null" json:"id"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`

	Alias     string                           `dorm:"alias" gorm:"alias;not_null" json:"alias"`
	Address   string                           `dorm:"address" gorm:"address;not_null" json:"address"`
	Addition  string                           `dorm:"addition" gorm:"addition;not_null" json:"addition"`
	ChainType uiptypes.ChainTypeUnderlyingType `dorm:"chain_type" gorm:"chain_type;not_null" json:"chain_type"`
	ChainID   uiptypes.ChainIDUnderlyingType   `dorm:"chain_id" gorm:"chain_id;not_null" json:"chain_id"`
}

// TableName specification
func (Account) TableName() string {
	return "account"
}

func NewAccount() *Account {
	return &Account{}
}

func (ci Account) migrate(dep *modelModule) error {
	if err := dep.accountTraits.Migrate(); err != nil {
		return err
	}

	return dep.GormDB.Model(&ci).AddUniqueIndex("ci_ac", "address", "chain_id").Error
}

func (ci Account) migration(dep *modelModule) func() error {
	return func() error {
		return ci.migrate(dep)
	}
}

func (ci Account) GetID() uint {
	return ci.ID
}

func (accountDB AccountDB) Create(ci *Account) (int64, error) {
	return accountDB.module.accountTraits.Create(ci)
}

func (accountDB AccountDB) Update(ci *Account) (int64, error) {
	return accountDB.module.accountTraits.Update(ci)
}

func (accountDB AccountDB) UpdateFields(ci *Account, fields []string) (int64, error) {
	return accountDB.module.accountTraits.UpdateFields(ci, fields)
}

func (accountDB AccountDB) UpdateFields_(ci *Account, db *dorm.DB, fields []string) (int64, error) {
	return accountDB.module.accountTraits.UpdateFields_(db, ci, fields)
}

func (accountDB AccountDB) UpdateFields__(ci *Account, db dorm.SQLCommon, fields []string) (int64, error) {
	return accountDB.module.accountTraits.UpdateFields__(db, ci, fields)
}

func (accountDB AccountDB) Delete(ci *Account) (int64, error) {
	return accountDB.Delete(ci)
}

type AccountDB struct {
	db     *gorm.DB
	module *modelModule
}

func NewAccountDB(m Module) (*AccountDB, error) {
	return &AccountDB{db: m.GormDB(), module: m.ModelModule()}, nil
}

func GetAccountDB(_ module.Module) (*AccountDB, error) {
	return new(AccountDB), nil
}

func (accountDB *AccountDB) ID(id uint) (account *Account, err error) {
	return wrapToAccount(accountDB.module.accountTraits.ID(id))
}

func (accountDB *AccountDB) ID_(db *gorm.DB, id uint) (account *Account, err error) {
	return wrapToAccount(accountDB.module.accountTraits.ID_(db, id))
}

func (accountDB *AccountDB) InvertFind(acc uiptypes.Account) (account *Account, err error) {
	return wrapToAccount(accountDB.module.accountTraits.accountInvertFind(
		acc.GetChainId(), encodeAddress(acc.GetAddress())))
}

func (accountDB *AccountDB) QueryAlias(alias string) (account *Account, err error) {
	return wrapToAccount(accountDB.module.accountTraits.accountQueryAlias(alias))
}

func (accountDB *AccountDB) FindAccounts(id uint, chainID uiptypes.ChainIDUnderlyingType) ([]uiptypes.Account, error) {
	var mid []string
	var err = accountDB.db.Where("id = ? and chain_id = ?", id, chainID).
		Select("address").
		Scan(&mid).Error
	if err != nil {
		return nil, err
	}
	var results []uiptypes.Account
	for i := range mid {
		add, err := encoding.DecodeBase64(mid[i])
		if err != nil {
			return nil, err
		}
		results = append(results, &base_account.Account{
			ChainId: chainID,
			Address: add,
		})
	}
	return results, nil
}

type AccountQuery struct {
	db *gorm.DB
}

func (accountDB *AccountDB) QueryChain() *AccountQuery {
	return &AccountQuery{db: accountDB.db}
}

func (accountDB *AccountQuery) Order(order string, reorder ...bool) *AccountQuery {
	accountDB.db = accountDB.db.Order(order, reorder...)
	return accountDB
}

func (accountDB *AccountQuery) Page(page, pageSize int) *AccountQuery {
	accountDB.db = accountDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return accountDB
}
func (accountDB *AccountQuery) BeforeID(id uint) *AccountQuery {
	accountDB.db = accountDB.db.Where("id <= ?", id)
	return accountDB
}

func (accountDB *AccountQuery) Preload() *AccountQuery {
	accountDB.db = accountDB.db.Set("gorm:auto_preload", true)
	return accountDB
}

func (accountDB *AccountQuery) Query() (accounts []Account, err error) {
	err = accountDB.db.Find(&accounts).Error
	return
}

func (accountDB *AccountQuery) Scan(desc interface{}) (err error) {
	err = accountDB.db.Scan(desc).Error
	return
}

type AccountDBInterface interface {
	Create(ci *Account) (int64, error)
	Update(ci *Account) (int64, error)
	UpdateFields(ci *Account, fields []string) (int64, error)
	UpdateFields_(ci *Account, db *dorm.DB, fields []string) (int64, error)
	UpdateFields__(ci *Account, db dorm.SQLCommon, fields []string) (int64, error)
	Delete(ci *Account) (int64, error)
	ID(id uint) (account *Account, err error)
	ID_(db *gorm.DB, id uint) (account *Account, err error)
	InvertFind(acc uiptypes.Account) (account *Account, err error)
	QueryAlias(alias string) (account *Account, err error)
	FindAccounts(id uint, chainID uiptypes.ChainIDUnderlyingType) ([]uiptypes.Account, error)
	QueryChain() *AccountQuery
}
