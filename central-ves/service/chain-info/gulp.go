package chainInfoservice

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	types2 "github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)


func (svc *Service) CreateEntity(id uint) interface{} {
	return &model.ChainInfo{ID: id}
}

func (svc *Service) GetEntity(id uint) (interface{}, error) {
	return svc.db.ID(id)
}

func (svc *Service) ResponsePost(obj interface{}) interface{} {
	return svc.AfterPost(control.SerializePostChainInfoReply(types2.CodeOK, obj.(*model.ChainInfo)))
}

func (svc *Service) DeleteHook(c controller.MContext, obj interface{}) bool {
	return svc.deleteHook(c, obj.(*model.ChainInfo))
}

func (svc *Service) ResponseGet(_ controller.MContext, obj interface{}) interface{} {
	return control.SerializeGetChainInfoReply(types2.CodeOK, obj.(*model.ChainInfo))
}

func (svc *Service) ResponseInspect(_ controller.MContext, obj interface{}) interface{} {
	return control.SerializeInspectChainInfoReply(types2.CodeOK, obj.(*model.ChainInfo))
}

func (svc *Service) ProcessListResults(_ controller.MContext, result interface{}) interface{} {
	return control.PSerializeListChainInfosReply(types2.CodeOK, result.([]model.ChainInfo))
}

func (svc *Service) CreateFilter() interface{} {
	return new(model.Filter)
}

func (svc *Service) GetPutRequest() interface{} {
	return new(control.PutChainInfoRequest)
}

func (svc *Service) FillPutFields(c controller.MContext, chainInfo interface{}, req interface{}) []string {
	return svc.fillPutFields(c, chainInfo.(*model.ChainInfo), req.(*control.PutChainInfoRequest))
}

func (svc *Service) ListChainInfos(c controller.MContext) {
	svc.List(c)
	return
}

func (svc *Service) GetChainInfo(c controller.MContext) {
	svc.Get(c)
	return
}

func (svc *Service) PutChainInfo(c controller.MContext) {
	svc.Put(c)
	return
}

func (svc *Service) PostChainInfo(c controller.MContext) {
	svc.Post(c)
	return
}

func (svc *Service) InspectChainInfo(c controller.MContext) {
	svc.Inspect(c)
	return
}
