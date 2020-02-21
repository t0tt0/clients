package provider

import (
	"github.com/HyperService-Consortium/go-ves/ves/model/internal/database"
)

func (s *Provider) Enforcer() *database.Enforcer {
	return s.enforcer
}
