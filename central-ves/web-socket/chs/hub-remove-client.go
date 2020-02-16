package chs

func (h *Hub) removeClient(client *Client) {
	h.mapMutex.Lock()
	if _, ok := h.clients[client]; ok {
		delete(h.clients, client)
		h.mapMutex.Unlock()
		close(client.Send)
		delete(h.reverseClients, client.User.GetID())
		delete(h.reverseNameClients, client.User.GetName())
	} else {
		h.mapMutex.Unlock()
	}
}
