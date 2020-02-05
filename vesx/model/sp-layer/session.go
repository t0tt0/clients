package splayer

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	dblayer "github.com/Myriad-Dreamin/go-ves/vesx/model/db-layer"
)

type Session = dblayer.Session

type SessionDB struct {
	dblayer.SessionDB
}

func NewSessionDB(m module.Module) (*SessionDB, error) {
	cdb, err := dblayer.NewSessionDB(m)
	if err != nil {
		return nil, err
	}
	db := new(SessionDB)
	db.SessionDB = *cdb
	return db, nil
}

func GetSessionDB(m module.Module) (*SessionDB, error) {
	cdb, err := dblayer.GetSessionDB(m)
	if err != nil {
		return nil, err
	}
	db := new(SessionDB)
	db.SessionDB = *cdb
	return db, nil
}

func (s *Provider) SessionDB() *SessionDB {
	return s.sessionDB
}
