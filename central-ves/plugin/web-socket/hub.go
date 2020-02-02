// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package centered_ves

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"sync"
	"time"
	"unsafe"
)

const (
	localChain       = uint64((127 << 24) + 1)
	placeHolderChain = uint64((127 << 24) + 2)
)

type broMessage = writeMessageTask
type uniMessage struct {
	target uiptypes.Account
	task   *writeMessageTask
}

type clientKey struct {
	chainID uint64
	address string
}

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

	server *CVESWebSocketPlugin

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
				client, ok := h.reverseNameClients[
					string(message.target.GetAddress())]
				h.unicastMessage(client, ok, message)
			} else {
				tag := md5.Sum(message.task.b)
				user, err := h.server.userDB.InvertFind(message.target)
				if err != nil {
					h.server.logger.Info("debugging unknown aim",
						"err", err,
						"tag", hex.EncodeToString(tag[:]),
						"chain id", message.target.GetChainId(),
						"address", hex.EncodeToString(message.target.GetAddress()))
				} else if user == nil {
					h.server.logger.Info("debugging unknown aim",
						"err", "!!not found",
						"tag", hex.EncodeToString(tag[:]),
						"chain id", message.target.GetChainId(),
						"address", hex.EncodeToString(message.target.GetAddress()))
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
		// do nothing
	case <-time.After(5 * time.Second):
		close(client.send)
		client.closeChan()
		return
	}

	h.clients[client] = true
	h.reverseClients[client.user.GetID()] = client
	h.reverseNameClients[client.user.GetName()] = client
}

func (h *Hub) broadcastMessage(message *writeMessageTask) {
	tag := md5.Sum(message.b)
	h.server.logger.Info("message broadcasting", "tag", hex.EncodeToString(tag[:]))
	for client := range h.clients {
		select {
		case client.send <- message:
		default:
			h.removeClient(client)
		}
	}
	//message.cb()
}

func (h *Hub) unicastMessage(client *Client, ok bool, message *uniMessage) {
	tag := md5.Sum(message.task.b)
	if !ok {
		h.server.logger.Info("message unicasting",
			"tag", hex.EncodeToString(tag[:]),
			"chain id", message.target.GetChainId(),
			"address", hex.EncodeToString(message.target.GetAddress()))
		select {
		case client.send <- message.task:
		default:
			h.server.logger.Info("remove no response client",
				"tag", hex.EncodeToString(tag[:]),
				"chain id", message.target.GetChainId(),
				"address", hex.EncodeToString(message.target.GetAddress()))
			h.removeClient(client)
		}
	} else {
		h.server.logger.Info("debugging unknown aim",
			"tag", hex.EncodeToString(tag[:]),
			"chain id", message.target.GetChainId(),
			"address", hex.EncodeToString(message.target.GetAddress()))
	}
}
