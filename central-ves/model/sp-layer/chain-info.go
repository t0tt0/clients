package splayer

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	dblayer "github.com/Myriad-Dreamin/go-ves/central-ves/model/db-layer"
)

type ChainInfo = dblayer.ChainInfo

type ChainInfoDB struct {
	dblayer.ChainInfoDB
}

func NewChainInfoDB(m module.Module) (*ChainInfoDB, error) {
	cdb, err := dblayer.NewChainInfoDB(m)
	if err != nil {
		return nil, err
	}
	db := new(ChainInfoDB)
	db.ChainInfoDB = *cdb
	return db, nil
}

func GetChainInfoDB(m module.Module) (*ChainInfoDB, error) {
	cdb, err := dblayer.GetChainInfoDB(m)
	if err != nil {
		return nil, err
	}
	db := new(ChainInfoDB)
	db.ChainInfoDB = *cdb
	return db, nil
}

func (s *Provider) ChainInfoDB() *ChainInfoDB {
	return s.chainInfoDB
}
