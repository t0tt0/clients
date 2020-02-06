package model

import (
	dblayer "github.com/Myriad-Dreamin/go-ves/vesx/model/db-layer"
	splayer "github.com/Myriad-Dreamin/go-ves/vesx/model/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Session = splayer.Session
type SessionDB = splayer.SessionDB

func NewSessionDB(m module.Module) (*SessionDB, error) {
	return splayer.NewSessionDB(m)
}

func GetSessionDB(m module.Module) (*SessionDB, error) {
	return splayer.GetSessionDB(m)
}

func NewSession(iscAddress []byte) *Session {
	return dblayer.NewSession(iscAddress)
}
