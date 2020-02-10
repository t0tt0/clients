package model

import (
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	dblayer "github.com/Myriad-Dreamin/go-ves/ves/model/db-layer"
	splayer "github.com/Myriad-Dreamin/go-ves/ves/model/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type SessionAccount = splayer.SessionAccount
type SessionAccountDB = splayer.SessionAccountDB

func NewSessionAccountDB(m module.Module) (*SessionAccountDB, error) {
	return splayer.NewSessionAccountDB(m)
}

func GetSessionAccountDB(m module.Module) (*SessionAccountDB, error) {
	return splayer.GetSessionAccountDB(m)
}

func SessionAccountsToUIPAccounts(accounts []SessionAccount) (uipAccounts []uiptypes.Account) {
	uipAccounts = make([]uiptypes.Account, len(accounts))
	for i := range accounts {
		uipAccounts[i] = accounts[i]
	}
	return uipAccounts
}

func NewSessionAccount(chainID uiptypes.ChainIDUnderlyingType, address []byte) *SessionAccount {
	return dblayer.NewSessionAccount(chainID, address)
}
