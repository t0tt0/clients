package centered_ves

import (
	"github.com/gorilla/websocket"
	"time"
)

func (c *Client) closeChan() {
	if c.closed.CAS(false, true) {
		message := websocket.FormatCloseMessage(
			websocket.ClosePolicyViolation,
			"client hello please",
		)
		err := c.conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(writeWait))
		if err != nil && err != websocket.ErrCloseSent {
			c.hub.server.logger.Error(
				"write close message error",
				"address", c.conn.RemoteAddr(), "error", err)
			c.close()
		}
	}
}
