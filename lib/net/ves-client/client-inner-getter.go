package vesclient

import (
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/HyperService-Consortium/go-ves/grpc/uiprpc-base"
	"github.com/HyperService-Consortium/go-ves/grpc/wsrpc"
)

func (vc *VesClient) getClientHello() *wsrpc.ClientHello {
	return new(wsrpc.ClientHello)
}

func (vc *VesClient) getClientHelloReply() *wsrpc.ClientHelloReply {
	return new(wsrpc.ClientHelloReply)
}

func (vc *VesClient) getShortSendMessage() *wsrpc.Message {
	return new(wsrpc.Message)
}

func (vc *VesClient) putShortSendMessage(_ *wsrpc.Message) {

}

func (vc *VesClient) getRawMessage() *wsrpc.RawMessage {
	return new(wsrpc.RawMessage)
}

func (vc *VesClient) combineRawMessage(dst *uiprpc_base.Account, code wsrpc.MessageType, b []byte) *wsrpc.RawMessage {
	var msg = vc.getRawMessage()
	msg.To = dst
	msg.MessageType = uint32(code)
	msg.Contents = b
	return msg
}

func (vc *VesClient) getShortReplyMessage() *wsrpc.Message {
	return new(wsrpc.Message)
}

func (vc *VesClient) getUserRegisterRequest() *wsrpc.UserRegisterRequest {
	return new(wsrpc.UserRegisterRequest)
}

func (vc *VesClient) getUserRegisterReply() *wsrpc.UserRegisterReply {
	return new(wsrpc.UserRegisterReply)
}

func (vc *VesClient) getRequestComingRequest() *wsrpc.RequestComingRequest {
	return new(wsrpc.RequestComingRequest)
}

func (vc *VesClient) getRequestComingReply() *wsrpc.RequestComingReply {
	return new(wsrpc.RequestComingReply)
}

func (vc *VesClient) getRequestGrpcServiceRequest() *wsrpc.RequestGrpcServiceRequest {
	return new(wsrpc.RequestGrpcServiceRequest)
}

func (vc *VesClient) getRequestGrpcServiceReply() *wsrpc.RequestGrpcServiceReply {
	return new(wsrpc.RequestGrpcServiceReply)
}

func (vc *VesClient) getRequestNsbServiceRequest() *wsrpc.RequestNsbServiceRequest {
	return new(wsrpc.RequestNsbServiceRequest)
}

func (vc *VesClient) getRequestNsbServiceReply() *wsrpc.RequestNsbServiceReply {
	return new(wsrpc.RequestNsbServiceReply)
}

func (vc *VesClient) getuserRegisterRequest() *wsrpc.UserRegisterRequest {
	return new(wsrpc.UserRegisterRequest)
}

func (vc *VesClient) getuserRegisterReply() *wsrpc.UserRegisterReply {
	return new(wsrpc.UserRegisterReply)
}

func (vc *VesClient) getSessionListRequest() *wsrpc.SessionListRequest {
	return new(wsrpc.SessionListRequest)
}

func (vc *VesClient) getSessionListReply() *wsrpc.SessionListReply {
	return new(wsrpc.SessionListReply)
}

func (vc *VesClient) getTransactionListRequest() *wsrpc.TransactionListRequest {
	return new(wsrpc.TransactionListRequest)
}

func (vc *VesClient) getTransactionListReply() *wsrpc.TransactionListReply {
	return new(wsrpc.TransactionListReply)
}

func (vc *VesClient) getSessionStart() *uiprpc.SessionStartRequest {
	return new(uiprpc.SessionStartRequest)
}

func (vc *VesClient) getSessionFinishedRequest() *wsrpc.SessionFinishedRequest {
	return new(wsrpc.SessionFinishedRequest)
}

func (vc *VesClient) getSessionFinishedReply() *wsrpc.SessionFinishedReply {
	return new(wsrpc.SessionFinishedReply)
}

func (vc *VesClient) getSendAttestationReceiveRequest() *wsrpc.AttestationReceiveRequest {
	return new(wsrpc.AttestationReceiveRequest)
}

func (vc *VesClient) combineSendAttestationReceiveRequest(
	src, dst *uiprpc_base.Account, attestation *uiprpc_base.Attestation,
	gRPCHost string, sessionId []byte,
) *wsrpc.AttestationReceiveRequest {
	req := vc.getSendAttestationReceiveRequest()
	req.SessionId = sessionId
	req.Src = src
	req.Dst = dst
	req.Atte = attestation
	req.GrpcHost = gRPCHost
	return req
}

func (vc *VesClient) getAttestationReceiveRequest() *wsrpc.AttestationReceiveRequest {
	return new(wsrpc.AttestationReceiveRequest)
}

func (vc *VesClient) getAttestationReceiveReply() *wsrpc.AttestationReceiveReply {
	return new(wsrpc.AttestationReceiveReply)
}

func (vc *VesClient) getCloseSessionRequest() *wsrpc.CloseSessionRequest {
	return new(wsrpc.CloseSessionRequest)
}

func (vc *VesClient) getVESHost() (string, error) {
	return vc.vesHost, nil
}
