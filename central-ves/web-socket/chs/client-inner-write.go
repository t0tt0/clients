package chs

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"reflect"
	"time"
)

// writePump pumps messages from the hub to the websocket connection.
////
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Close()
	}()
	var err error
	for {
		c.setWriteDeadLine()
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.CloseChan()
				return
			}

			switch msg := message.Message.(type) {
			case proto.Message:
				err = c.Conn.PostMessage(message.MessageType, msg)
			case []byte:
				err = c.Conn.PostRawPacket(message.MessageType, msg)
			default:
				c.Hub.Server.Logger.Error(
					"bad message type", "msgT", reflect.TypeOf(msg))
			}
			if err != nil {
				c.Hub.Server.Logger.Error(
					"post error", "error", err)

			}
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
