package vesclient

import "github.com/HyperService-Consortium/go-ves/grpc/wsrpc"

func (vc *VesClient) ProcessMessage(message []byte, messageID wsrpc.MessageType) {
	switch messageID {
	/* 0~4 */
	case wsrpc.CodeMessageRequest:
		//vc.processMessageRequest(message)
	case wsrpc.CodeMessageReply:
		vc.processMessageReply(message)
	case wsrpc.CodeRawProto:
		//vc.processRawProto(message)
	case wsrpc.CodeClientHelloRequest:
		//vc.processClientHelloRequest(message)
	case wsrpc.CodeClientHelloReply:
		vc.processClientHelloReply(message)

	/* 5~9 */
	case wsrpc.CodeRequestComingRequest:
		vc.processRequestComingRequest(message)
	case wsrpc.CodeRequestComingReply:
		//vc.processRequestComingReply(message)
	case wsrpc.CodeRequestGrpcServiceRequest:
		//vc.processRequestGrpcServiceRequest(message)
	case wsrpc.CodeRequestGrpcServiceReply:
		//vc.processRequestGrpcServiceReply(message)
	case wsrpc.CodeRequestNsbServiceRequest:
		//vc.processRequestNsbServiceRequest(message)

	/* 10~14 */
	case wsrpc.CodeRequestNsbServiceReply:
		//vc.processRequestNsbServiceReply(message)
	case wsrpc.CodeSessionListRequest:
		//vc.processSessionListRequest(message)
	case wsrpc.CodeSessionListReply:
		//vc.processSessionListReply(message)
	case wsrpc.CodeTransactionListRequest:
		//vc.processTransactionListRequest(message)
	case wsrpc.CodeTransactionListReply:
		//vc.processTransactionListReply(message)

	/* 15~19 */
	case wsrpc.CodeUserRegisterRequest:
		//vc.processUserRegisterRequest(message)
	case wsrpc.CodeUserRegisterReply:
		//vc.processUserRegisterReply(message)
		// todo: ignoring
		vc.logger.Info("user request request successfully")
	case wsrpc.CodeSessionFinishedRequest:
		//vc.processSessionFinishedRequest(message)
	case wsrpc.CodeSessionFinishedReply:
		//vc.processSessionFinishedReply(message)
	case wsrpc.CodeSessionRequireTransactRequest:
		//vc.processSessionRequireTransactRequest(message)

	/* 20~24 */
	case wsrpc.CodeSessionRequireTransactReply:
		//vc.processSessionRequireTransactReply(message)
	case wsrpc.CodeAttestationReceiveRequest:
		vc.processAttestationReceiveRequest(message)
	case wsrpc.CodeAttestationReceiveReply:
		//vc.processAttestationReceiveReply(message)
	case wsrpc.CodeAttestationSendingRequest:
		vc.processAttestationSendingRequest(message)
	case wsrpc.CodeAttestationSendingReply:
		//vc.processAttestationSendingReply(message)

	/* 25~26 */
	case wsrpc.CodeCloseSessionRequest:
		vc.processCloseSessionRequest(message)
	case wsrpc.CodeCloseSessionReply:
		//vc.processCloseSessionReply(message)
	default:
		// abort
		if !vc.ignoreUnknownMessage {
			vc.logger.Warn("aborting message", "id", messageID)
		}
	}
}

func (vc *VesClient) processMessageReply(message []byte) {
	vc.unmarshalProto(message, vc.getShortReplyMessage())
	// todo
	//if bytes.Equal(messageReply.To, vc.getName()) {
	//	vc.logger.Info("%v is saying: %v\n", "source", string(messageReply.From), "content", messageReply.Contents)
	//}
}

func (vc *VesClient) processClientHelloReply(message []byte) {
	if target := vc.getClientHelloReply(); vc.unmarshalProto(message, target) {
		vc.ProcessClientHelloReply(target)
	}
}

func (vc *VesClient) processRequestComingRequest(message []byte) {
	if target := vc.getRequestComingRequest(); vc.unmarshalProto(message, target) {
		vc.ProcessRequestComingRequest(target)
	}
}

func (vc *VesClient) processAttestationReceiveRequest(message []byte) {
	if target := vc.getAttestationReceiveRequest(); vc.unmarshalProto(message, target) {
		vc.ProcessAttestationReceiveRequest(target)
	}
}

func (vc *VesClient) processAttestationSendingRequest(message []byte) {
	// attestation sending request has the same format with request
	// coming request
	if target := vc.getRequestComingRequest(); vc.unmarshalProto(message, target) {
		vc.ProcessAttestationSendingRequest(target)
	}
}

func (vc *VesClient) processCloseSessionRequest(message []byte) {
	if target := vc.getCloseSessionRequest(); vc.unmarshalProto(message, target) {
		vc.ProcessCloseSessionRequest(target)
	}
}
