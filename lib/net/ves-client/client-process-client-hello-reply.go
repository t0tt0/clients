package vesclient

import (
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
	nsbcli "github.com/Myriad-Dreamin/go-ves/lib/net/nsb-client"
)

func (vc *VesClient) ProcessClientHelloReply(req *wsrpc.ClientHelloReply) {
	vc.grpcip = req.GetGrpcHost()
	vc.logger.Info("adding default grpc ip ", "ip", vc.grpcip)

	vc.nsbip = req.GetNsbHost()
	vc.logger.Info("adding default nsb ip ", "ip", vc.nsbip)

	// todo: restrict scope
	vc.nsbClient = nsbcli.NewNSBClient(vc.nsbip)
}
