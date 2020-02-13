package vesclient

import (
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
)

func (vc *VesClient) SendMessage(to *uiprpc_base.Account, msg []byte) (err error) {
	shortSendMessage := vc.getShortSendMessage()
	//shortSendMessage.From = vc.name
	shortSendMessage.To = to
	shortSendMessage.Contents = string(msg)

	err = vc.conn.PostMessage(wsrpc.CodeMessageRequest, shortSendMessage)
	vc.putShortSendMessage(shortSendMessage)
	return
}
