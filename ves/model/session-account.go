package model

import (
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/go-ves/ves/model/internal/database"
	dblayer "github.com/Myriad-Dreamin/go-ves/ves/model/internal/db-layer"
	splayer "github.com/Myriad-Dreamin/go-ves/ves/model/internal/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type SessionAccount = database.SessionAccount
type SessionAccountDB = abstraction.SessionAccountDB

func (p DBLayerModule) NewSessionAccountDB(m module.Module) (SessionAccountDB, error) {
	return dblayer.NewSessionAccountDB(p.newTraits, m)
}

func (p SPLayerModule) NewSessionAccountDB(base abstraction.SessionAccountDB, m module.Module) (SessionAccountDB, error) {
	return splayer.NewSessionAccountDB(base, m)
}

func (p Module) NewSessionAccountDB(m module.Module) (SessionAccountDB, error) {
	base, err := p.dbLayer.NewSessionAccountDB(m)
	if err != nil {
		return nil, err
	}
	return p.spLayer.NewSessionAccountDB(base, m)
}

func NewSessionAccountDB(m module.Module) (SessionAccountDB, error) {
	return p.NewSessionAccountDB(m)
}
func SessionAccountsToUIPAccounts(accounts []SessionAccount) (uipAccounts []uip.Account) {
	return database.SessionAccountsToUIPAccounts(accounts)
}

func NewSessionAccount(chainID uip.ChainIDUnderlyingType, address []byte) *SessionAccount {
	return database.NewSessionAccount(chainID, address)
}