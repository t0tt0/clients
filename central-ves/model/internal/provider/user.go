package provider

import "github.com/HyperService-Consortium/go-ves/central-ves/model/internal/abstraction"

func (s *Provider) UserDB() abstraction.UserDB {
	return s.userDB
}
