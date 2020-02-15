package client

import "time"

func (c *Client) setWriteDeadLine() {
	if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
		c.Hub.Server.Logger.Error("set write ddl error", "error", err)
	}
}


