package model

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/model/internal/abstraction"
	"github.com/HyperService-Consortium/go-ves/central-ves/model/internal/fset"
)

type AccountFSet = fset.AccountFSet

func NewAccountFSet(p abstraction.Provider) *AccountFSet {
	return fset.NewAccountFSet(p)
}
