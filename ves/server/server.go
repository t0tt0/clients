package server

import (
	"context"
	"encoding/hex"
	"github.com/DeanThompson/ginpprof"
	base_account "github.com/HyperService-Consortium/go-uip/base-account"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	xconfig "github.com/Myriad-Dreamin/go-ves/config"
	"github.com/Myriad-Dreamin/go-ves/lib/jwt"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/go-ves/ves/control"
	"github.com/Myriad-Dreamin/go-ves/ves/control/router"
	"github.com/Myriad-Dreamin/go-ves/ves/lib/plugin"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
	dblayer "github.com/Myriad-Dreamin/go-ves/ves/model/db-layer"
	"github.com/Myriad-Dreamin/go-ves/ves/model/index"
	"github.com/Myriad-Dreamin/go-ves/ves/service"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"io"
	"os"
	"reflect"
	"sync"
	"syscall"
)

type Server struct {
	Cfg             *config.ServerConfig
	Logger          types.Logger
	Module          module.Module
	CloseHandler    types.CloseHandler
	ServiceProvider *service.Provider
	ModelProvider   *model.Provider
	RouterProvider  *router.Provider
	plugins         []plugin.Plugin

	LoggerWriter io.Writer
	levelDB      index.Engine
	RedisPool    *redis.Pool

	HTTPEngine *control.HttpEngine
	GRPCEngine *control.GRPCEngine
	Router     *router.RootRouter

	jwtMW *jwt.Middleware
	//var authMW *privileger.MiddleWare
	routerAuthMW *controller.Middleware
	corsMW       gin.HandlerFunc
}

func NewServer() *Server {
	return &Server{Module: make(module.Module)}
}

func (srv *Server) Terminate() {
	_ = srv.CloseHandler.Close()
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

type OptionCloseHandler struct {
	OptionImpl
	Handler types.CloseHandler
}

func newServer(options []Option) (srv *Server, err error) {
	srv = NewServer()

	for i := range options {
		switch option := options[i].(type) {
		case OptionRouterLoggerWriter:
			srv.LoggerWriter = option.Writer
		case *OptionRouterLoggerWriter:
			srv.LoggerWriter = option.Writer
		case OptionCloseHandler:
			srv.CloseHandler = option.Handler
		case *OptionCloseHandler:
			srv.CloseHandler = option.Handler
		}
	}

	if srv.LoggerWriter == nil {
		srv.LoggerWriter = os.Stdout
	}
	if srv.CloseHandler == nil {
		srv.CloseHandler = newCloseHandler()
	}

	if srv.Logger == nil {
		err = srv.instantiateLogger()
	}

	srv.Module.Provide(config.ModulePath.Global.CloseHandler, srv.CloseHandler)
	srv.Module.Provide(config.ModulePath.Global.LoggerWriter, srv.LoggerWriter)
	srv.Module.Provide(config.ModulePath.Service.ChainDNS, xconfig.ChainDNS)

	srv.ServiceProvider = new(service.Provider)
	srv.ModelProvider = model.NewProvider(config.ModulePath.Minimum.Provider.Model)
	srv.RouterProvider = router.NewProvider(config.ModulePath.Minimum.Provider.Router)
	srv.HTTPEngine = control.NewHttpEngine(srv.Module)
	srv.GRPCEngine = control.NewGRPCEngine(srv.Module)

	_ = model.SetProvider(srv.ModelProvider)
	srv.Module.Provide(config.ModulePath.Minimum.Provider.Service, srv.ServiceProvider)
	srv.Module.Provide(config.ModulePath.Minimum.Provider.Model, srv.ModelProvider)
	srv.Module.Provide(config.ModulePath.Minimum.Provider.Router, srv.RouterProvider)

	return
}

type CloseHandler struct {
	c []io.Closer
}

func (c *CloseHandler) Close() error {
	for i := range c.c {
		_ = c.c[i].Close()
	}
	return nil
}

func (c *CloseHandler) Handle(closer io.Closer) {
	c.c = append(c.c, closer)
}

func newCloseHandler() types.CloseHandler {
	return &CloseHandler{}
}

func (srv *Server) prepareBeforeLocalInit(cfgPath string) bool {

	if !(srv.InitRespAccount() &&
		srv.PrepareFileSystem() &&
		srv.PrepareRemoteService() &&
		srv.PrepareDatabase()) {
		return false
	}
	return true
}

func New(cfgPath string, options ...Option) (srv *Server, err error) {
	srv, err = newServer(options)
	if err != nil {
		return
	}
	if !srv.LoadConfig(cfgPath) {
		panic("build error")
	}

	if !srv.prepareBeforeLocalInit(cfgPath) {
		panic("build failed")
	}

	defer func() {
		if err := recover(); err != nil {
			srv.handlerPanicError(err)
		} else if srv == nil {
			srv.Terminate()
		}
	}()

	if !(srv.PrepareMiddleware() &&
		srv.PrepareService() &&
		srv.BuildRouter()) {
		panic("build failed")
		return
	}

	if err = srv.Module.Install(srv.RouterProvider); err != nil {
		return
	}
	if err = srv.Module.Install(srv.ModelProvider); err != nil {
		return
	}
	//
	//if !PreparePlugin(cfg) {
	//	panic("build failed")
	//return
	//}

	// Pressure()
	return
}

func (srv *Server) Inject(plugins ...plugin.Plugin) (injectSuccess bool) {
	defer func() {
		if err := recover(); err != nil {
			srv.handlerPanicError(err)
		} else if injectSuccess == false {
			srv.Terminate()
		}
	}()

	for _, plg := range plugins {
		plg = plg.Configuration(srv.Logger, srv.FetchConfig, srv.Cfg)
		if plg == nil {
			return false
		}
		plg = plg.Inject(srv.ServiceProvider, srv.ModelProvider, srv.Module)
		if plg == nil {
			return false
		}
		srv.plugins = append(srv.plugins, plg)
	}
	return true
}

func (srv *Server) Serve(httpPort, gRPCPort string) {

	// serve recover from panic
	defer func() {
		if err := recover(); err != nil {
			srv.handlerPanicError(err)
		}
	}()

	// lazy build
	sugar.HandlerError0(srv.HTTPEngine.Build(srv.Module))
	sugar.HandlerError0(srv.GRPCEngine.Build(srv.Module))

	srv.Cfg.BaseParametersConfig.ExposeHost = "127.0.0.1" + gRPCPort

	// all is ready
	srv.Module.Debug(srv.Logger)
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	for _, plg := range srv.plugins {
		go plg.Work(ctx)
	}

	if err := dblayer.GetRawInstance().Ping(); err != nil {
		srv.Logger.Debug("database died", "error", err)
		return
	}

	var wg sync.WaitGroup
	type pTask struct {
		engine control.RunnableEngine
		port   string
	}
	for _, task := range []pTask{
		{srv.HTTPEngine, httpPort},
		{srv.GRPCEngine, gRPCPort},
	} {
		wg.Add(1)
		go func(task pTask) {
			if err := task.engine.Run(task.port); err != nil {
				srv.Logger.Debug("run error",
					"error", err,
					"engine", reflect.TypeOf(task.engine), "port", task.port)
			}
			wg.Done()
		}(task)
	}
	//before engine built
	wg.Wait()
	// ensure engine built
}

func (srv *Server) ServeWithPProf(httpPort, gRPCPort string) {
	ginpprof.Wrap(srv.HTTPEngine.Engine)
	srv.Serve(httpPort, gRPCPort)
}

func (srv *Server) handlerPanicError(err interface{}) {
	sugar.PrintStack()
	srv.Logger.Error("panic error", "error", err)
	srv.Terminate()
}

func (srv *Server) InitRespAccount() bool {
	signer := sugar.HandlerError(signaturer.NewTendermintNSBSigner(
		sugar.HandlerError(
			hex.DecodeString(srv.Cfg.BaseParametersConfig.NSBSignerPrivateKey)).([]byte))).(uiptypes.Signer)
	srv.Module.Provide(config.ModulePath.Global.Signer, signer)
	////&uipbase.Account{Address: server.Signer.GetPublicKey(), ChainId: 3}
	srv.Module.Provide(config.ModulePath.Global.RespAccount, &base_account.Account{
		ChainId: srv.Cfg.BaseParametersConfig.NSBSignerChainID,
		Address: signer.GetPublicKey(),
	})
	srv.Logger.Info("using resp account",
		"chain-id", srv.Cfg.BaseParametersConfig.NSBSignerChainID,
		"public-address", hex.EncodeToString(signer.GetPublicKey()))
	return true
}
