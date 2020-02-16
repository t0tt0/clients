package chs

func (c *Client) Close() {
	if c.Closed.CAS(false, true) {
		c.close()
	}
}

func (c *Client) close() {
	err := c.Conn.Close()
	if err != nil {
		c.Hub.Server.Logger.Error("close error", "address", c.Conn.RemoteAddr(), "error", err)
	}
}
