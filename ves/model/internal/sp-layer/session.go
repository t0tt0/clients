package splayer

import (
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type SessionDB struct {
	abstraction.SessionDB
}

func NewSessionDB(base abstraction.SessionDB, m module.Module) (*SessionDB, error) {
	db := new(SessionDB)
	db.SessionDB = base
	return db, nil
}
