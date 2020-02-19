package provider

import "github.com/Myriad-Dreamin/go-ves/ves/model/internal/abstraction"

func (s *Provider) SessionAccountDB() abstraction.SessionAccountDB {
	return s.sessionAccountDB
}
