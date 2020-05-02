package vesclient

import (
	"encoding/hex"
	nsbcli "github.com/HyperService-Consortium/NSB/lib/nsb-client"
	"github.com/HyperService-Consortium/go-ves/grpc/wsrpc"
)

func (vc *VesClient) ProcessRequestComingRequest(requestComingRequest *wsrpc.RequestComingRequest) {

	vc.logger.Info(
		"new session request coming",
		"session id", hex.EncodeToString(requestComingRequest.GetSessionId()),
		"responsible address", hex.EncodeToString(requestComingRequest.GetAccount().GetAddress()),
	)

	signer, err := vc.getNSBSigner()
	if err != nil {
		vc.logger.Error("VesClient.read.RequestComingRequest.getNSBSigner", "error", err)
		return
	}

	// todo: new nsbclient
	if ret, err := nsbcli.NewNSBClient(requestComingRequest.GetNsbHost()).UserAck(
		signer,
		requestComingRequest.GetSessionId(),
		requestComingRequest.GetAccount().GetAddress(),
		// todo: signature
		[]byte("123"),
	); err != nil {
		vc.logger.Error("VesClient.read.RequestComingRequest.UserAck", "host", requestComingRequest.GetNsbHost(), "error", err)
		return
	} else {
		vc.logger.Info(
			"user ack to nsb",
			"info", ret.Info, "data", string(ret.Data), "log", ret.Log, "events", ret.Events,
		)
	}

	x, err := signer.Sign(requestComingRequest.GetSessionId())
	if err != nil {
		vc.logger.Error("VesClient.read.RequestComingRequest.SignX", "error", err)
		return
	}
	if err = vc.sendAck(
		requestComingRequest.GetAccount(),
		requestComingRequest.GetSessionId(),
		requestComingRequest.GetGrpcHost(),
		x.Bytes(),
	); err != nil {
		vc.logger.Error("VesClient.read.RequestComingRequest.sendAck", "error", err)
		return
	} else {
		vc.logger.Info(
			"user ack to server",
		)
	}

}
