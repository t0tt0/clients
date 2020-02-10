package service

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"golang.org/x/net/context"

	uiprpc "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
)

type SessionRequireTransactService struct {
	model.VESDB
	context.Context
	*uiprpc.SessionRequireTransactRequest
}

func (s SessionRequireTransactService) Serve() (*uiprpc.SessionRequireTransactReply, error) {
	// todo errors.New("TODO")
	s.ActivateSession(s.GetSessionId())
	defer s.InactivateSession(s.GetSessionId())
	return &uiprpc.SessionRequireTransactReply{
		// Tx: true,
	}, nil
}
