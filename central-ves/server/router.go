package server

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/control/router"
)

func (srv *Server) BuildRouter() bool {
	srv.Router = router.NewRootRouter(srv.Module)
	return true
}
