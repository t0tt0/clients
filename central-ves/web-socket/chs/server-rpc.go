package chs

import (
	"context"
	"encoding/hex"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/HyperService-Consortium/go-ves/grpc/uiprpc-base"
	"github.com/HyperService-Consortium/go-ves/grpc/wsrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func (srv *Server) ListenAndServeRpc(_ context.Context, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		srv.Logger.Fatal("failed to listen", "error", err)
	}
	s := grpc.NewServer()
	uiprpc.RegisterCenteredVESServer(s, srv)
	reflection.Register(s)

	srv.Logger.Info("prepare to serve rpc", "port", port)

	if err := s.Serve(lis); err != nil {
		srv.Logger.Fatal("failed to serve", "error", err)
	}
	return
}

func (srv *Server) InternalRequestComing(
	ctx context.Context,
	in *uiprpc.InternalRequestComingRequest,
) (*uiprpc.InternalRequestComingReply, error) {
	if err := srv.RequestComing(func() (accs []uip.Account) {
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

func (srv *Server) InternalAttestationSending(
	ctx context.Context,
	in *uiprpc.InternalRequestComingRequest,
) (*uiprpc.InternalRequestComingReply, error) {
	if err := srv.AttestationSending(func() (accs []uip.Account) {
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
func (srv *Server) RequestComing(accounts []uip.Account, iscAddress []byte, grpcHost string) (err error) {
	// fmt.Println("rpc...", accounts)
	for _, acc := range accounts {
		// fmt.Println("hex", acc.GetChainId(), hex.EncodeToString(acc.GetAddress()))
		srv.Logger.Info("sending session request", "chain id", acc.GetChainId(), "address", hex.EncodeToString(acc.GetAddress()))
		if err = srv.requestComing(acc, iscAddress, grpcHost); err != nil {
			return
		}
	}
	return nil
}

// AttestationSending do the service of retransmitting attestation
func (srv *Server) AttestationSending(accounts []uip.Account, iscAddress []byte, grpcHost string) (err error) {
	// fmt.Println("rpc...", accounts)
	for _, acc := range accounts {
		srv.Logger.Info("sending attestation request", "chain id", acc.GetChainId(), "address", hex.EncodeToString(acc.GetAddress()))
		if err = srv.attestationSending(acc, iscAddress, grpcHost); err != nil {
			return
		}
	}
	return nil
}

func (srv *Server) requestComing(acc uip.Account, iscAddress []byte, grpcHost string) error {
	var msg wsrpc.RequestComingRequest
	msg.NsbHost = srv.Nsbip
	msg.GrpcHost = grpcHost
	msg.SessionId = iscAddress
	msg.Account = &uiprpc_base.Account{
		Address: acc.GetAddress(),
		ChainId: acc.GetChainId(),
	}
	srv.hub.Unicast <- &UniMessage{
		Target: acc, Task: NewWriteMessageTask(
			wsrpc.CodeRequestComingRequest, &msg)}
	return nil
}

func (srv *Server) attestationSending(acc uip.Account, iscAddress []byte, grpcHost string) error {
	var msg wsrpc.RequestComingRequest
	msg.NsbHost = srv.Nsbip
	msg.GrpcHost = grpcHost
	msg.SessionId = iscAddress
	msg.Account = &uiprpc_base.Account{
		Address: acc.GetAddress(),
		ChainId: acc.GetChainId(),
	}

	// log.Infof("attestating network gate", )

	srv.hub.Unicast <- &UniMessage{Target: acc, Task: NewWriteMessageTask(
		wsrpc.CodeAttestationSendingRequest, &msg)}
	return nil
}

func (srv *Server) InternalCloseSession(
	ctx context.Context,
	in *uiprpc.InternalCloseSessionRequest,
) (*uiprpc.InternalCloseSessionReply, error) {
	if err := srv.CloseSession(func() (accs []uip.Account) {
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
func (srv *Server) CloseSession(accounts []uip.Account, iscAddress []byte, grpcHost, nsbHost string) (err error) {
	// fmt.Println("rpc...", accounts)
	for _, acc := range accounts {
		srv.Logger.Info("sending close session", "chain id", acc.GetChainId(), "address", hex.EncodeToString(acc.GetAddress()))
		if err = srv.closeSession(acc, iscAddress, grpcHost, nsbHost); err != nil {
			return
		}
	}
	return nil
}

func (srv *Server) closeSession(acc uip.Account, iscAddress []byte, grpcHost, nsbHost string) error {
	var msg wsrpc.CloseSessionRequest
	msg.NsbHost = nsbHost
	msg.GrpcHost = grpcHost
	msg.SessionId = iscAddress

	srv.hub.Unicast <- &UniMessage{Target: acc, Task: NewWriteMessageTask(
		wsrpc.CodeCloseSessionRequest,
		&msg,
	)}
	return nil
}
