package splayer

import (
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type TransactionDB struct {
	abstraction.TransactionDB
}

func NewTransactionDB(base abstraction.TransactionDB, m module.Module) (*TransactionDB, error) {
	db := new(TransactionDB)
	db.TransactionDB = base
	return db, nil
}
