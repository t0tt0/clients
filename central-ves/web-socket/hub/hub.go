// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hub

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/web-socket/client"
	"github.com/Myriad-Dreamin/go-ves/central-ves/web-socket/server"
	"sync"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*client.Client]bool

	// Registered clients.
	reverseClients map[uint]*client.Client

	// Registered clients.
	reverseNameClients map[string]*client.Client

	// Inbound messages from the clients.
	Broadcast chan *BroMessage

	// messages to single clients
	Unicast chan *UniMessage

	// Register requests from the clients.
	Register chan *client.Client

	// Unregister requests from clients.
	Unregister chan *client.Client

	Server *server.Server

	mapMutex sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:          make(chan *BroMessage),
		Unicast:            make(chan *UniMessage),
		reverseClients:     make(map[uint]*client.Client),
		reverseNameClients: make(map[string]*client.Client),
		Register:           make(chan *client.Client),
		Unregister:         make(chan *client.Client),
		clients:            make(map[*client.Client]bool),
	}
}

