package vesclient

import (
	"github.com/Myriad-Dreamin/go-ves/lib/miris"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/kataras/iris"
	"net/http"
)

func (vc *VesClient) ListenHTTP() error {
	r := iris.Default()
	r.Get("/ping", miris.ToIrisHandler(func(c controller.MContext) {
		c.JSON(http.StatusOK, iris.Map{"result": "pong"})
	}))

	v1 := r.Party("/v1")
	v1.PartyFunc("", vc.buildAccountRPCApis)

	return r.Run(iris.Addr("0.0.0.0:26670"))
}
