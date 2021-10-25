package config

import (
	"github.com/HyperService-Consortium/go-uip/uip"
)

type accountProvider struct {
}

func (a accountProvider) AccountBase() uip.AccountBase {
	return a
}

func (accountProvider) Get(name string, chainId uint64) (uip.Account, error) {
	return searchAccount(name, chainId)
}

func (accountProvider) GetRelay(domain uint64) (uip.Account, error) {
	return getRelay(domain)
}

func (accountProvider) GetTransactionProofType(chainId uint64) (uip.MerkleProofType, error) {
	return getTransactionProofType(chainId)
}

var UserMap = accountProvider{}
