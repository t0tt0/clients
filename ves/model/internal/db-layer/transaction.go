package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/abstraction"
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/database"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

func wrapToTransaction(transaction interface{}, err error) (*database.Transaction, error) {
	if transaction == nil {
		return nil, err
	}
	return transaction.(*database.Transaction), err
}

type TransactionDB struct {
	traits abstraction.ORMTraits
}

func NewTransactionDB(cb func(interface{}) abstraction.ORMTraits, _ module.Module) (*TransactionDB, error) {
	return &TransactionDB{traits: cb(Transaction{})}, nil
}

func (transactionDB TransactionDB) GetTraits() abstraction.ORMTraits {
	return transactionDB.traits
}

func (transactionDB TransactionDB) Query(options ...abstraction.TransactionQueryOption) (objs []database.Transaction, err error) {
	panic("implement me")
}

func (transactionDB TransactionDB) Scan(desc interface{}, options ...abstraction.TransactionQueryOption) (err error) {
	panic("implement me")
}

func (transactionDB TransactionDB) Migrate() error {
	return transactionDB.traits.Migrate()
}

func (transactionDB TransactionDB) Create(d *database.Transaction) (aff int64, err error) {
	return transactionDB.traits.Create(d)
}

func (transactionDB TransactionDB) Update(d *database.Transaction) (aff int64, err error) {
	return transactionDB.traits.Update(d)
}

func (transactionDB TransactionDB) UpdateFields(d *database.Transaction, fields []string) (aff int64, err error) {
	return transactionDB.traits.UpdateFields(d, fields)
}

func (transactionDB TransactionDB) UpdateFields_(d *database.Transaction, db *dorm.DB, fields []string) (aff int64, err error) {
	return transactionDB.traits.UpdateFields_(db, d, fields)
}

func (transactionDB TransactionDB) UpdateFields__(d *database.Transaction, db dorm.SQLCommon, fields []string) (aff int64, err error) {
	return transactionDB.traits.UpdateFields__(db, d, fields)
}

func (transactionDB TransactionDB) Delete(d *database.Transaction) (aff int64, err error) {
	return transactionDB.traits.Delete(d)
}

func (transactionDB *TransactionDB) Filter(f *database.Filter) (user []database.Transaction, err error) {
	err = transactionDB.traits.Filter(f, &user)
	return
}

func (transactionDB *TransactionDB) FilterI(f interface{}) (interface{}, error) {
	return transactionDB.Filter(f.(*database.Filter))
}

func (transactionDB *TransactionDB) ID(id uint) (transaction *database.Transaction, err error) {
	return wrapToTransaction(transactionDB.traits.ID(id))
}

func (transactionDB *TransactionDB) ID_(db *gorm.DB, id uint) (transaction *database.Transaction, err error) {
	return wrapToTransaction(transactionDB.traits.ID_(db, id))
}

type TransactionQuery struct {
	db *gorm.DB
}

func findTransaction(g *gorm.DB, d interface{}) (has bool, err error) {
	db := g.Find(d)
	if db.RecordNotFound() {
		return false, nil
	} else if db.Error != nil {
		return false, db.Error
	}
	return true, nil
}

func (transactionDB *TransactionDB) FindSessionIndex(d *Transaction, sessionID string, idx int64) (bool, error) {
	return findTransaction(transactionDB.traits.GetGormDB().Where("session_id = ? and idx = ?", sessionID, idx), d)
}

func (transactionDB *TransactionDB) QueryChain() *TransactionQuery {
	return &TransactionQuery{db: transactionDB.traits.GetGormDB()}
}

func (transactionDB *TransactionQuery) Order(order string, reorder ...bool) *TransactionQuery {
	transactionDB.db = transactionDB.db.Order(order, reorder...)
	return transactionDB
}

func (transactionDB *TransactionQuery) Page(page, pageSize int) *TransactionQuery {
	transactionDB.db = transactionDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return transactionDB
}
func (transactionDB *TransactionQuery) BeforeID(id uint) *TransactionQuery {
	transactionDB.db = transactionDB.db.Where("id <= ?", id)
	return transactionDB
}

func (transactionDB *TransactionQuery) Preload() *TransactionQuery {
	transactionDB.db = transactionDB.db.Set("gorm:auto_preload", true)
	return transactionDB
}

func (transactionDB *TransactionQuery) Query() (transactions []database.Transaction, err error) {
	err = transactionDB.db.Find(&transactions).Error
	return
}

func (transactionDB *TransactionQuery) Scan(desc interface{}) (err error) {
	err = transactionDB.db.Scan(desc).Error
	return
}
