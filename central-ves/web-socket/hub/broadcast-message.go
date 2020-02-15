package hub

func (h *Hub) broadcastMessage(message *WriteMessageTask) {
	for client := range h.clients {
		select {
		case client.Send <- message:
		default:
			h.removeClient(client)
		}
	}
	//message.callback()
}
