package vesclient

import (
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
	"github.com/gogo/protobuf/proto"
)

// for retransmitting to other client
func (vc *VesClient) PostRawMessage(code wsrpc.MessageType, dst *uiprpc_base.Account, msg proto.Message) error {

	b, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	return vc.conn.PostMessage(wsrpc.CodeRawProto, vc.combineRawMessage(dst, code, b))
}
