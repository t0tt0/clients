package hub

import (
	"encoding/hex"
	"github.com/Myriad-Dreamin/go-ves/central-ves/web-socket/client"
)

func (h *Hub) unicastMessage(client *client.Client, ok bool, message *UniMessage) {
	if ok {
		h.Server.Logger.Info("message unicasting",
			"chain id", message.Target.GetChainId(),
			"address", hex.EncodeToString(message.Target.GetAddress()))
		select {
		case client.Send <- message.Task:
		default:
			h.Server.Logger.Info("remove no response client",
				"chain id", message.Target.GetChainId(),
				"address", hex.EncodeToString(message.Target.GetAddress()))
			h.removeClient(client)
		}
	} else {
		h.Server.Logger.Info("debugging unknown aim",
			"chain id", message.Target.GetChainId(),
			"address", hex.EncodeToString(message.Target.GetAddress()))
	}
}

