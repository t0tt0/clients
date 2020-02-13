package centered_ves

func (c *Client) Close() {
	if c.closed.CAS(false, true) {
		c.close()
	}
}

func (c *Client) close() {
	err := c.conn.Close()
	if err != nil {
		c.hub.server.logger.Error("close error", "address", c.conn.RemoteAddr(), "error", err)
	}
}
