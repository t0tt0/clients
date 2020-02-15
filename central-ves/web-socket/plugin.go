package central_ves

import (
	"context"
	"github.com/Myriad-Dreamin/go-ves/central-ves/config"
	"github.com/Myriad-Dreamin/go-ves/central-ves/lib/plugin"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/fset"
	"github.com/Myriad-Dreamin/go-ves/central-ves/web-socket/server"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
)

type CVESWebSocketPlugin struct {
	server.Server
}

func New() *CVESWebSocketPlugin {
	return &CVESWebSocketPlugin{}
}

func (srv *CVESWebSocketPlugin) Configuration(logger plugin.Logger, loader plugin.ConfigLoader, cfg *plugin.ServerConfig) plugin.Plugin {
	//options := parseOptions(rOptions)
	return sugar.HandlerError(server.NewServer(
		cfg.BaseParametersConfig.RPCPort,
		cfg.BaseParametersConfig.WSPort,
		nil,
		logger,
		server.NSBHostOption(cfg.BaseParametersConfig.NSBHost))).(plugin.Plugin)
}

func (srv *CVESWebSocketPlugin) Inject(services *plugin.ServiceProvider, dbs *plugin.DatabaseProvider, module plugin.Module) plugin.Plugin {
	srv.UserDB = module.Require(config.ModulePath.Global.UserDB).(*fset.AccountFSet)
	return srv
}

func (srv *CVESWebSocketPlugin) Work(ctx context.Context) {
	if err := srv.Start(ctx); err != nil {
		srv.Logger.Error("work error", "error", err)
		return
	}
}
