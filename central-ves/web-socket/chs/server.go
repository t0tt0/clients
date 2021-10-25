package chs

import (
	"context"
	"github.com/HyperService-Consortium/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"net/http"
)


//websocket server as a plugin in a central-ves server
// Server is a client manager, named centered ves
// it is not in the standard of uip
type Server struct {
	Logger logger.Logger
	*http.Server
	hub     *Hub
	UserDB  *model.AccountFSet
	rpcPort string
	Nsbip   string
}

// NewServer return a pointer of c-ves websocket Server, as the configuration injected in central-ves server
//addr is the websocket port defined in the config.go file
func NewServer(rpcPort, addr string, db *model.AccountFSet, rOptions ...interface{}) (srv *Server, err error) {
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

// Start the service of web socket ves in center-ves
func (srv *Server) Start(ctx context.Context) error {
	go srv.ListenAndServeRpc(ctx, srv.rpcPort)
	return srv.ListenAndServe(ctx, srv.Addr)
}

func (srv *Server) ProvideUserDB(db *model.AccountFSet) {
	srv.UserDB = db
	srv.hub.db = db
}
