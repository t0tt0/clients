package provider

import "github.com/HyperService-Consortium/go-ves/ves/model/internal/abstraction"

func (s *Provider) SessionDB() abstraction.SessionDB {
	return s.sessionDB
}
