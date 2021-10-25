package provider

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/model/internal/database"
)

func (s *Provider) Enforcer() *database.Enforcer {
	return s.enforcer
}
