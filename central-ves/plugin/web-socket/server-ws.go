package centered_ves

import (
	"context"
	ves_websocket "github.com/Myriad-Dreamin/go-ves/lib/ves-websocket"
	"go.uber.org/atomic"
	"net/http"
)

func (c *Server) ListenAndServe(ctx context.Context, port string) error {
	go c.hub.run(ctx)
	c.Addr = port
	c.Handler.(*http.ServeMux).HandleFunc("/", c.serveWs)
	c.logger.Info("prepare to serve ws", "port", c.Addr)
	return c.Server.ListenAndServe()
}

// serveWs handles websocket requests from the peer.
func (c *Server) serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		c.logger.Error("failed to upgrade to tcp", "error", err)
		return
	}

	c.logger.Info("new ws serving\n", "remote", r.RemoteAddr)
	client := &Client{
		hub: c.hub, closed: atomic.NewBool(false), helloed: make(chan bool, 1), send: make(chan *writeMessageTask, 256)}
	client.conn, err = ves_websocket.NewVESSocket(conn, client.ProcessMessage, c.logger)
	if err != nil {
		c.logger.Error("failed to upgrade to tcp", "error", err)
		client.conn = ves_websocket.Raw(conn)
		client.closeChan()
		return
	}
	client.conn.SetCloseHandler(client.closeHandler)

	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

