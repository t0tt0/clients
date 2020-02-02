// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package centered_ves

import (
	"context"
	"encoding/hex"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/instance"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	uiprpc_base "github.com/Myriad-Dreamin/go-ves/grpc/uiprpc-base"
	"github.com/Myriad-Dreamin/go-ves/grpc/wsrpc"
	"github.com/Myriad-Dreamin/go-ves/lib/net/help-func"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"

	"github.com/HyperService-Consortium/go-uip/uiptypes"
)

// func serveHome(w http.ResponseWriter, r *http.Request) {
// 	log.Println(r.URL)
// 	if r.URL.Path != "/" {
// 		http.Error(w, "Not found", http.StatusNotFound)
// 		return
// 	}
// 	if r.Method != "GET" {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	http.ServeFile(w, r, "home.html")
// }

// CVESWebSocketPlugin is a client manager, named centered ves
// it is not in the standard of uip
type CVESWebSocketPlugin struct {
	logger logger.Logger
	*http.Server
	hub     *Hub
	userDB  *instance.VESInstance
	rpcPort string
	nsbip   []byte
}

type NSBHostOption string

type ServerOptions struct {
	logger  logger.Logger
	nsbHost NSBHostOption
}

func defaultServerOptions() ServerOptions {
	return ServerOptions{
		logger:  logger.NewStdLogger(),
		nsbHost: "127.0.0.1:26657",
	}
}

func parseOptions(rOptions []interface{}) ServerOptions {
	var options = defaultServerOptions()
	for i := range rOptions {
		switch option := rOptions[i].(type) {
		case logger.Logger:
			options.logger = option
		case NSBHostOption:
			options.nsbHost = option
		}
	}
	return options
}

// NewServer return a pointer of CVESWebSocketPlugin
func NewServer(rpcport, addr string, db *instance.VESInstance, rOptions ...interface{}) (srv *CVESWebSocketPlugin, err error) {
	options := parseOptions(rOptions)
	srv = &CVESWebSocketPlugin{Server: new(http.Server)}
	srv.nsbip, err = helper.HostFromString(string(options.nsbHost))
	srv.hub = newHub()
	srv.hub.server = srv
	srv.userDB = db
	srv.Handler = http.NewServeMux()
	srv.Addr = addr
	srv.rpcPort = rpcport
	srv.logger = options.logger
	return
}

func (c *CVESWebSocketPlugin) ListenAndServeRpc(ctx context.Context, port string) {
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

// Start the service of centered ves
func (c *CVESWebSocketPlugin) Start(ctx context.Context) error {
	go c.hub.run(ctx)
	go c.ListenAndServeRpc(ctx, c.rpcPort)
	c.Handler.(*http.ServeMux).HandleFunc("/", c.serveWs)
	c.logger.Info("prepare to serve ws", "port", c.Addr)

	return c.ListenAndServe()
}

// serveWs handles websocket requests from the peer.
func (c *CVESWebSocketPlugin) serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		c.logger.Error("failed to upgrade to tcp", "error", err)
		return
	}

	c.logger.Info("new ws serving\n", "remote", r.RemoteAddr)
	client := &Client{hub: c.hub, helloed: make(chan bool, 1), conn: conn, send: make(chan *writeMessageTask, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

func (c *CVESWebSocketPlugin) InternalRequestComing(
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

func (c *CVESWebSocketPlugin) InternalAttestationSending(
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
func (c *CVESWebSocketPlugin) RequestComing(accounts []uiptypes.Account, iscAddress, grpcHost []byte) (err error) {
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
func (c *CVESWebSocketPlugin) AttestationSending(accounts []uiptypes.Account, iscAddress, grpcHost []byte) (err error) {
	// fmt.Println("rpc...", accounts)
	for _, acc := range accounts {
		c.logger.Info("sending attestation request", "chain id", acc.GetChainId(), "address", hex.EncodeToString(acc.GetAddress()))
		if err = c.attestationSending(acc, iscAddress, grpcHost); err != nil {
			return
		}
	}
	return nil
}

func (c *CVESWebSocketPlugin) requestComing(acc uiptypes.Account, iscAddress, grpcHost []byte) error {
	var msg wsrpc.RequestComingRequest
	msg.NsbHost = c.nsbip
	msg.GrpcHost = grpcHost
	msg.SessionId = iscAddress
	msg.Account = &uiprpc_base.Account{
		Address: acc.GetAddress(),
		ChainId: acc.GetChainId(),
	}
	packet, err := wsrpc.GetDefaultSerializer().Serial(wsrpc.CodeRequestComingRequest, &msg)
	if err != nil {
		return err
	}
	c.hub.unicast <- &uniMessage{
		target: acc, task: &writeMessageTask{
			b: packet.Bytes(), cb: func() {
				wsrpc.GetDefaultSerializer().Put(packet)
			},
		}}
	return nil
}

func (c *CVESWebSocketPlugin) attestationSending(acc uiptypes.Account, iscAddress, grpcHost []byte) error {
	var msg wsrpc.RequestComingRequest
	msg.NsbHost = c.nsbip
	msg.GrpcHost = grpcHost
	msg.SessionId = iscAddress
	msg.Account = &uiprpc_base.Account{
		Address: acc.GetAddress(),
		ChainId: acc.GetChainId(),
	}

	// log.Infof("attestating network gate", )

	packet, err := wsrpc.GetDefaultSerializer().Serial(wsrpc.CodeAttestationSendingRequest, &msg)
	if err != nil {
		return err
	}
	c.hub.unicast <- &uniMessage{target: acc, task: &writeMessageTask{
		b: packet.Bytes(), cb: func() {
			wsrpc.GetDefaultSerializer().Put(packet)
		}}}
	return nil
}

func (c *CVESWebSocketPlugin) InternalCloseSession(
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
func (c *CVESWebSocketPlugin) CloseSession(accounts []uiptypes.Account, iscAddress, grpcHost, nsbHost []byte) (err error) {
	// fmt.Println("rpc...", accounts)
	for _, acc := range accounts {
		c.logger.Info("sending close session", "chain id", acc.GetChainId(), "address", hex.EncodeToString(acc.GetAddress()))
		if err = c.closeSession(acc, iscAddress, grpcHost, nsbHost); err != nil {
			return
		}
	}
	return nil
}

func (c *CVESWebSocketPlugin) closeSession(acc uiptypes.Account, iscAddress, grpcHost, nsbHost []byte) error {
	var msg wsrpc.CloseSessionRequest
	msg.NsbHost = nsbHost
	msg.GrpcHost = grpcHost
	msg.SessionId = iscAddress

	packet, err := wsrpc.GetDefaultSerializer().Serial(wsrpc.CodeCloseSessionRequest, &msg)
	if err != nil {
		return err
	}
	c.hub.unicast <- &uniMessage{target: acc, task: &writeMessageTask{
		b: packet.Bytes(), cb: func() {
			wsrpc.GetDefaultSerializer().Put(packet)
		}}}
	return nil
}
