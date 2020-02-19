package provider

import (
	"fmt"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/abstraction"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/database"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"path"
)

var provider *Provider

type Provider struct {
	module.BaseModuler
	objectDB abstraction.ObjectDB

	chainInfoDB abstraction.ChainInfoDB
	userDB      abstraction.UserDB
	enforcer    *database.Enforcer
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
	case abstraction.ChainInfoDB:
		s.chainInfoDB = ss
	case *database.Enforcer:
		s.enforcer = ss
	case abstraction.UserDB:
		s.userDB = ss
	case abstraction.ObjectDB:
		s.objectDB = ss
	default:
		if mm, ok := ss.(module.Moduler); ok {
			// todo:
			_ = mm
		}
	}
}

