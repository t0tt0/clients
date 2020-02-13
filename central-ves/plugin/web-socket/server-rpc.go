package centered_ves

import (
	"context"
	"encoding/hex"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func (c *Server) ListenAndServeRpc(_ context.Context, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		c.logger.Fatal("failed to listen", "error", err)
	}
	s := grpc.NewServer()
	uiprpc.RegisterCenteredVESServer(s, c)
	reflection.Register(s)

	c.logger.Info("prepare to serve rpc", "port", port)

	if err := s.Serve(lis); err != nil {
		c.logger.Fatal("failed to serve", "error", err)
	}
	return
}

func (c *Server) InternalRequestComing(
	ctx context.Context,
	in *uiprpc.InternalRequestComingRequest,
) (*uiprpc.InternalRequestComingReply, error) {
	if err := c.RequestComing(func() (accs []uiptypes.Account) {
		for _, acc := range in.GetAccounts() {
			accs = append(accs, acc)
		}
		return accs
	}(), in.GetSessionId(), in.GetHost()); err != nil {
		return nil, err
	}
	return &uiprpc.InternalRequestComingReply{
		Ok: true,
	}, nil
}

func (c *Server) InternalAttestationSending(
	ctx context.Context,
	in *uiprpc.InternalRequestComingRequest,
) (*uiprpc.InternalRequestComingReply, error) {
	if err := c.AttestationSending(func() (accs []uiptypes.Account) {
		for _, acc := range in.GetAccounts() {
			accs = append(accs, acc)
		}
		return accs
	}(), in.GetSessionId(), in.GetHost()); err != nil {
		return nil, err
	}
	return &uiprpc.InternalRequestComingReply{
		Ok: true,
	}, nil
}

// RequestComing do the service of retransmitting message of new session event
func (c *Server) RequestComing(accounts []uiptypes.Account, iscAddress []byte, grpcHost string) (err error) {
	// fmt.Println("rpc...", accounts)
	for _, acc := range accounts {
		// fmt.Println("hex", acc.GetChainId(), hex.EncodeToString(acc.GetAddress()))
		c.logger.Info("sending session request", "chain id", acc.GetChainId(), "address", hex.EncodeToString(acc.GetAddress()))
		if err = c.requestComing(acc, iscAddress, grpcHost); err != nil {
			return
		}
	}
	return nil
}

// AttestationSending do the service of retransmitting attestation
func (c *Server) AttestationSending(accounts []uiptypes.Account, iscAddress []byte, grpcHost string) (err error) {
	// fmt.Println("rpc...", accounts)
	for _, acc := range accounts {
		c.logger.Info("sending attestation request", "chain id", acc.GetChainId(), "address", hex.EncodeToString(acc.GetAddress()))
		if err = c.attestationSending(acc, iscAddress, grpcHost); err != nil {
			return
		}
	}
	return nil
}

func (c *Server) requestComing(acc uiptypes.Account, iscAddress []byte, grpcHost string) error {
	var msg wsrpc.RequestComingRequest
	msg.NsbHost = c.nsbip
	msg.GrpcHost = grpcHost
	msg.SessionId = iscAddress
	msg.Account = &uiprpc_base.Account{
		Address: acc.GetAddress(),
		ChainId: acc.GetChainId(),
	}
	c.hub.unicast <- &uniMessage{
		target: acc, task: newWriteMessageTask(
			wsrpc.CodeRequestComingRequest, &msg)}
	return nil
}

func (c *Server) attestationSending(acc uiptypes.Account, iscAddress []byte, grpcHost string) error {
	var msg wsrpc.RequestComingRequest
	msg.NsbHost = c.nsbip
	msg.GrpcHost = grpcHost
	msg.SessionId = iscAddress
	msg.Account = &uiprpc_base.Account{
		Address: acc.GetAddress(),
		ChainId: acc.GetChainId(),
	}

	// log.Infof("attestating network gate", )

	c.hub.unicast <- &uniMessage{target: acc, task: newWriteMessageTask(
		wsrpc.CodeAttestationSendingRequest, &msg)}
	return nil
}

func (c *Server) InternalCloseSession(
	ctx context.Context,
	in *uiprpc.InternalCloseSessionRequest,
) (*uiprpc.InternalCloseSessionReply, error) {
	if err := c.CloseSession(func() (accs []uiptypes.Account) {
		for _, acc := range in.GetAccounts() {
			accs = append(accs, acc)
		}
		return accs
	}(), in.GetSessionId(), in.GetGrpcHost(), in.GetNsbHost()); err != nil {
		return nil, err
	}
	return &uiprpc.InternalCloseSessionReply{
		Ok: true,
	}, nil
}

// CloseSession do the service of retransmitting attestation
func (c *Server) CloseSession(accounts []uiptypes.Account, iscAddress []byte, grpcHost, nsbHost string) (err error) {
	// fmt.Println("rpc...", accounts)
	for _, acc := range accounts {
		c.logger.Info("sending close session", "chain id", acc.GetChainId(), "address", hex.EncodeToString(acc.GetAddress()))
		if err = c.closeSession(acc, iscAddress, grpcHost, nsbHost); err != nil {
			return
		}
	}
	return nil
}

func (c *Server) closeSession(acc uiptypes.Account, iscAddress []byte, grpcHost, nsbHost string) error {
	var msg wsrpc.CloseSessionRequest
	msg.NsbHost = nsbHost
	msg.GrpcHost = grpcHost
	msg.SessionId = iscAddress

	c.hub.unicast <- &uniMessage{target: acc, task: newWriteMessageTask(
		wsrpc.CodeCloseSessionRequest,
		&msg,
	)}
	return nil
}
