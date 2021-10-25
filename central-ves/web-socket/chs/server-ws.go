package chs

import (
	"context"
	ves_websocket "github.com/HyperService-Consortium/go-ves/lib/net/ves-websocket"
	"go.uber.org/atomic"
	"net/http"
)


//handle messages in servews, clients are objectes, receiving messages from different peers
func (srv *Server) ListenAndServe(ctx context.Context, port string) error {
	go srv.hub.Run(ctx)
	srv.Addr = port
	srv.Handler.(*http.ServeMux).HandleFunc("/", srv.serveWs)
	srv.Logger.Info("prepare to serve ws", "port", srv.Addr)

	//http requests, the handshake process in the websocket fuctions is conducted through http request
	return srv.Server.ListenAndServe()
}

// serveWs handles websocket requests from the peer. be used in listenandserve /web-socket/chs/client-process.go
func (srv *Server) serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		srv.Logger.Error("failed to upgrade to tcp", "error", err)
		return
	}

	srv.Logger.Info("new ws serving\n", "remote", r.RemoteAddr)
	c := &Client{
		Hub: srv.hub, Closed: atomic.NewBool(false), Helloed: make(chan bool, 1), Send: make(chan *WriteMessageTask, 256)}
	c.Conn, err = ves_websocket.NewVESSocket(conn, c.ProcessMessage, srv.Logger)
	if err != nil {
		srv.Logger.Error("failed to upgrade to tcp", "error", err)
		c.Conn = ves_websocket.Raw(conn)
		c.CloseChan()
		return
	}
	c.Conn.SetCloseHandler(c.CloseHandler)

	c.Hub.Register <- c

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go c.WritePump()
	go c.ReadPump()
}
