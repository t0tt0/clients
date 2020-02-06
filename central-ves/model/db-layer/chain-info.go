package dblayer

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
	chainInfoTraits     Traits
	chainInfoInvertFind extend_traits.Where2Func
)

func injectChainInfoTraits() error {
	chainInfoTraits = NewTraits(ChainInfo{})
	chainInfoInvertFind = chainInfoTraits.Where2("chain_id = ? and address = ?")
	return nil
}

func wrapToChainInfo(chainInfo interface{}, err error) (*ChainInfo, error) {
	if chainInfo == nil {
		return nil, err
	}
	return chainInfo.(*ChainInfo), err
}

type ChainInfo struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null" json:"id"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`

	UserID  uint                           `dorm:"user_id" gorm:"column:user_id;not_null" json:"user_id"`
	Address string                         `dorm:"address" gorm:"address;not_null" json:"address"`
	ChainID uiptypes.ChainIDUnderlyingType `dorm:"chain_id" gorm:"chain_id;not_null" json:"chain_id"`
}

// TableName specification
func (ChainInfo) TableName() string {
	return "chain_info"
}

func (ci ChainInfo) migrate() error {
	if err := chainInfoTraits.Migrate(); err != nil {
		return err
	}

	return p.GormDB.Model(&ci).AddUniqueIndex("ci_ac", "address", "chain_id").Error
}

func (ci ChainInfo) GetID() uint {
	return ci.ID
}

func (ci *ChainInfo) Create() (int64, error) {
	return chainInfoTraits.Create(ci)
}

func (ci *ChainInfo) Update() (int64, error) {
	return chainInfoTraits.Update(ci)
}

func (ci *ChainInfo) UpdateFields(fields []string) (int64, error) {
	return chainInfoTraits.UpdateFields(ci, fields)
}

func (ci *ChainInfo) UpdateFields_(db *dorm.DB, fields []string) (int64, error) {
	return chainInfoTraits.UpdateFields_(db, ci, fields)
}

func (ci *ChainInfo) UpdateFields__(db dorm.SQLCommon, fields []string) (int64, error) {
	return chainInfoTraits.UpdateFields__(db, ci, fields)
}

func (ci *ChainInfo) Delete() (int64, error) {
	return chainInfoTraits.Delete(ci)
}

type ChainInfoDB struct{}

func NewChainInfoDB(_ module.Module) (*ChainInfoDB, error) {
	return new(ChainInfoDB), nil
}

func GetChainInfoDB(_ module.Module) (*ChainInfoDB, error) {
	return new(ChainInfoDB), nil
}

func (chainInfoDB *ChainInfoDB) Filter(f *Filter) (user []ChainInfo, err error) {
	err = chainInfoTraits.Filter(f, &user)
	return
}

func (chainInfoDB *ChainInfoDB) FilterI(f interface{}) (interface{}, error) {
	return chainInfoDB.Filter(f.(*Filter))
}

func (chainInfoDB *ChainInfoDB) ID(id uint) (chainInfo *ChainInfo, err error) {
	return wrapToChainInfo(chainInfoTraits.ID(id))
}

func (chainInfoDB *ChainInfoDB) ID_(db *gorm.DB, id uint) (chainInfo *ChainInfo, err error) {
	return wrapToChainInfo(chainInfoTraits.ID_(db, id))
}

func (chainInfoDB *ChainInfoDB) InvertFind(acc uiptypes.Account) (chainInfo *ChainInfo, err error) {
	return wrapToChainInfo(chainInfoInvertFind(acc.GetChainId(), acc.GetAddress()))
}

func decodeBase64(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

func (chainInfoDB *ChainInfoDB) FindAccounts(id uint, chainID uiptypes.ChainIDUnderlyingType) ([]uiptypes.Account, error) {
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

type ChainInfoQuery struct {
	db *gorm.DB
}

func (chainInfoDB *ChainInfoDB) QueryChain() *ChainInfoQuery {
	return &ChainInfoQuery{db: p.GormDB}
}

func (chainInfoDB *ChainInfoQuery) Order(order string, reorder ...bool) *ChainInfoQuery {
	chainInfoDB.db = chainInfoDB.db.Order(order, reorder...)
	return chainInfoDB
}

func (chainInfoDB *ChainInfoQuery) Page(page, pageSize int) *ChainInfoQuery {
	chainInfoDB.db = chainInfoDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return chainInfoDB
}
func (chainInfoDB *ChainInfoQuery) BeforeID(id uint) *ChainInfoQuery {
	chainInfoDB.db = chainInfoDB.db.Where("id <= ?", id)
	return chainInfoDB
}

func (chainInfoDB *ChainInfoQuery) Preload() *ChainInfoQuery {
	chainInfoDB.db = chainInfoDB.db.Set("gorm:auto_preload", true)
	return chainInfoDB
}

func (chainInfoDB *ChainInfoQuery) Query() (chainInfos []ChainInfo, err error) {
	err = chainInfoDB.db.Find(&chainInfos).Error
	return
}

func (chainInfoDB *ChainInfoQuery) Scan(desc interface{}) (err error) {
	err = chainInfoDB.db.Scan(desc).Error
	return
}
