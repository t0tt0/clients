package provider

import "github.com/Myriad-Dreamin/go-ves/ves/model/internal/abstraction"

func (s *Provider) SessionDB() abstraction.SessionDB {
	return s.sessionDB
}
