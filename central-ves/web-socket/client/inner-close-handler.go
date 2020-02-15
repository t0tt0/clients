package client

import "github.com/gorilla/websocket"

func (c *Client) CloseHandler(code int, text string) error {
	c.Closed.CAS(false, true)
	if code != websocket.CloseNoStatusReceived {
		c.Hub.Server.Logger.Info("closed", "code", code, "text", text)
	}
	return nil
}
