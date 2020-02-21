package plugin

import (
	"context"
	"github.com/HyperService-Consortium/go-ves/central-ves/config"
	"github.com/HyperService-Consortium/go-ves/central-ves/model"
	"github.com/HyperService-Consortium/go-ves/central-ves/service"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Logger = types.Logger
type ConfigLoader = types.ConfigLoader
type ServiceProvider = service.Provider
type DatabaseProvider = model.Provider
type ServerConfig = config.ServerConfig
type Module = module.Module

type Plugin interface {
	Configuration(logger Logger, loader ConfigLoader, cfg *ServerConfig) (plg Plugin, err error)
	Inject(services *ServiceProvider, dbs DatabaseProvider, module Module) (plg Plugin, err error)
	Work(ctx context.Context) (err error)
}
