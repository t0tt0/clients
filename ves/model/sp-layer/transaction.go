package splayer

import (
	dblayer "github.com/Myriad-Dreamin/go-ves/ves/model/db-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Transaction = dblayer.Transaction

type TransactionDB struct {
	dblayer.TransactionDB
}

func NewTransactionDB(m module.Module) (*TransactionDB, error) {
	cdb, err := dblayer.NewTransactionDB(m)
	if err != nil {
		return nil, err
	}
	db := new(TransactionDB)
	db.TransactionDB = *cdb
	return db, nil
}

func GetTransactionDB(m module.Module) (*TransactionDB, error) {
	cdb, err := dblayer.GetTransactionDB(m)
	if err != nil {
		return nil, err
	}
	db := new(TransactionDB)
	db.TransactionDB = *cdb
	return db, nil
}

func (s *Provider) TransactionDB() *TransactionDB {
	return s.transactionDB
}
