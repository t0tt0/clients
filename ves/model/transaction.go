package model

import (
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/database"
	dblayer "github.com/Myriad-Dreamin/go-ves/ves/model/internal/db-layer"
	splayer "github.com/Myriad-Dreamin/go-ves/ves/model/internal/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Transaction = database.Transaction
type TransactionDB = abstraction.TransactionDB

func (p DBLayerModule) NewTransactionDB(m module.Module) (TransactionDB, error) {
	return dblayer.NewTransactionDB(p.newTraits, m)
}

func (p SPLayerModule) NewTransactionDB(base abstraction.TransactionDB, m module.Module) (TransactionDB, error) {
	return splayer.NewTransactionDB(base, m)
}

func (p Module) NewTransactionDB(m module.Module) (TransactionDB, error) {
	base, err := p.dbLayer.NewTransactionDB(m)
	if err != nil {
		return nil, err
	}
	return p.spLayer.NewTransactionDB(base, m)
}

func NewTransactionDB(m module.Module) (TransactionDB, error) {
	return p.NewTransactionDB(m)
}
