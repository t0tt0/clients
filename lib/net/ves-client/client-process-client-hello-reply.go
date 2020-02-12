package vesclient

import (
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
	helper "github.com/Myriad-Dreamin/go-ves/lib/net/help-func"
	nsbcli "github.com/Myriad-Dreamin/go-ves/lib/net/nsb-client"
)


func (vc *VesClient) ProcessClientHelloReply(req *wsrpc.ClientHelloReply) {
	var err error
	vc.grpcip, err = helper.DecodeIP(req.GetGrpcHost())
	if err != nil {
		vc.logger.Error("VesClient.read.ClientHelloReply.decodeGRPCHost", "error", err)
	} else {
		vc.logger.Info("adding default grpc ip ", "ip", vc.grpcip)
	}

	vc.nsbip, err = helper.DecodeIP(req.GetNsbHost())
	if err != nil {
		vc.logger.Error("VesClient.read.ClientHelloReply.decodeNSBHost", "error", err)
	} else {
		vc.logger.Info("adding default nsb ip ", "ip", vc.nsbip)
	}

	// todo: restrict scope
	vc.nsbClient = nsbcli.NewNSBClient(vc.nsbip)
}
