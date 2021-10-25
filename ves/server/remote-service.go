package server

import (
	"fmt"
	nsbcli "github.com/HyperService-Consortium/NSB/lib/nsb-client"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	xconfig "github.com/HyperService-Consortium/go-ves/config"
	"github.com/HyperService-Consortium/go-ves/grpc/uiprpc"
	"github.com/HyperService-Consortium/go-ves/lib/bni/getter"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/ves/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	centerAddress = "127.0.0.1:23352"
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

	initializer, err := opintent.NewInitializer(xconfig.UserMap, getter.NewBlockChainGetter(xconfig.ChainDNS))
	if err != nil {
		srv.Logger.Error("init op intent initializer error", "error", err)
		return false
	}

	srv.Module.Provide(config.ModulePath.Service.OpIntentInitializer, initializer)
	srv.Module.Provide(config.ModulePath.Global.CentralVESClient, uiprpc.NewCenteredVESClient(conn))
	fmt.Println(srv.Module.Require(config.ModulePath.Global.CentralVESClient))
	srv.Module.Provide(config.ModulePath.Global.NSBClient, nsbcli.NewNSBClient(srv.Cfg.BaseParametersConfig.NSBHost))
	return true
}
