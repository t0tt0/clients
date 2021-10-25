package plugin

import (
	"context"
	types2 "github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/ves/config"
	"github.com/HyperService-Consortium/go-ves/ves/model"
	"github.com/HyperService-Consortium/go-ves/ves/service"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Logger = types2.Logger
type ConfigLoader = types2.ConfigLoader
type ServiceProvider = service.Provider
type DatabaseProvider = model.Provider
type ServerConfig = config.ServerConfig
type Module = module.Module

type Plugin interface {
	Configuration(logger Logger, loader ConfigLoader, cfg *ServerConfig) (plg Plugin)
	Inject(services *ServiceProvider, dbs DatabaseProvider, module Module) (plg Plugin)
	Work(ctx context.Context)
}
