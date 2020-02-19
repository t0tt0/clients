package chs

import (
	"context"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/fset"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"net/http"
)

// Server is a client manager, named centered ves
// it is not in the standard of uip
type Server struct {
	Logger logger.Logger
	*http.Server
	hub     *Hub
	UserDB  *fset.AccountFSet
	rpcPort string
	Nsbip   string
}

// NewServer return a pointer of Server
func NewServer(rpcPort, addr string, db *fset.AccountFSet, rOptions ...interface{}) (srv *Server, err error) {
	options := parseOptions(rOptions)
	srv = &Server{
		Server: &http.Server{
			Handler: http.NewServeMux(),
			Addr:    addr,
		},
		hub:     NewHub(options.logger, db),
		UserDB:  db,
		rpcPort: rpcPort,
		Logger:  options.logger,
		Nsbip:   options.nsbHost,
	}
	srv.hub.Server = srv
	return
}

// Start the service of centered ves
func (srv *Server) Start(ctx context.Context) error {
	go srv.ListenAndServeRpc(ctx, srv.rpcPort)
	return srv.ListenAndServe(ctx, srv.Addr)
}

func (srv *Server) ProvideUserDB(db *fset.AccountFSet) {
	srv.UserDB = db
	srv.hub.db = db
}
