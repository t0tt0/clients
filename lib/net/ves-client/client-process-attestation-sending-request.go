package vesclient

import (
	"encoding/hex"
	nsbcli "github.com/HyperService-Consortium/NSB/lib/nsb-client"
	TxState "github.com/HyperService-Consortium/go-uip/const/transaction_state_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc-base"
	"github.com/HyperService-Consortium/go-ves/grpc/wsrpc"
)

func (vc *VesClient) ProcessAttestationSendingRequest(attestationSendingRequest *wsrpc.RequestComingRequest) {

	vc.logger.Info(
		"new transaction's attestation must be created",
		"session id", hex.EncodeToString(attestationSendingRequest.GetSessionId()),
		"address", hex.EncodeToString(attestationSendingRequest.GetAccount().GetAddress()),
	)

	transactionReply, err := vc.GetRawTransaction(
		attestationSendingRequest.GetSessionId(),
		attestationSendingRequest.GetGrpcHost(),
	)
	if err != nil {
		vc.logger.Error("VesClient.read.AttestationSendingRequest.getRawTransaction", "error", err)
		return
	}

	//vc.logger.Info(
	//	"the instance of the transaction intent is",
	//	"tid", transactionReply.Tid,
	//	"tx", hex.EncodeToString(transactionReply.RawTransaction),
	//)

	//vc.logger.Info("the instance of the transaction")

	signer, err := vc.getNSBSigner()
	if err != nil {
		vc.logger.Error("VesClient.read.AttestationSendingRequest.getNSBSigner", "error", err)
		return
	}

	// packet attestation
	var sendingAtte = vc.getAttestationReceiveRequest()
	sendingAtte.SessionId = attestationSendingRequest.GetSessionId()
	sendingAtte.GrpcHost = attestationSendingRequest.GetGrpcHost()

	sigg, err := signer.Sign(transactionReply.RawTransaction)
	if err != nil {
		vc.logger.Error("VesClient.read.AttestationSendingRequest.Sign", "host", attestationSendingRequest.GetNsbHost(), "error", err)
		return
	}

	sendingAtte.Atte = &uiprpc_base.Attestation{
		Tid:     transactionReply.Tid,
		Aid:     TxState.Instantiating,
		Content: transactionReply.RawTransaction,
		Signatures: append(make([]*uiprpc_base.Signature, 0, 1), &uiprpc_base.Signature{
			// todo use src.nsbSigner to sign
			SignatureType: uip.SignatureTypeUnderlyingType(sigg.GetSignatureType()),
			Content:       sigg.GetContent(),
		}),
	}
	sendingAtte.Src = transactionReply.Src
	sendingAtte.Dst = transactionReply.Dst

	if ret, err := nsbcli.NewNSBClient(attestationSendingRequest.GetNsbHost()).InsuranceClaim( //delievered in this function i think
		signer,
		sendingAtte.SessionId,
		sendingAtte.Atte.Tid,
		TxState.Instantiating,
	); err != nil {
		vc.logger.Error("VesClient.read.AttestationSendingRequest.InsuranceClaim", "host", attestationSendingRequest.GetNsbHost(), "error", err)
		return
	} else {
		vc.logger.Info(
			"insurance claiming in sending, just send out",
			"tid", sendingAtte.Atte.Tid,
			"aid", TxState.Description(TxState.Instantiating),
			"info", ret.Info,
			"data", string(ret.Data),
			"log", ret.Log,
			"events", ret.Events,
		)
	}

	//if _, err := nsbcli.NewNSBClient(attestationSendingRequest.GetNsbHost()).InsuranceClaim(
	//	signer,
	//	sendingAtte.SessionId,
	//	sendingAtte.Atte.Tid,
	//	TxState.Instantiating,
	//); err != nil {
	//	vc.logger.Error("VesClient.read.AttestationSendingRequest.InsuranceClaim", "host", attestationSendingRequest.GetNsbHost(), "error", err)
	//	return
	//} else {
	//	//vc.logger.Info("insurance claiming")
	//}

	err = vc.PostRawMessage(wsrpc.CodeAttestationReceiveRequest, transactionReply.Dst, sendingAtte)
	if err != nil {
		vc.logger.Error("VesClient.read.AttestationSendingRequest.postRawMessage", "error", err)
		return
	}
	vc.logger.Info("post next attestation request successfully")
}
