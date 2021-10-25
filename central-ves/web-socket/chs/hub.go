// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chs

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"sync"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	Server *Server
	Logger logger.Logger
	// Registered clients.
	clients map[*Client]bool

	// Registered clients.
	reverseClients map[uint]*Client

	// Registered clients.
	reverseNameClients map[string]*Client

	// Inbound messages from the clients.
	Broadcast chan *BroMessage

	// messages to single clients
	Unicast chan *UniMessage

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	mapMutex sync.Mutex
	db       *model.AccountFSet
}

func NewHub(logger logger.Logger, db *model.AccountFSet) *Hub {
	return &Hub{
		Logger:             logger,
		db:                 db,
		Broadcast:          make(chan *BroMessage),
		Unicast:            make(chan *UniMessage),
		reverseClients:     make(map[uint]*Client),
		reverseNameClients: make(map[string]*Client),
		Register:           make(chan *Client),
		Unregister:         make(chan *Client),
		clients:            make(map[*Client]bool),
	}
}
