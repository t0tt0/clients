package centered_ves

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"reflect"
	"time"
)

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Close()
	}()
	var err error
	for {
		c.setWriteDeadLine()
		select {
		case message, ok := <-c.send:
			if !ok {
				c.closeChan()
				return
			}

			switch msg := message.message.(type) {
			case proto.Message:
				err = c.conn.PostMessage(message.messageType, msg)
			case []byte:
				err = c.conn.PostRawPacket(message.messageType, msg)
			default:
				c.hub.server.logger.Error(
					"bad message type", "msgT", reflect.TypeOf(msg))
			}
			if err != nil {
				c.hub.server.logger.Error(
					"post error", "error", err)

			}
		case <-ticker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
