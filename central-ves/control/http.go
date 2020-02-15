package control

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/config"
	mgin "github.com/Myriad-Dreamin/go-ves/lib/gin"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/gin-gonic/gin"
	"io"
)

type HttpEngine struct {
	*gin.Engine
}

func NewHttpEngine(m module.Module) *HttpEngine {
	engine := gin.New()

	w := m.Require(config.ModulePath.Global.LoggerWriter).(io.Writer)
	gin.DefaultErrorWriter = w
	gin.DefaultWriter = w
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: w,
	}), gin.Recovery())

	m.Provide(config.ModulePath.Minimum.Global.HttpEngine, engine)
	return &HttpEngine{Engine: engine}
}

func (h HttpEngine) Build(m module.Module) error {
	h.Engine.Use(m.Require(config.ModulePath.Minimum.Middleware.CORS).(gin.HandlerFunc))

	router := m.Require(config.ModulePath.Global.Router).(*controller.Router)
	router.Build(mgin.NewGinRouter(h))
	return nil
}

func (h HttpEngine) Run(port string) error {
	return h.Engine.Run(port)
}
