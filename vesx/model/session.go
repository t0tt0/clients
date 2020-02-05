package model

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	splayer "github.com/Myriad-Dreamin/go-ves/vesx/model/sp-layer"
)

type Session = splayer.Session
type SessionDB = splayer.SessionDB

func NewSessionDB(m module.Module) (*SessionDB, error) {
	return splayer.NewSessionDB(m)
}

func GetSessionDB(m module.Module) (*SessionDB, error) {
	return splayer.GetSessionDB(m)
}
