package vesclient

import (
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
)

func (vc *VesClient) SendMessage(to, msg []byte) (err error) {
	shortSendMessage := vc.getShortSendMessage()
	shortSendMessage.From = vc.name
	shortSendMessage.To = to
	shortSendMessage.Contents = string(msg)

	err = vc.postMessage(wsrpc.CodeMessageRequest, shortSendMessage)
	vc.putShortSendMessage(shortSendMessage)
	return
}
