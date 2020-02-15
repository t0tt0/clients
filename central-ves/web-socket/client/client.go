// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/go-ves/central-ves/web-socket/hub"
	"github.com/Myriad-Dreamin/go-ves/lib/ves-websocket"
	"go.uber.org/atomic"
)


// Client is a middleman between the websocket connection and the hub.
type Client struct {
	Hub *hub.Hub

	// The websocket connection.
	Conn ves_websocket.VESWSSocket

	// owned user
	User *model.User

	// Buffered channel of outbound messages.
	Send chan *hub.WriteMessageTask

	// client hello sended
	Helloed chan bool
	Closed  *atomic.Bool
}


