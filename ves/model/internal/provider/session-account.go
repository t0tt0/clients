package provider

import "github.com/HyperService-Consortium/go-ves/ves/model/internal/abstraction"

func (s *Provider) SessionAccountDB() abstraction.SessionAccountDB {
	return s.sessionAccountDB
}
