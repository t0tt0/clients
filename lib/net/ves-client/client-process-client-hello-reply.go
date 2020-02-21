package vesclient

import (
	"github.com/HyperService-Consortium/go-ves/grpc/wsrpc"
	nsbcli "github.com/HyperService-Consortium/go-ves/lib/net/nsb-client"
)

func (vc *VesClient) ProcessClientHelloReply(req *wsrpc.ClientHelloReply) {
	vc.vesHost = req.GetGrpcHost()
	vc.logger.Info("adding default grpc ip ", "ip", vc.vesHost)

	vc.nsbHost = req.GetNsbHost()
	vc.logger.Info("adding default nsb ip ", "ip", vc.nsbHost)

	// todo: restrict scope
	vc.nsbClient = nsbcli.NewNSBClient(vc.nsbHost)
}
