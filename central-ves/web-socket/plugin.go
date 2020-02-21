package central_ves

import (
	"context"
	"github.com/HyperService-Consortium/go-ves/central-ves/config"
	"github.com/HyperService-Consortium/go-ves/central-ves/lib/plugin"
	"github.com/HyperService-Consortium/go-ves/central-ves/model"
	"github.com/HyperService-Consortium/go-ves/central-ves/web-socket/chs"
)

type CVESWebSocketPlugin struct {
	*chs.Server
}

func New() *CVESWebSocketPlugin {
	return &CVESWebSocketPlugin{}
}

func (srv *CVESWebSocketPlugin) Configuration(logger plugin.Logger, loader plugin.ConfigLoader, cfg *plugin.ServerConfig) (p plugin.Plugin, err error) {
	//options := parseOptions(rOptions)
	p = srv
	srv.Server, err = chs.NewServer(
		cfg.BaseParametersConfig.RPCPort,
		cfg.BaseParametersConfig.WSPort,
		nil,
		logger,
		chs.NSBHostOption(cfg.BaseParametersConfig.NSBHost))
	return
}

func (srv *CVESWebSocketPlugin) Inject(services *plugin.ServiceProvider, dbs plugin.DatabaseProvider, module plugin.Module) (plugin.Plugin, error) {
	srv.ProvideUserDB(module.Require(config.ModulePath.Global.UserDB).(*model.AccountFSet))
	return srv, nil
}

func (srv *CVESWebSocketPlugin) Work(ctx context.Context) error {
	if err := srv.Start(ctx); err != nil {
		srv.Logger.Error("work error", "error", err)
		return err
	}
	return nil
}
