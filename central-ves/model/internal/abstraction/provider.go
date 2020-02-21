package abstraction

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/database"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Provider interface {
	module.Moduler
	Register(s string, db interface{})
	ChainInfoDB() ChainInfoDB
	UserDB() UserDB
	ObjectDB() ObjectDB
	Enforcer() *database.Enforcer
}
