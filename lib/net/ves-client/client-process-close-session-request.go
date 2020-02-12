package vesclient

import "github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"


func (vc *VesClient) ProcessCloseSessionRequest(req *wsrpc.CloseSessionRequest) {
	vc.logger.Info("session closed")
	vc.emitClose(req.SessionId)
}
