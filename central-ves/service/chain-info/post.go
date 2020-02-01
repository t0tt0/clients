package chainInfoservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	base_service "github.com/Myriad-Dreamin/go-ves/central-ves/lib/base-service"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	ginhelper "github.com/Myriad-Dreamin/go-ves/central-ves/service/gin-helper"
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
