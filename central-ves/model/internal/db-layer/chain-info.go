package dblayer

import (
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/go-ves/lib/backend/extend-traits"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

func wrapToChainInfo(chainInfo interface{}, err error) (*ChainInfo, error) {
	if chainInfo == nil {
		return nil, err
	}
	return chainInfo.(*ChainInfo), err
}

type ChainInfoDB struct {
	traits     abstraction.ORMTraits
	invertFind extend_traits.Where2Func
}

func (c ChainInfoDB) GetTraits() abstraction.ORMTraits {
	return c.traits
}

func NewChainInfoDB(cb func(interface{}) abstraction.ORMTraits, _ module.Module) (*ChainInfoDB, error) {
	traits := cb(ChainInfo{})
	return &ChainInfoDB{
		traits:     traits,
		invertFind: traits.Where2("chain_id = ? and address = ?"),
	}, nil
}

func (c ChainInfoDB) Migrate() error {
	if err := c.traits.Migrate(); err != nil {
		return err
	}

	return c.traits.GetGormDB().Model(new(ChainInfo)).
		AddUniqueIndex("ci_ac", "address", "chain_id").Error
}

func (c ChainInfoDB) Create(ci *ChainInfo) (aff int64, err error) {
	return c.traits.Create(ci)
}

func (c ChainInfoDB) Update(ci *ChainInfo) (aff int64, err error) {
	return c.traits.Update(ci)
}

func (c ChainInfoDB) UpdateFields(ci *ChainInfo, fields []string) (aff int64, err error) {
	return c.traits.UpdateFields(ci, fields)
}

func (c ChainInfoDB) UpdateFields_(ci *ChainInfo, db *dorm.DB, fields []string) (aff int64, err error) {
	return c.traits.UpdateFields_(db, ci, fields)
}

func (c ChainInfoDB) UpdateFields__(ci *ChainInfo, db dorm.SQLCommon, fields []string) (aff int64, err error) {
	return c.traits.UpdateFields__(db, ci, fields)
}

func (c ChainInfoDB) Delete(ci *ChainInfo, ) (aff int64, err error) {
	return c.traits.Delete(ci)
}

func (c ChainInfoDB) Filter(f *ChainInfoFilter) (user []ChainInfo, err error) {
	err = c.traits.Filter(f, &user)
	return
}

func (c ChainInfoDB) ID(id uint) (chainInfo *ChainInfo, err error) {
	return wrapToChainInfo(c.traits.ID(id))
}

func (c ChainInfoDB) ID_(db *gorm.DB, id uint) (chainInfo *ChainInfo, err error) {
	return wrapToChainInfo(c.traits.ID_(db, id))
}

func (c ChainInfoDB) InvertFind(acc uip.Account) (chainInfo *ChainInfo, err error) {
	return wrapToChainInfo(c.invertFind(acc.GetChainId(), encodeAddress(acc.GetAddress())))
}

func (c ChainInfoDB) FindAccounts(id uint, chainID uip.ChainIDUnderlyingType) ([]uip.Account, error) {
	var mid []string
	var err = c.traits.GetGormDB().Where("id = ? and chain_id = ?", id, chainID).
		Select("address").
		Scan(&mid).Error
	if err != nil {
		return nil, err
	}
	var results []uip.Account
	for i := range mid {
		add, err := decodeAddress(mid[i])
		if err != nil {
			return nil, err
		}
		results = append(results, &uip.AccountImpl{
			ChainId: chainID,
			Address: add,
		})
	}
	return results, nil
}

func (c ChainInfoDB) FilterI(f interface{}) (interface{}, error) {
	return c.Filter(f.(*ChainInfoFilter))
}

func (c ChainInfoDB) Query(options ...abstraction.ChainInfoQueryOption) (cis []ChainInfo, err error) {
	panic("implement me")
}

func (c ChainInfoDB) Scan(desc interface{}, options ...abstraction.ChainInfoQueryOption) (err error) {
	panic("implement me")
}

func (c *ChainInfoDB) QueryChain() *ChainInfoQuery {
	return &ChainInfoQuery{db: c.traits.GetGormDB()}
}

type ChainInfoQuery struct {
	db *gorm.DB
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
