package provider

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/database"
)

func (s *Provider) Enforcer() *database.Enforcer {
	return s.enforcer
}
