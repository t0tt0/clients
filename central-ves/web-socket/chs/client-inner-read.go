package chs

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		//c.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetPongHandler(c.pongHandler)
	if err := c.pongHandler(""); err != nil {
		return
	}
	c.Conn.ReadRoutine()
}
