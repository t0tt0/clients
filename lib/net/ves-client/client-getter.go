package vesclient

import (
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
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

func (vc *VesClient) getShortReplyMessage() *wsrpc.Message {
	return new(wsrpc.Message)
}

func (vc *VesClient) getUserRegisterRequest() *wsrpc.UserRegisterRequest {
	return new(wsrpc.UserRegisterRequest)
}

func (vc *VesClient) getUserRegisterReply() *wsrpc.UserRegisterReply {
	return new(wsrpc.UserRegisterReply)
}

func (vc *VesClient) getrequestComingRequest() *wsrpc.RequestComingRequest {
	return new(wsrpc.RequestComingRequest)
}

func (vc *VesClient) getrequestComingReply() *wsrpc.RequestComingReply {
	return new(wsrpc.RequestComingReply)
}

func (vc *VesClient) getrequestGrpcServiceRequest() *wsrpc.RequestGrpcServiceRequest {
	return new(wsrpc.RequestGrpcServiceRequest)
}

func (vc *VesClient) getrequestGrpcServiceReply() *wsrpc.RequestGrpcServiceReply {
	return new(wsrpc.RequestGrpcServiceReply)
}

func (vc *VesClient) getrequestNsbServiceRequest() *wsrpc.RequestNsbServiceRequest {
	return new(wsrpc.RequestNsbServiceRequest)
}

func (vc *VesClient) getrequestNsbServiceReply() *wsrpc.RequestNsbServiceReply {
	return new(wsrpc.RequestNsbServiceReply)
}

func (vc *VesClient) getuserRegisterRequest() *wsrpc.UserRegisterRequest {
	return new(wsrpc.UserRegisterRequest)
}

func (vc *VesClient) getuserRegisterReply() *wsrpc.UserRegisterReply {
	return new(wsrpc.UserRegisterReply)
}

func (vc *VesClient) getsessionListRequest() *wsrpc.SessionListRequest {
	return new(wsrpc.SessionListRequest)
}

func (vc *VesClient) getsessionListReply() *wsrpc.SessionListReply {
	return new(wsrpc.SessionListReply)
}

func (vc *VesClient) gettransactionListRequest() *wsrpc.TransactionListRequest {
	return new(wsrpc.TransactionListRequest)
}

func (vc *VesClient) gettransactionListReply() *wsrpc.TransactionListReply {
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

func (vc *VesClient) getReceiveAttestationReceiveRequest() *wsrpc.AttestationReceiveRequest {
	return new(wsrpc.AttestationReceiveRequest)
}

func (vc *VesClient) getAttestationReceiveReply() *wsrpc.AttestationReceiveReply {
	return new(wsrpc.AttestationReceiveReply)
}

func (vc *VesClient) getCloseSessionRequest() *wsrpc.CloseSessionRequest {
	return new(wsrpc.CloseSessionRequest)
}
