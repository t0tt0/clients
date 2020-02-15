package server

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/control/router"
)

func (srv *Server) BuildRouter() bool {
	srv.Router = router.NewRootRouter(srv.Module)
	return true
}
