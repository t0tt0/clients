package abstraction

import (
	database2 "github.com/HyperService-Consortium/go-ves/ves/model/internal/database"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Provider interface {
	module.Moduler
	Register(s string, db interface{})
	SessionAccountDB() SessionAccountDB
	SessionDB() SessionDB
	TransactionDB() TransactionDB
	ObjectDB() ObjectDB
	Enforcer() *database2.Enforcer
}
