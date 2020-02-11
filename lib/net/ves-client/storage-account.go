package vesclient

import (
	"encoding/base64"
	base_account "github.com/HyperService-Consortium/go-uip/base-account"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/go-ves/lib/extend-traits"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	accountTraits     traits
	accountInvertFind extend_traits.Where2Func
)

func injectAccountTraits() error {
	accountTraits = newTraits(Account{})
	accountInvertFind = accountTraits.Where2("chain_id = ? and address = ?")
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

	Address string                         `dorm:"address" gorm:"address;not_null" json:"address"`
	ChainID uiptypes.ChainIDUnderlyingType `dorm:"chain_id" gorm:"chain_id;not_null" json:"chain_id"`
}

// TableName specification
func (Account) TableName() string {
	return "chain_info"
}

func (ci Account) migrate() error {
	if err := accountTraits.Migrate(); err != nil {
		return err
	}

	return p.GormDB.Model(&ci).AddUniqueIndex("ci_ac", "address", "chain_id").Error
}

func (ci Account) GetID() uint {
	return ci.ID
}

func (ci *Account) Create() (int64, error) {
	return accountTraits.Create(ci)
}

func (ci *Account) Update() (int64, error) {
	return accountTraits.Update(ci)
}

func (ci *Account) UpdateFields(fields []string) (int64, error) {
	return accountTraits.UpdateFields(ci, fields)
}

func (ci *Account) UpdateFields_(db *dorm.DB, fields []string) (int64, error) {
	return accountTraits.UpdateFields_(db, ci, fields)
}

func (ci *Account) UpdateFields__(db dorm.SQLCommon, fields []string) (int64, error) {
	return accountTraits.UpdateFields__(db, ci, fields)
}

func (ci *Account) Delete() (int64, error) {
	return accountTraits.Delete(ci)
}

type AccountDB struct{}

func NewAccountDB(_ module.Module) (*AccountDB, error) {
	return new(AccountDB), nil
}

func GetAccountDB(_ module.Module) (*AccountDB, error) {
	return new(AccountDB), nil
}

func (accountDB *AccountDB) ID(id uint) (account *Account, err error) {
	return wrapToAccount(accountTraits.ID(id))
}

func (accountDB *AccountDB) ID_(db *gorm.DB, id uint) (account *Account, err error) {
	return wrapToAccount(accountTraits.ID_(db, id))
}

func (accountDB *AccountDB) InvertFind(acc uiptypes.Account) (account *Account, err error) {
	return wrapToAccount(accountInvertFind(acc.GetChainId(), acc.GetAddress()))
}

func decodeBase64(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

func (accountDB *AccountDB) FindAccounts(id uint, chainID uiptypes.ChainIDUnderlyingType) ([]uiptypes.Account, error) {
	var mid []string
	var err = p.GormDB.Where("id = ? and chain_id = ?", id, chainID).
		Select("address").
		Scan(&mid).Error
	if err != nil {
		return nil, err
	}
	var results []uiptypes.Account
	for i := range mid {
		add, err := decodeBase64(mid[i])
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
	return &AccountQuery{db: p.GormDB}
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
