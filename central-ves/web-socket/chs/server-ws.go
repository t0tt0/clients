package chs

import (
	"context"
	ves_websocket "github.com/Myriad-Dreamin/go-ves/lib/ves-websocket"
	"go.uber.org/atomic"
	"net/http"
)

func (srv *Server) ListenAndServe(ctx context.Context, port string) error {
	go srv.hub.Run(ctx)
	srv.Addr = port
	srv.Handler.(*http.ServeMux).HandleFunc("/", srv.serveWs)
	srv.Logger.Info("prepare to serve ws", "port", srv.Addr)
	return srv.Server.ListenAndServe()
}

// serveWs handles websocket requests from the peer.
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
