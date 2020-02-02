package centered_ves

import (
	"context"
	"github.com/Myriad-Dreamin/go-ves/central-ves/config"
	"github.com/Myriad-Dreamin/go-ves/central-ves/lib/plugin"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/instance"
	helper "github.com/Myriad-Dreamin/go-ves/lib/net/help-func"
	"net/http"
)

func New() *CVESWebSocketPlugin {
	return &CVESWebSocketPlugin{
		Server: new(http.Server),
	}
}

func (c *CVESWebSocketPlugin) Configuration(logger plugin.Logger, loader plugin.ConfigLoader, cfg *plugin.ServerConfig) plugin.Plugin {
	//options := parseOptions(rOptions)

	c.logger = logger
	var err error
	c.nsbip, err = helper.HostFromString(cfg.BaseParametersConfig.NSBHost)
	if err != nil {
		panic(err)
	}
	c.hub = newHub()
	c.hub.server = c
	c.Handler = http.NewServeMux()
	c.Addr = cfg.BaseParametersConfig.WSPort
	c.rpcPort = cfg.BaseParametersConfig.RPCPort
	return c
}

func (c *CVESWebSocketPlugin) Inject(services *plugin.ServiceProvider, dbs *plugin.DatabaseProvider, module plugin.Module) plugin.Plugin {
	c.userDB = module.Require(config.ModulePath.Global.UserDB).(*instance.VESInstance)
	return c
}

func (c *CVESWebSocketPlugin) Work(ctx context.Context) {
	if err := c.Start(ctx); err != nil {
		c.logger.Error("work error", "error", err)
		return
	}
}






