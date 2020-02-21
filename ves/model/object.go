package model

import (
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/database"
	dblayer "github.com/Myriad-Dreamin/go-ves/ves/model/internal/db-layer"
	splayer "github.com/Myriad-Dreamin/go-ves/ves/model/internal/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Object = database.Object
type ObjectDB = abstraction.ObjectDB

func (p DBLayerModule) NewObjectDB(m module.Module) (ObjectDB, error) {
	return dblayer.NewObjectDB(p.newTraits, m)
}

func (p SPLayerModule) NewObjectDB(base abstraction.ObjectDB, m module.Module) (ObjectDB, error) {
	return splayer.NewObjectDB(base, m)
}

func (p Module) NewObjectDB(m module.Module) (ObjectDB, error) {
	base, err := p.dbLayer.NewObjectDB(m)
	if err != nil {
		return nil, err
	}
	return p.spLayer.NewObjectDB(base, m)
}

func NewObjectDB(m module.Module) (ObjectDB, error) {
	return p.NewObjectDB(m)
}
