package vesclient

import (
	ginhelper "github.com/Myriad-Dreamin/go-ves/lib/backend/gin-helper"
	"github.com/Myriad-Dreamin/go-ves/lib/backend/miris"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/kataras/iris"
	"net/http"
)

func (vc *VesClient) buildSessionRPCApis(p iris.Party) {
	//p.Get("session-list", vc.IrisListSession)
	id := p.Party("/session-list/{pid}")
	p.Post("/session", miris.ToIrisHandler(vc.IrisPostSession))
	id.Delete("", miris.ToIrisHandler(vc.IrisDeleteSession))
	id.Put("", miris.ToIrisHandler(vc.IrisPutSession))
}

func (vc *VesClient) IrisDeleteSession(c controller.MContext) {
	id, ok := ginhelper.ParseUint(c, "pid")
	if ok {
		c.JSON(http.StatusOK, iris.Map{
			"code": types.CodeOK,
			"id":   id,
		})
	}
	return
}

func (vc *VesClient) IrisPutSession(c controller.MContext) {

	c.JSON(http.StatusOK, iris.Map{
		"code": types.CodeToDo,
		"todo": 1,
	})
	return
}
