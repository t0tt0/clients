package splayer

import (
	dblayer "github.com/Myriad-Dreamin/go-ves/ves/model/db-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type SessionAccount = dblayer.SessionAccount

type SessionAccountDB struct {
	dblayer.SessionAccountDB
}

func NewSessionAccountDB(m module.Module) (*SessionAccountDB, error) {
	cdb, err := dblayer.NewSessionAccountDB(m)
	if err != nil {
		return nil, err
	}
	db := new(SessionAccountDB)
	db.SessionAccountDB = *cdb
	return db, nil
}

func GetSessionAccountDB(m module.Module) (*SessionAccountDB, error) {
	cdb, err := dblayer.GetSessionAccountDB(m)
	if err != nil {
		return nil, err
	}
	db := new(SessionAccountDB)
	db.SessionAccountDB = *cdb
	return db, nil
}

func (s *Provider) SessionAccountDB() *SessionAccountDB {
	return s.sessionAccountDB
}
