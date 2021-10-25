package service

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/control"
	authservice "github.com/HyperService-Consortium/go-ves/central-ves/service/auth"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

// go:generate go run github.com/Myriad-Dreamin/minimum-lib/code-gen/test-gen -source ./ -destination ../../test/

type AuthService = control.AuthService

func NewAuthService(m module.Module) (AuthService, error) {
	return authservice.NewService(m)
}

func (s *Provider) AuthService() AuthService {
	return s.authService
}
