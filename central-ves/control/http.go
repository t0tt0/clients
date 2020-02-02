package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	mgin "github.com/Myriad-Dreamin/go-ves/central-ves/lib/gin"
	"github.com/gin-gonic/gin"
)

type HttpEngine = gin.Engine

func BuildHttp(router *controller.Router, engine *HttpEngine) {
	router.Build(mgin.NewGinRouter(engine))
}