package control

import (
	mgin "github.com/Myriad-Dreamin/go-ves/lib/gin"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/gin-gonic/gin"
)

type HttpEngine = gin.Engine

func BuildHttp(router *controller.Router, engine *HttpEngine) {
	router.Build(mgin.NewGinRouter(engine))
}
