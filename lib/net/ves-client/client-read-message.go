package vesclient

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
)

func (vc *VesClient) ReadMessage() (message []byte, messageID uint16, err error) {
	_, message, err = vc.conn.ReadMessage()
	if err != nil {
		err = wrap(CodeReadMessageError, err)
		return
	}

	var buf = bytes.NewBuffer(message)
	// todo: BigEndian
	err = binary.Read(buf, binary.BigEndian, &messageID)
	if err != nil {
		err = wrap(CodeReadMessageIDError, err)
		return
	}

	tag := md5.Sum(message)
	vc.logger.Info("message from server", "tag", hex.EncodeToString(tag[:]), "type", wsrpc.MessageType(messageID))

	message = buf.Bytes()
	return
}
