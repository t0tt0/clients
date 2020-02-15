package control

import (
	mgin "github.com/Myriad-Dreamin/go-ves/lib/gin"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/gin-gonic/gin"
	"io"
)

type HttpEngine = gin.Engine

func NewHttpEngine(m module.Module) *HttpEngine {
	engine := gin.New()

	w := m.Require(config.ModulePath.Global.LoggerWriter).(io.Writer)
	gin.DefaultErrorWriter = w
	gin.DefaultWriter = w
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: w,
	}), gin.Recovery())

	engine.Use(m.Require(config.ModulePath.Minimum.Middleware.CORS).(gin.HandlerFunc))

	m.Provide(config.ModulePath.Minimum.Global.HttpEngine, engine)
	return engine
}

func BuildHttp(router *controller.Router, engine *HttpEngine) {
	router.Build(mgin.NewGinRouter(engine))
}

