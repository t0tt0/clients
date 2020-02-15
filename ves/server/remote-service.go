package server

import (
	"fmt"
	base_account "github.com/HyperService-Consortium/go-uip/base-account"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	nsbcli "github.com/Myriad-Dreamin/go-ves/lib/net/nsb-client"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	centerAddress = "127.0.0.1:23352"
	nsbHost       = "39.100.145.91:26657"
)

func (srv *Server) PrepareRemoteService() bool {
	// lis2, err := net.Listen("tcp", ":33334")
	// if err != nil {
	// 	return fmt.Errorf("failed to listen: %v", err)
	// }

	conn, err := grpc.Dial(centerAddress, grpc.WithInsecure(), grpc.WithKeepaliveParams(keepalive.ClientParameters{}))
	if err != nil {
		srv.Logger.Error("did not connect", "error", err)
		return false
	}

	srv.Module.Require(config.ModulePath.Global.CloseHandler).(types.CloseHandler).Handle(conn)
	// conn.Close()

	srv.Module.Provide(config.ModulePath.Global.CentralVESClient, uiprpc.NewCenteredVESClient(conn))
	fmt.Println(srv.Module.Require(config.ModulePath.Global.CentralVESClient))
	srv.Module.Provide(config.ModulePath.Global.NSBClient, nsbcli.NewNSBClient(nsbHost))

	srv.Module.Provide(config.ModulePath.Global.Signer, sugar.HandlerError(signaturer.NewTendermintNSBSigner(make([]byte, 64))))
	srv.Module.Provide(config.ModulePath.Global.RespAccount, &base_account.Account{})
	return true
}
