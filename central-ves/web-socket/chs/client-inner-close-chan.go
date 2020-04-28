package chs

import (
	"github.com/gorilla/websocket"
	"time"
)

func (c *Client) CloseChan() {
	if c.Closed.CAS(false, true) {
		message := websocket.FormatCloseMessage(
			websocket.ClosePolicyViolation,
			"client hello please, or io timeout",
		)
		err := c.Conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(writeWait))
		if err != nil && err != websocket.ErrCloseSent {
			c.Hub.Server.Logger.Error(
				"write close message error",
				"address", c.Conn.RemoteAddr(), "error", err)
			c.close()
		}
	}
}
