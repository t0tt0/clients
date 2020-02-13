package vesclient

import (
	"crypto/md5"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
	"github.com/Myriad-Dreamin/go-ves/lib/encoding"
)

func (vc *VesClient) ReadMessage() (message []byte, messageID wsrpc.MessageType, err error) {
	_, message, err = vc.conn.ReadMessage()
	if err != nil {
		err = wrap(CodeReadMessageError, err)
		return
	}
	tag := md5.Sum(message)
	vc.logger.Info("message from server", "tag", encoding.EncodeHex(tag[:]), "type", messageID)

	message, messageID, err = wsrpc.Deserialize(message)
	if err != nil {
		err = wrap(CodeReadMessageIDError, err)
		return
	}

	return
}
