package chainInfoservice

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/control"
	"github.com/HyperService-Consortium/go-ves/central-ves/model"
	ginhelper "github.com/HyperService-Consortium/go-ves/lib/backend/gin-helper"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (svc *Service) SerializePost(c controller.MContext) interface{} {
	var req control.PostChainInfoRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}

	var obj = new(model.ChainInfo)
	// fill here
	obj.UserID = req.UserId
	obj.Address = req.Address
	obj.ChainID = req.ChainId
	return obj
}

func (svc *Service) AfterPost(obj interface{}) interface{} {
	return obj
}
