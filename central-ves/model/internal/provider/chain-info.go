package provider

import "github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/abstraction"

func (s *Provider) ChainInfoDB() abstraction.ChainInfoDB {
	return s.chainInfoDB
}
