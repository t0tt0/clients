package vesclient

import (
	"fmt"
	"github.com/HyperService-Consortium/go-ves/lib/backend/miris"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/kataras/iris"
	"net/http"
)

func (vc *VesClient) ListenHTTP(port string) error {
	r := iris.Default()
	r.Get("/ping", miris.ToIrisHandler(func(c controller.MContext) {
		c.JSON(http.StatusOK, iris.Map{"result": "pong"})
	}))

	v1 := r.Party("/v1")
	v1.PartyFunc("", vc.buildAccountRPCApis)
	v1.PartyFunc("", vc.buildSessionRPCApis)

	fmt.Println(r.GetRoutes())

	return r.Run(iris.Addr(port))
}

//	//		if err = vc.SendMessage(
//	//			bytes.TrimSpace(toBytes),
//	//			bytes.TrimSpace(buf.Bytes()),
//	//		); err != nil {
//	//			vc.logger.Error("error found", "error", err)
//	//			continue
//	//		}
