package vesclient

import (
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
)

func (vc *VesClient) postMessage(code wsrpc.MessageType, msg proto.Message) error {
	buf, err := wsrpc.GetDefaultSerializer().Serial(code, msg)
	if err != nil {
		return err
	}
	err = vc.conn.WriteMessage(websocket.BinaryMessage, buf.Bytes())
	if err != nil {
		return err
	}
	wsrpc.GetDefaultSerializer().Put(buf)
	return nil
}
