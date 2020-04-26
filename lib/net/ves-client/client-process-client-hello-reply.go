package vesclient

import (
	nsbcli "github.com/HyperService-Consortium/NSB/lib/nsb-client"
	"github.com/HyperService-Consortium/go-ves/grpc/wsrpc"
)

func (vc *VesClient) ProcessClientHelloReply(req *wsrpc.ClientHelloReply) {
	vc.vesHost = req.GetGrpcHost()
	vc.logger.Info("adding default grpc ip ", "ip", vc.vesHost)

	vc.nsbHost = req.GetNsbHost()
	vc.logger.Info("adding default nsb ip ", "ip", vc.nsbHost)

	// todo: restrict scope
	vc.nsbClient = nsbcli.NewNSBClient(vc.nsbHost)
}
