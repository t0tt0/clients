package splayer

import (
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"path"
)

var provider *Provider

type Provider struct {
	module.BaseModuler
	objectDB *ObjectDB

    transactionDB *TransactionDB
	sessionAccountDB *SessionAccountDB
	sessionDB        *SessionDB
	enforcer         *Enforcer
}

func NewProvider(namespace string) *Provider {
	return &Provider{
		BaseModuler: module.BaseModuler{
			Namespace: namespace,
		},
	}
}

func (s *Provider) Register(name string, database interface{}) {
	if err := s.Provide(path.Join(s.Namespace, name), database); err != nil {
		panic(fmt.Errorf("unknown database %T", database))
	}

	switch ss := database.(type) {
    case *TransactionDB:
        s.transactionDB = ss
	case *SessionAccountDB:
		s.sessionAccountDB = ss
	case *SessionDB:
		s.sessionDB = ss
	case *Enforcer:
		s.enforcer = ss
	case *ObjectDB:
		s.objectDB = ss
	default:
		if mm, ok := ss.(module.Moduler); ok {
			// todo:
			_ = mm
		}
	}
}

func SetProvider(p *Provider) (op *Provider) {
	op = provider
	provider = p
	return
}
