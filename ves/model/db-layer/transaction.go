package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	transactionTraits Traits
)

func injectTransactionTraits() error {
	transactionTraits = NewTraits(Transaction{})
	return nil
}

func wrapToTransaction(transaction interface{}, err error) (*Transaction, error) {
	if transaction == nil {
		return nil, err
	}
	return transaction.(*Transaction), err
}

type Transaction struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null" json:"id"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`

	SessionID string `dorm:"session_id" gorm:"column:session_id;not null;" json:"session_id"`
	Index     int64    `dorm:"index" gorm:"column:index;not null;" json:"index"`
	Content   string `dorm:"contents" gorm:"column:content;not null;" json:"content"`
}

// TableName specification
func (Transaction) TableName() string {
	return "transaction"
}

func (Transaction) migrate() error {
	return transactionTraits.Migrate()
}

func (d Transaction) GetID() uint {
	return d.ID
}

func findObject(g *gorm.DB, d interface{}) (has bool, err error) {
	db := g.Find(&d)
	if db.RecordNotFound() {
		return false, nil
	} else if db.Error != nil {
		return false, db.Error
	}
	return true, nil
}

func (d *Transaction) FindSessionIndex(sessionID string, index int64) (bool, error) {
	return findObject(p.GormDB.Where("session_id = ? and index = ?", sessionID, index), d)
}

func (d *Transaction) Create() (int64, error) {
	return transactionTraits.Create(d)
}

func (d *Transaction) Update() (int64, error) {
	return transactionTraits.Update(d)
}

func (d *Transaction) UpdateFields(fields []string) (int64, error) {
	return transactionTraits.UpdateFields(d, fields)
}

func (d *Transaction) UpdateFields_(db *dorm.DB, fields []string) (int64, error) {
	return transactionTraits.UpdateFields_(db, d, fields)
}

func (d *Transaction) UpdateFields__(db dorm.SQLCommon, fields []string) (int64, error) {
	return transactionTraits.UpdateFields__(db, d, fields)
}

func (d *Transaction) Delete() (int64, error) {
	return transactionTraits.Delete(d)
}

type TransactionDB struct{}

func NewTransactionDB(_ module.Module) (*TransactionDB, error) {
	return new(TransactionDB), nil
}

func GetTransactionDB(_ module.Module) (*TransactionDB, error) {
	return new(TransactionDB), nil
}

func (transactionDB *TransactionDB) Filter(f *Filter) (user []Transaction, err error) {
	err = transactionTraits.Filter(f, &user)
	return
}

func (transactionDB *TransactionDB) FilterI(f interface{}) (interface{}, error) {
	return transactionDB.Filter(f.(*Filter))
}

func (transactionDB *TransactionDB) ID(id uint) (transaction *Transaction, err error) {
	return wrapToTransaction(transactionTraits.ID(id))
}

func (transactionDB *TransactionDB) ID_(db *gorm.DB, id uint) (goods *Transaction, err error) {
	return wrapToTransaction(transactionTraits.ID_(db, id))
}

type TransactionQuery struct {
	db *gorm.DB
}

func (transactionDB *TransactionDB) QueryChain() *TransactionQuery {
	return &TransactionQuery{db: p.GormDB}
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

func (transactionDB *TransactionQuery) Query() (transactions []Transaction, err error) {
	err = transactionDB.db.Find(&transactions).Error
	return
}

func (transactionDB *TransactionQuery) Scan(desc interface{}) (err error) {
	err = transactionDB.db.Scan(desc).Error
	return
}
