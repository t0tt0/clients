package client

import "github.com/gorilla/websocket"

func (c *Client) handleReadMessageError(err error) {
	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
		c.Hub.Server.Logger.Info("close error", "error", err)
	}
}
