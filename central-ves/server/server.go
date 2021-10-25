package server

import (
	"context"
	"github.com/DeanThompson/ginpprof"
	"github.com/HyperService-Consortium/go-ves/central-ves/config"
	"github.com/HyperService-Consortium/go-ves/central-ves/control"
	"github.com/HyperService-Consortium/go-ves/central-ves/control/router"
	"github.com/HyperService-Consortium/go-ves/central-ves/lib/plugin"
	"github.com/HyperService-Consortium/go-ves/central-ves/model"
	"github.com/HyperService-Consortium/go-ves/central-ves/service"
	"github.com/HyperService-Consortium/go-ves/lib/backend/jwt"
	types2 "github.com/HyperService-Consortium/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"io"
	"os"
	"sync"
	"syscall"
)

type Server struct {
	Cfg          *config.ServerConfig
	Logger       types2.Logger
	LoggerWriter io.Writer

	RedisPool  *redis.Pool
	HttpEngine *control.HttpEngine
	Router     *router.RootRouter

	contestPath string

	jwtMW *jwt.Middleware
	//var authMW *privileger.MiddleWare
	routerAuthMW *controller.Middleware
	corsMW       gin.HandlerFunc

	Module          module.Module
	ServiceProvider *service.Provider
	ModelProvider   model.Provider
	RouterProvider  *router.Provider

	plugins []plugin.Plugin
}

func NewServer() *Server {
	return &Server{Module: make(module.Module)}
}

func (srv *Server) Terminate() {
	model.Close(srv.Module)
	syscall.Exit(0)
}

type Option interface {
	MinimumServerOption() bool
}

type OptionImpl struct{}

func (OptionImpl) MinimumServerOption() bool { return false }

type OptionRouterLoggerWriter struct {
	OptionImpl
	Writer io.Writer
}

func newServer(options []Option) *Server {
	srv := NewServer()

	for i := range options {
		switch option := options[i].(type) {
		case OptionRouterLoggerWriter:
			srv.LoggerWriter = option.Writer
		case *OptionRouterLoggerWriter:
			srv.LoggerWriter = option.Writer
		}
	}

	if srv.LoggerWriter == nil {
		srv.LoggerWriter = os.Stdout
	}
	srv.Module.Provide(config.ModulePath.Global.LoggerWriter, srv.LoggerWriter)

	srv.ModelProvider = model.NewProvider(config.ModulePath.Minimum.Provider.Model)
	srv.RouterProvider = router.NewProvider(config.ModulePath.Minimum.Provider.Router)
	srv.ServiceProvider = new(service.Provider)
	srv.HttpEngine = control.NewHttpEngine(srv.Module)

	srv.Module.Provide(config.ModulePath.Minimum.Provider.Service, srv.ServiceProvider)
	srv.Module.Provide(config.ModulePath.Minimum.Provider.Model, srv.ModelProvider)
	srv.Module.Provide(config.ModulePath.Minimum.Provider.Router, srv.RouterProvider)
	srv.Module.Provide(config.ModulePath.Global.UserDB, model.NewAccountFSet(srv.ModelProvider))
	return srv
}

func New(cfgPath string, options ...Option) (srv *Server) {
	srv = newServer(options)
	if !(srv.InstantiateLogger() &&
		srv.LoadConfig(cfgPath) &&
		srv.PrepareFileSystem() &&
		srv.PrepareDatabase()) {
		srv = nil
		return
	}
	defer func() {
		if err := recover(); err != nil {
			sugar.PrintStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		} else if srv == nil {
			srv.Terminate()
		}
	}()

	if !(srv.PrepareMiddleware() &&
		srv.PrepareService() &&
		srv.BuildRouter()) {
		srv = nil
		return
	}

	if err := srv.Module.Install(srv.RouterProvider); err != nil {
		srv.println("install router provider error", err)
	}
	if err := srv.Module.Install(srv.ModelProvider); err != nil {
		srv.println("install database provider error", err)
	}
	//
	//if !PreparePlugin(cfg) {
	//	srv = nil
	//return
	//}

	// Pressure()
	return
}




//inject() starts here
//
func (srv *Server) Inject(plugins ...plugin.Plugin) (injectSuccess bool) {
	defer func() {
		if err := recover(); err != nil {
			sugar.PrintStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		} else if injectSuccess == false {
			srv.Terminate()
		}
	}()

	for _, plg := range plugins {
// websocket center ves server is configured here		//
//configuration(): new websocket server in the plugin
		plg, _ = plg.Configuration(srv.Logger, srv.FetchConfig, srv.Cfg)
		if plg == nil {
			return false
		}
//inject(): inject the websocket server in the plugin to the center server
		plg, _ = plg.Inject(srv.ServiceProvider, srv.ModelProvider, srv.Module)
		if plg == nil {
			return false
		}
		srv.plugins = append(srv.plugins, plg)
	}
	return true
}
//inject() ends here





//http port is in the parameters
func (srv *Server) Serve(port string) {
	defer func() {
		if err := recover(); err != nil {
			sugar.PrintStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		}
	}()

	err := srv.HttpEngine.Build(srv.Module)
	if err != nil {
		panic(err)
	}
	srv.Module.Debug(srv.Logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	for _, plg := range srv.plugins {
		go plg.Work(ctx)
	}

	if err := model.GetRawInstance().Ping(); err != nil {
		srv.Logger.Debug("database died", "error", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := srv.HttpEngine.Run(port); err != nil {
			srv.Logger.Debug("IRouter run error", "error", err)
		}
		wg.Done()
	}()

	//do something
	wg.Wait()
}

func (srv *Server) ServeWithPProf(port string) {
	ginpprof.Wrap(srv.HttpEngine.Engine)
	srv.Serve(port)
}
