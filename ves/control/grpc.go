package control

import (
	"github.com/Myriad-Dreamin/go-ves/grpc/uiprpc"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type GRPCEngine struct {
	*grpc.Server
	logger logger.Logger
	cfg *config.ServerConfig
}

func NewGRPCEngine(m module.Module) *GRPCEngine {
	s := grpc.NewServer()
	reflection.Register(s)
	return &GRPCEngine{
		Server: s,
		logger: m.Require(config.ModulePath.Minimum.Global.Logger).(logger.Logger),
	}
}

func (engine *GRPCEngine) Build(m Dependencies) error {
	uiprpc.RegisterVESServer(engine.Server,
		m.Require(config.ModulePath.Service.VESServer).(uiprpc.VESServer))
	return nil
}

func (engine *GRPCEngine) Run(port string) error {
	engine.logger.Info("prepare to serve rpc", "port", port)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	if err := engine.Server.Serve(lis); err != nil {
		return err
	}

	return nil
}
