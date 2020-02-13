package vesclient

import (
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
	"github.com/gogo/protobuf/proto"
)

//err = svc.postRawMessage(wsrpc.CodeAttestationReceiveRequest, s.GetSrc(), sendingAtte)
//	if err != nil {
//		svc.client.logger.Error("postRawMessage", "error", err)
//		return nil
//	}

// for retransmitting to other client
func (vc *VesClient) PostRawMessage(code wsrpc.MessageType, dst *uiprpc_base.Account, msg proto.Message) error {

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
