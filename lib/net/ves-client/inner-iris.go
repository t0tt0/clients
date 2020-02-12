package vesclient

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/kataras/iris"
)

func BindJSON(ctx iris.Context, obj interface{}) error {
	return binding.JSON.Bind(ctx.Request(), obj)
}

func BindQuery(ctx iris.Context, obj interface{}) error {
	return binding.Query.Bind(ctx.Request(), obj)
}

func Bind(ctx iris.Context, obj interface{}) error {
	return BindJSON(ctx, obj)
}

func (vc *VesClient) ContextJSON(ctx iris.Context, v interface{}) {
	_, err := ctx.JSON(v)
	if err != nil {
		vc.logger.Error("serialize json error", "error", err)
	}
}
