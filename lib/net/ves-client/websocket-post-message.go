package vesclient

import (
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
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

func (vc *VesClient) postRawMessage(code wsrpc.MessageType, dst *uiprpc_base.Account, msg proto.Message) error {

	buf, err := wsrpc.GetDefaultSerializer().Serial(code, msg)
	/// fmt.Println(buf.Bytes())
	if err != nil {
		return err
	}
	var s = vc.getRawMessage()
	s.To, err = proto.Marshal(dst)
	if err != nil {
		return err
	}
	s.From = vc.name
	s.Contents = make([]byte, buf.Len())
	copy(s.Contents, buf.Bytes())
	// fmt.Println(s.Contents)
	wsrpc.GetDefaultSerializer().Put(buf)
	return vc.postMessage(wsrpc.CodeRawProto, s)
}

