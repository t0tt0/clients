package splayer

import (
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type SessionAccountDB struct {
	abstraction.SessionAccountDB
}

func NewSessionAccountDB(base abstraction.SessionAccountDB, m module.Module) (*SessionAccountDB, error) {
	db := new(SessionAccountDB)
	db.SessionAccountDB = base
	return db, nil
}
