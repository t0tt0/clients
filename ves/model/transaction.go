package model

import (
	splayer "github.com/Myriad-Dreamin/go-ves/ves/model/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Transaction = splayer.Transaction
type TransactionDB = splayer.TransactionDB

func NewTransactionDB(m module.Module) (*TransactionDB, error) {
	return splayer.NewTransactionDB(m)
}

func GetTransactionDB(m module.Module) (*TransactionDB, error) {
	return splayer.GetTransactionDB(m)
}
