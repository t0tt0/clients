package objectservice

import (
	base_service "github.com/Myriad-Dreamin/go-ves/lib/base-service"
	"github.com/Myriad-Dreamin/go-ves/lib/gin-helper"
	"github.com/Myriad-Dreamin/go-ves/vesx/control"
	"github.com/Myriad-Dreamin/go-ves/vesx/model"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (svc *Service) SerializePost(c controller.MContext) base_service.CRUDEntity {
	var req control.PostObjectRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}

	var obj = new(model.Object)
	// fill here
	return obj
}

func (svc *Service) AfterPost(obj interface{}) interface{} {
	return obj
}
