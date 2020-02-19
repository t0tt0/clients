package vesclient

import (
	"encoding/hex"
	"github.com/HyperService-Consortium/go-uip/const/transaction_state_type"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
)

type attestationReceiveRequestService struct {
	client *VesClient
	req    *wsrpc.AttestationReceiveRequest

	signer uiptypes.Signer
	newReq *wsrpc.AttestationReceiveRequest
}

func (vc *VesClient) ProcessAttestationReceiveRequest(
	req *wsrpc.AttestationReceiveRequest) {
	svc := attestationReceiveRequestService{client: vc, req: req}
	(&svc).serve()
}

func (svc *attestationReceiveRequestService) serve() {
	attestation := svc.req.GetAtte()
	if attestation == nil {
		svc.client.logger.Error("omit request without attestation")
		return
	}

	switch attestation.Aid {
	/* 0~4 */
	case TxState.Unknown:
		svc.procTxStateUnknown()
	case TxState.Initing:
		svc.procTxStateInitializing()
	case TxState.Inited:
		svc.procTxStateInitialized()
	case TxState.Instantiating:
		svc.procTxStateInstantiating()
	case TxState.Instantiated:
		svc.procTxStateInstantiated()

	/* 5~7 */
	case TxState.Open:
		svc.procTxStateOpen()
	case TxState.Opened:
		svc.procTxStateOpened()
	case TxState.Closed:
		svc.procTxStateClosed()

	default:
		svc.client.logger.Warn("received unknown tx state", "state", attestation.Aid)
	}
}

func (svc *attestationReceiveRequestService) procTxStateUnknown() {
	svc.client.logger.Info(
		"transaction is of the status unknown",
		"tid", svc.req.Atte.Tid)
}

func (svc *attestationReceiveRequestService) procTxStateInitializing() {
	svc.client.logger.Info(
		"transaction is of the status initializing",
		"tid", svc.req.Atte.Tid)
}

func (svc *attestationReceiveRequestService) procTxStateInitialized() {
	svc.client.logger.Info(
		"transaction is of the status initialized",
		"tid", svc.req.Atte.Tid)
}

func (svc *attestationReceiveRequestService) procTxStateInstantiating() {
	if svc.newReq = svc.generateNewAttestationFromOld(); svc.newReq != nil {
		svc.tellOthers()
	}
}

func (svc *attestationReceiveRequestService) procTxStateInstantiated() {
	if svc.newReq = svc.generateNewAttestationFromOld(); svc.newReq != nil {
		svc.tellOthers()
	}
}

func (svc *attestationReceiveRequestService) procTxStateOpen() {
	if svc.newReq = svc.generateNewAttestationFromOld(); svc.newReq != nil && svc.doTransaction() {
		svc.tellOthers()
	}
}

func (svc *attestationReceiveRequestService) procTxStateOpened() {
	if svc.newReq = svc.generateNewAttestationFromOld(); svc.newReq != nil {
		svc.tellOthers()
	}
}

func (svc *attestationReceiveRequestService) procTxStateClosed() {
	svc.client.logger.Info(
		"transaction is of the status closed",
		"tid", svc.req.GetAtte().Tid)
}

func (svc *attestationReceiveRequestService) generateNewAttestationFromOld() *wsrpc.AttestationReceiveRequest {
	oldAttestation := svc.req.Atte

	svc.client.logger.Info(
		"generate new attestation with status",
		"tid", oldAttestation.Tid,
		"aid", TxState.Description(oldAttestation.Aid+1))

	if signature := svc.checkSignaturesAndSign(
		oldAttestation.GetSignatures()); signature == nil {
		return nil
	} else {
		return svc.client.combineSendAttestationReceiveRequest(
			svc.req.GetDst(), svc.req.GetSrc(),
			nextAttestation(oldAttestation, SignatureFromSTDToRPC(signature)),
			svc.req.GetGrpcHost(),
			svc.req.GetSessionId())
	}
}

func (svc *attestationReceiveRequestService) tellOthers() {
	//todo: move logger errors
	if !svc.client.ensureGetNSBSigner(&svc.signer) {
		return
	}
	ret, err := svc.client.nsbClient.InsuranceClaim(
		svc.signer,
		svc.newReq.SessionId,
		svc.newReq.Atte.Tid, svc.newReq.Atte.Aid,
	)
	if err != nil {
		svc.client.logger.Error("InsuranceClaim", "error", err)
		return
	}

	svc.client.logger.Info(
		"insurance claiming",
		"tid", svc.newReq.Atte.Tid,
		"aid", TxState.Description(svc.newReq.Atte.Aid),
		"info", ret.Info,
		"data", string(ret.Data),
		"log", ret.Log,
		"tags", ret.Tags,
	)

	err = svc.client.RetransmitAttestationReceiveRequest(svc.newReq.Dst, svc.newReq)
	if err != nil {
		svc.client.logger.Error("postRawMessage", "error", err)
		return
	}

	svc.client.informAttestation(svc.req.GrpcHost, svc.newReq)
}

func (svc *attestationReceiveRequestService) checkSignaturesAndSign(
	signatures []*uiprpc_base.Signature) (signature uiptypes.Signature) {
	if !svc.client.ensureGetNSBSigner(&svc.signer) {
		return
	}

	// todo: [iter the attestation (copy or refer it? ), before 2.15] -> [check?]
	var err error
	signature, err = svc.signer.Sign(signatures[len(signatures)-1].GetContent())
	if err != nil {
		svc.client.logger.Error("sign chain error", "error", err)
		signature = nil
		return
	}
	return
}

type _ARRDoTransactionService struct {
	*attestationReceiveRequestService

	router uiptypes.Router
	translator uiptypes.Translator
}

func (svc *attestationReceiveRequestService) doTransaction() bool {
	acc := svc.req.Dst
	svc.client.logger.Info(
		"the resp is",
		"address", hex.EncodeToString(acc.GetAddress()),
		"chain id", acc.GetChainId())

	ctx := _ARRDoTransactionService{
		attestationReceiveRequestService: svc}

	if !svc.client.ensureRouter(acc.ChainId, &ctx.router) {
		return false
	}

	if !svc.client.ensureTranslator(acc.ChainId, &ctx.translator) {
		return false
	}

	return (&ctx).doTransaction()
}

func (svc *_ARRDoTransactionService) doTransaction() bool {
	acc := svc.req.Dst

	if svc.router.MustWithSigner() && !svc.decorateRouterWithRespSigner(acc) {
		return false
	}

	rawTx, err := svc.translator.Deserialize(svc.req.Atte.Content)
	if err != nil {
		svc.client.logger.Error("translator.Deserialize", "error", err)
		return false
	}

	receipt, err := svc.router.RouteRaw(acc.ChainId, rawTx)
	if err != nil {
		svc.client.logger.Error("router.RouteRaw", "error", err)
		return false
	}
	svc.client.logger.Info("routing", "receipt", hex.EncodeToString(receipt))

	bid, additional, err := svc.router.WaitForTransact(acc.ChainId, receipt, svc.client.waitOpt)
	if err != nil {
		svc.client.logger.Error("router.WaitForTransact", "error", err)
		return false
	}
	svc.client.logger.Info("route result", "block id", bid)

	blockStorage, err := svc.client.getBlockStorage(acc.ChainId)
	if err != nil {
		svc.client.logger.Error("getBlockStorage:", "error", err)
		return false
	}

	proof, err := blockStorage.GetTransactionProof(acc.GetChainId(), bid, additional)
	if err != nil {
		svc.client.logger.Error("blockStorage.GetTransactionProof", "error", err)
		return false
	}

	if !svc.client.ensureGetNSBSigner(&svc.signer) {
		return false
	}
	cb, err := svc.client.nsbClient.AddMerkleProof(
		svc.signer, nil, uiptypes.MerkleProofTypeUnderlyingType(proof.GetType()), proof.GetRootHash(), proof.GetProof(), proof.GetKey(), proof.GetValue())
	if err != nil {
		svc.client.logger.Error("nsbClient.AddMerkleProof", "error", err)
		return false
	}
	svc.client.logger.Info("adding merkle proof", "result", cb)

	// todo: add const TransactionsRoot
	cb, err = svc.client.nsbClient.AddBlockCheck(
		svc.signer, nil, acc.ChainId, bid, proof.GetRootHash(), 1)
	if err != nil {
		svc.client.logger.Error("nsbClient.AddBlockCheck", "error", err)
		return false
	}
	svc.client.logger.Info("adding block check", "result", cb)

	// todo XXXXXXXXXXXXXXXXXXX
	return true
}

func (svc *_ARRDoTransactionService) decorateRouterWithRespSigner(
	acc *uiprpc_base.Account) bool {
	respSigner, err := svc.client.getRespSigner(acc)
	if err != nil {
		svc.client.logger.Error(
			"getRespSigner error",
			"error", err,
			"chainID", acc.ChainId,
			"account address", hex.EncodeToString(acc.Address))
		return false
	}

	svc.router, err = svc.router.RouteWithSigner(respSigner)
	if err != nil {
		svc.client.logger.Error("RouteWithSigner", "error", err)
		return false
	}
	return true
}
