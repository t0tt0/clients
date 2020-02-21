package splayer

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type UserDB struct {
	abstraction.UserDB
}

func NewUserDB(base abstraction.UserDB, m module.Module) (*UserDB, error) {
	db := new(UserDB)
	db.UserDB = base
	return db, nil
}
