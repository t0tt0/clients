package model

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/database"
	dblayer "github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/db-layer"
	splayer "github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type User = database.User
type UserDB = abstraction.UserDB

func (p DBLayerModule) NewUserDB(m module.Module) (UserDB, error) {
	return dblayer.NewUserDB(p.newTraits, m)
}

func (p SPLayerModule) NewUserDB(base abstraction.UserDB, m module.Module) (UserDB, error) {
	return splayer.NewUserDB(base, m)
}

func (p Module) NewUserDB(m module.Module) (UserDB, error) {
	base, err := p.dbLayer.NewUserDB(m)
	if err != nil {
		return nil, err
	}
	return p.spLayer.NewUserDB(base, m)
}

func NewUserDB(m module.Module) (UserDB, error) {
	return p.NewUserDB(m)
}

