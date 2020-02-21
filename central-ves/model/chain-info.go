package model

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/model/internal/abstraction"
	"github.com/HyperService-Consortium/go-ves/central-ves/model/internal/database"
	dblayer "github.com/HyperService-Consortium/go-ves/central-ves/model/internal/db-layer"
	splayer "github.com/HyperService-Consortium/go-ves/central-ves/model/internal/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type ChainInfo = database.ChainInfo
type ChainInfoDB = abstraction.ChainInfoDB

func (p DBLayerModule) NewChainInfoDB(m module.Module) (ChainInfoDB, error) {
	return dblayer.NewChainInfoDB(p.newTraits, m)
}

func (p SPLayerModule) NewChainInfoDB(base abstraction.ChainInfoDB, m module.Module) (ChainInfoDB, error) {
	return splayer.NewChainInfoDB(base, m)
}

func (p Module) NewChainInfoDB(m module.Module) (ChainInfoDB, error) {
	base, err := p.dbLayer.NewChainInfoDB(m)
	if err != nil {
		return nil, err
	}
	return p.spLayer.NewChainInfoDB(base, m)
}

func NewChainInfoDB(m module.Module) (ChainInfoDB, error) {
	return p.NewChainInfoDB(m)
}
