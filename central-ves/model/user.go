package model

import (
	splayer "github.com/Myriad-Dreamin/go-ves/central-ves/model/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type User = splayer.User

type UserDB = splayer.UserDB

func NewUserDB(m module.Module) (*UserDB, error) {
	return splayer.NewUserDB(m)
}

func GetUserDB(m module.Module) (*UserDB, error) {
	return splayer.GetUserDB(m)
}
