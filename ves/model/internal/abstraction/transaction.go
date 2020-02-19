package abstraction

import (
	"github.com/Myriad-Dreamin/dorm"
	database2 "github.com/Myriad-Dreamin/go-ves/ves/model/internal/database"
	"github.com/jinzhu/gorm"
)

// TransactionDB ...
type TransactionDB interface {
	GetTraits() ORMTraits
	Create(d *database2.Transaction) (aff int64, err error)
	Update(d *database2.Transaction) (aff int64, err error)
	UpdateFields(d *database2.Transaction, fields []string) (aff int64, err error)
	UpdateFields_(d *database2.Transaction, db *dorm.DB, fields []string) (aff int64, err error)
	UpdateFields__(d *database2.Transaction, db dorm.SQLCommon, fields []string) (aff int64, err error)
	Delete(d *database2.Transaction) (aff int64, err error)
	Filter(f *database2.Filter) (user []database2.Transaction, err error)
	FilterI(f interface{}) (interface{}, error)
	ID(id uint) (transaction *database2.Transaction, err error)
	ID_(db *gorm.DB, id uint) (transaction *database2.Transaction, err error)
	FindSessionIndex(d *database2.Transaction, sessionID string, idx int64) (bool, error)

	Query(opts ...TransactionQueryOption) (objs []database2.Transaction, err error)
	Scan(desc interface{}, opts ...TransactionQueryOption) (err error)
}

type TransactionQueryOption interface {
	implementsTransactionQuery() TransactionQueryOption
}
