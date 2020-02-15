// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package centered_ves

import (
	"context"
	"encoding/hex"
	"sync"
	"time"
	"unsafe"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Registered clients.
	reverseClients map[uint]*Client

	// Registered clients.
	reverseNameClients map[string]*Client

	// Inbound messages from the clients.
	broadcast chan *broMessage

	// messages to single clients
	unicast chan *uniMessage

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	server *Server

	mapMutex sync.Mutex
}

func newHub() *Hub {
	return &Hub{
		broadcast:          make(chan *broMessage),
		unicast:            make(chan *uniMessage),
		reverseClients:     make(map[uint]*Client),
		reverseNameClients: make(map[string]*Client),
		register:           make(chan *Client),
		unregister:         make(chan *Client),
		clients:            make(map[*Client]bool),
	}
}

func unsafeBytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func (h *Hub) run(ctx context.Context) {
	for {
		select {
		case client := <-h.register:
			h.registerClient(client)
		case client := <-h.unregister:
			h.removeClient(client)
		case message := <-h.broadcast:
			h.broadcastMessage(message)
		case message := <-h.unicast:
			if message.target.GetChainId() == placeHolderChain {
				client, ok := h.reverseNameClients[string(message.target.GetAddress())]
				h.unicastMessage(client, ok, message)
			} else {
				user, err := h.server.userDB.InvertFind(message.target)
				if err != nil {
					h.server.logger.Info("debugging unknown aim",
						"err", err,
						"chain id", message.target.GetChainId(),
						"address", hex.EncodeToString(message.target.GetAddress()))
					continue
				} else if user == nil {
					h.server.logger.Info("debugging unknown aim",
						"err", "!!not found",
						"chain id", message.target.GetChainId(),
						"address", hex.EncodeToString(message.target.GetAddress()))
					continue
				}
				client, ok := h.reverseClients[user.ID]
				h.unicastMessage(client, ok, message)
			}
		}
	}
}

func (h *Hub) removeClient(client *Client) {
	h.mapMutex.Lock()
	if _, ok := h.clients[client]; ok {
		delete(h.clients, client)
		h.mapMutex.Unlock()
		close(client.send)
		delete(h.reverseClients, client.user.GetID())
		delete(h.reverseNameClients, client.user.GetName())
	} else {
		h.mapMutex.Unlock()
	}
}

func (h *Hub) registerClient(client *Client) {
	select {
	case <-client.helloed:
		h.clients[client] = true
		h.reverseClients[client.user.GetID()] = client
		h.reverseNameClients[client.user.GetName()] = client
	case <-time.After(5 * time.Second):
		close(client.send)
		client.closeChan()
		return
	}
}

func (h *Hub) broadcastMessage(message *writeMessageTask) {
	for client := range h.clients {
		select {
		case client.send <- message:
		default:
			h.removeClient(client)
		}
	}
	//message.callback()
}

func (h *Hub) unicastMessage(client *Client, ok bool, message *uniMessage) {
	if ok {
		h.server.logger.Info("message unicasting",
			"chain id", message.target.GetChainId(),
			"address", hex.EncodeToString(message.target.GetAddress()))
		select {
		case client.send <- message.task:
		default:
			h.server.logger.Info("remove no response client",
				"chain id", message.target.GetChainId(),
				"address", hex.EncodeToString(message.target.GetAddress()))
			h.removeClient(client)
		}
	} else {
		h.server.logger.Info("debugging unknown aim",
			"chain id", message.target.GetChainId(),
			"address", hex.EncodeToString(message.target.GetAddress()))
	}
}
