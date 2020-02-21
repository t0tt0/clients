package objectservice

import (
	"github.com/Myriad-Dreamin/go-ves/lib/backend/gin-helper"
	"github.com/Myriad-Dreamin/go-ves/ves/control"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (svc *Service) SerializePost(c controller.MContext) interface{} {
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
