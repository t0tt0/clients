package centered_ves

import (
	"context"
	"github.com/Myriad-Dreamin/go-ves/central-ves/config"
	"github.com/Myriad-Dreamin/go-ves/central-ves/lib/plugin"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/fset"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
)

type CVESWebSocketPlugin = Server

func New() *CVESWebSocketPlugin {
	return &CVESWebSocketPlugin{}
}

func (c *CVESWebSocketPlugin) Configuration(logger plugin.Logger, loader plugin.ConfigLoader, cfg *plugin.ServerConfig) plugin.Plugin {
	//options := parseOptions(rOptions)
	return sugar.HandlerError(NewServer(
		cfg.BaseParametersConfig.RPCPort,
		cfg.BaseParametersConfig.WSPort,
		nil,
		logger,
		NSBHostOption(cfg.BaseParametersConfig.NSBHost))).(plugin.Plugin)
}

func (c *CVESWebSocketPlugin) Inject(services *plugin.ServiceProvider, dbs *plugin.DatabaseProvider, module plugin.Module) plugin.Plugin {
	c.userDB = module.Require(config.ModulePath.Global.UserDB).(*fset.AccountFSet)
	return c
}

func (c *CVESWebSocketPlugin) Work(ctx context.Context) {
	if err := c.Start(ctx); err != nil {
		c.logger.Error("work error", "error", err)
		return
	}
}
