package provider

import (
	"fmt"
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/abstraction"
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/database"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"path"
)

var provider *Provider

type Provider struct {
	module.BaseModuler
	objectDB abstraction.ObjectDB

	sessionDB        abstraction.SessionDB
	sessionAccountDB abstraction.SessionAccountDB
	transactionDB    abstraction.TransactionDB
	enforcer         *database.Enforcer
}

func NewProvider(namespace string) *Provider {
	return &Provider{
		BaseModuler: module.BaseModuler{
			Namespace: namespace,
		},
	}
}

func (s *Provider) Register(name string, db interface{}) {
	if err := s.Provide(path.Join(s.Namespace, name), db); err != nil {
		panic(fmt.Errorf("unknown db %T", db))
	}

	switch ss := db.(type) {
	case *database.Enforcer:
		s.enforcer = ss
	case abstraction.SessionDB:
		s.sessionDB = ss
	case abstraction.SessionAccountDB:
		s.sessionAccountDB = ss
	case abstraction.TransactionDB:
		s.transactionDB = ss
	case abstraction.ObjectDB:
		s.objectDB = ss
	default:
		if mm, ok := ss.(module.Moduler); ok {
			// todo:
			_ = mm
		}
	}
}
