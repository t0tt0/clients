package model

import (
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/abstraction"
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/database"
	dblayer "github.com/HyperService-Consortium/go-ves/ves/model/internal/db-layer"
	splayer "github.com/HyperService-Consortium/go-ves/ves/model/internal/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Session = database.Session
type SessionDB = abstraction.SessionDB

func (p DBLayerModule) NewSessionDB(m module.Module) (SessionDB, error) {
	return dblayer.NewSessionDB(p.newTraits, m)
}

func (p SPLayerModule) NewSessionDB(base abstraction.SessionDB, m module.Module) (SessionDB, error) {
	return splayer.NewSessionDB(base, m)
}

func (p Module) NewSessionDB(m module.Module) (SessionDB, error) {
	base, err := p.dbLayer.NewSessionDB(m)
	if err != nil {
		return nil, err
	}
	return p.spLayer.NewSessionDB(base, m)
}

func NewSessionDB(m module.Module) (SessionDB, error) {
	return p.NewSessionDB(m)
}
