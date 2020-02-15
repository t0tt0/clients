package hub

import (
	"context"
	"encoding/hex"
)

func (h *Hub) Run(ctx context.Context) {
	for {
		select {
		case client := <-h.Register:
			h.registerClient(client)
		case client := <-h.Unregister:
			h.removeClient(client)
		case message := <-h.Broadcast:
			h.broadcastMessage(message)
		case message := <-h.Unicast:
			if message.Target.GetChainId() == PlaceHolderChain {
				client, ok := h.reverseNameClients[string(message.Target.GetAddress())]
				h.unicastMessage(client, ok, message)
			} else {
				user, err := h.Server.UserDB.InvertFind(message.Target)
				if err != nil {
					h.Server.Logger.Info("debugging unknown aim",
						"err", err,
						"chain id", message.Target.GetChainId(),
						"address", hex.EncodeToString(message.Target.GetAddress()))
					continue
				} else if user == nil {
					h.Server.Logger.Info("debugging unknown aim",
						"err", "!!not found",
						"chain id", message.Target.GetChainId(),
						"address", hex.EncodeToString(message.Target.GetAddress()))
					continue
				}
				client, ok := h.reverseClients[user.ID]
				h.unicastMessage(client, ok, message)
			}
		}
	}
}
