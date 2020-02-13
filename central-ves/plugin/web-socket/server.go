package centered_ves

import (
	"context"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/fset"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"net/http"
)

// Server is a client manager, named centered ves
// it is not in the standard of uip
type Server struct {
	logger logger.Logger
	*http.Server
	hub     *Hub
	userDB  *fset.AccountFSet
	rpcPort string
	nsbip   string
}

// NewServer return a pointer of Server
func NewServer(rpcPort, addr string, db *fset.AccountFSet, rOptions ...interface{}) (srv *Server, err error) {
	options := parseOptions(rOptions)
	srv = &Server{
		Server: &http.Server{
			Handler: http.NewServeMux(),
			Addr:    addr,
		},
		hub:     newHub(),
		userDB:  db,
		rpcPort: rpcPort,
		logger:  options.logger,
		nsbip:   options.nsbHost,
	}

	srv.hub.server = srv
	return
}

// Start the service of centered ves
func (c *Server) Start(ctx context.Context) error {
	go c.ListenAndServeRpc(ctx, c.rpcPort)
	return c.ListenAndServe(ctx, c.Addr)
}

