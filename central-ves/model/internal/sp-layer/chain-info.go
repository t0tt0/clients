package splayer

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)


type ChainInfoDB struct {
	abstraction.ChainInfoDB
}

func NewChainInfoDB(base abstraction.ChainInfoDB, _ module.Module) (*ChainInfoDB, error) {
	db := new(ChainInfoDB)
	db.ChainInfoDB = base
	return db, nil
}
