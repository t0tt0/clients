package chs

import "time"

func (c *Client) pongHandler(string) error {
	err := c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		c.Hub.Server.Logger.Error("set read ddl error", "address", c.Conn.RemoteAddr(), "error", err)
	}
	return nil
}
