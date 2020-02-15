package service

import (
	"github.com/Myriad-Dreamin/go-ves/ves/control"
	sessionservice "github.com/Myriad-Dreamin/go-ves/ves/service/session"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

// go:generate go run github.com/Myriad-Dreamin/minimum-lib/code-gen/test-gen -source ./ -destination ../../test/

type SessionService = control.SessionService

func NewSessionService(m module.Module) (SessionService, error) {
	return sessionservice.NewService(m)
}

func (s *Provider) SessionService() SessionService {
	return s.sessionService
}
