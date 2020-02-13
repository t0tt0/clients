package centered_ves

import "time"

func (c *Client) pongHandler(string) error {
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		c.hub.server.logger.Error("set read ddl error", "address", c.conn.RemoteAddr(), "error", err)
	}
	return nil
}
