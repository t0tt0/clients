// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package centered_ves

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/go-ves/lib/ves-websocket"
	"github.com/gorilla/websocket"
	"go.uber.org/atomic"
	"time"
)


// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn ves_websocket.VESWSSocket

	// owned user
	user *model.User

	// Buffered channel of outbound messages.
	send chan *writeMessageTask

	// client hello sended
	helloed chan bool
	closed  *atomic.Bool
}


func (c *Client) closeHandler(code int, text string) error {
	c.closed.CAS(false, true)
	if code != websocket.CloseNoStatusReceived {
		c.hub.server.logger.Info("closed", "code", code, "text", text)
	}
	return nil
}


func (c *Client) handleReadMessageError(err error) {
	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
		c.hub.server.logger.Info("close error", "error", err)
	}
}

func (c *Client) setWriteDeadLine() {
	if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
		c.hub.server.logger.Error("set write ddl error", "error", err)
	}
}
