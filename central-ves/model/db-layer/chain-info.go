package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	chainInfoTraits Traits
)

func injectChainInfoTraits() error {
	chainInfoTraits = NewTraits(ChainInfo{})
	return nil
}

func wrapToChainInfo(chainInfo interface{}, err error) (*ChainInfo, error) {
	if chainInfo == nil {
		return nil, err
	}
	return chainInfo.(*ChainInfo), err
}

type ChainInfo struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`
}

// TableName specification
func (ChainInfo) TableName() string {
	return "chain_info"
}

func (ChainInfo) migrate() error {
	return chainInfoTraits.Migrate()
}

func (d ChainInfo) GetID() uint {
	return d.ID
}

func (d *ChainInfo) Create() (int64, error) {
	return chainInfoTraits.Create(d)
}

func (d *ChainInfo) Update() (int64, error) {
	return chainInfoTraits.Update(d)
}

func (d *ChainInfo) UpdateFields(fields []string) (int64, error) {
	return chainInfoTraits.UpdateFields(d, fields)
}

func (d *ChainInfo) UpdateFields_(db *dorm.DB, fields []string) (int64, error) {
	return chainInfoTraits.UpdateFields_(db, d, fields)
}

func (d *ChainInfo) UpdateFields__(db dorm.SQLCommon, fields []string) (int64, error) {
	return chainInfoTraits.UpdateFields__(db, d, fields)
}

func (d *ChainInfo) Delete() (int64, error) {
	return chainInfoTraits.Delete(d)
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

func (chainInfoDB *ChainInfoDB) ID_(db *gorm.DB, id uint) (goods *ChainInfo, err error) {
	return wrapToChainInfo(chainInfoTraits.ID_(db, id))
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
