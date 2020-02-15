package hub

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/web-socket/client"
	"time"
)

func (h *Hub) registerClient(client *client.Client) {
	select {
	case <-client.Helloed:
		h.clients[client] = true
		h.reverseClients[client.User.GetID()] = client
		h.reverseNameClients[client.User.GetName()] = client
	case <-time.After(5 * time.Second):
		close(client.Send)
		client.CloseChan()
		return
	}
}

