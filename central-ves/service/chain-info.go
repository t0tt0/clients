package service

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	chainInfoservice "github.com/Myriad-Dreamin/go-ves/central-ves/service/chain-info"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

// go:generate go run github.com/Myriad-Dreamin/minimum-lib/code-gen/test-gen -source ./ -destination ../../test/

type ChainInfoService = control.ChainInfoService

func NewChainInfoService(m module.Module) (ChainInfoService, error) {
	return chainInfoservice.NewService(m)
}

func (s *Provider) ChainInfoService() ChainInfoService {
	return s.chainInfoService
}
