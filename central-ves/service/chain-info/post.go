package chainInfoservice

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	base_service "github.com/Myriad-Dreamin/go-ves/lib/base-service"
	ginhelper "github.com/Myriad-Dreamin/go-ves/lib/gin-helper"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (svc *Service) SerializePost(c controller.MContext) base_service.CRUDEntity {
	var req control.PostChainInfoRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}

	var obj = new(model.ChainInfo)
	// fill here
	obj.UserID = req.UserID
	obj.Address = req.Address
	obj.ChainID = req.ChainID
	return obj
}

func (svc *Service) AfterPost(obj interface{}) interface{} {
	return obj
}
