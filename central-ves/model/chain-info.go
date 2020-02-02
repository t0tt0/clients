package model

import (
	splayer "github.com/Myriad-Dreamin/go-ves/central-ves/model/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type ChainInfo = splayer.ChainInfo
type ChainInfoDB = splayer.ChainInfoDB

func NewChainInfoDB(m module.Module) (*ChainInfoDB, error) {
	return splayer.NewChainInfoDB(m)
}

func GetChainInfoDB(m module.Module) (*ChainInfoDB, error) {
	return splayer.GetChainInfoDB(m)
}
