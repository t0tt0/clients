package vesclient

import (
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
)

func (vc *VesClient) RetransmitAttestationReceiveRequest(
	target *uiprpc_base.Account, req *wsrpc.AttestationReceiveRequest) error {
	return vc.PostRawMessage(
		wsrpc.CodeAttestationReceiveRequest, target, req)
}
