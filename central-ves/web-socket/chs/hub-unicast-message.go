package chs

import (
	"encoding/hex"
)

func (h *Hub) unicastMessage(client *Client, ok bool, message *UniMessage) {
	if ok {
		h.Logger.Info("message unicasting",
			"chain id", message.Target.GetChainId(),
			"address", hex.EncodeToString(message.Target.GetAddress()))
		select {
		case client.Send <- message.Task:
		default:
			h.Logger.Error("remove no response client",
				"chain id", message.Target.GetChainId(),
				"address", hex.EncodeToString(message.Target.GetAddress()))
			h.removeClient(client)
		}
	} else {
		h.Logger.Info("debugging unknown aim",
			"chain id", message.Target.GetChainId(),
			"address", hex.EncodeToString(message.Target.GetAddress()))
	}
}
