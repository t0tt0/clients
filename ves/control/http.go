package control

import (
	mgin "github.com/HyperService-Consortium/go-ves/lib/backend/gin"
	"github.com/HyperService-Consortium/go-ves/ves/config"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/gin-gonic/gin"
	"io"
)

type HttpEngine struct {
	*gin.Engine
}

func NewHttpEngine(m Dependencies) *HttpEngine {
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

func (h HttpEngine) Build(m Dependencies) error {
	h.Engine.Use(m.Require(config.ModulePath.Minimum.Middleware.CORS).(gin.HandlerFunc))

	router := m.Require(config.ModulePath.Global.Router).(*controller.Router)
	router.Build(mgin.NewGinRouter(h))
	return nil
}

func (h HttpEngine) Run(port string) error {
	return h.Engine.Run(port)
}
