package chainInfoservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	base_service "github.com/Myriad-Dreamin/go-ves/central-ves/lib/base-service"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/go-ves/central-ves/types"
)

func (svc *Service) CreateEntity(id uint) base_service.CRUDEntity {
	return &model.ChainInfo{ID: id}
}

func (svc *Service) GetEntity(id uint) (base_service.CRUDEntity, error) {
	return svc.db.ID(id)
}

func (svc *Service) ResponsePost(obj base_service.CRUDEntity) interface{} {
	return svc.AfterPost(control.SerializePostChainInfoReply(types.CodeOK, obj.(*model.ChainInfo)))
}

func (svc *Service) DeleteHook(c controller.MContext, obj base_service.CRUDEntity) bool {
	return svc.deleteHook(c, obj.(*model.ChainInfo))
}

func (svc *Service) ResponseGet(_ controller.MContext, obj base_service.CRUDEntity) interface{} {
	return control.SerializeGetChainInfoReply(types.CodeOK, obj.(*model.ChainInfo))
}

func (svc *Service) ResponseInspect(_ controller.MContext, obj base_service.CRUDEntity) interface{} {
	return control.SerializeInspectChainInfoReply(types.CodeOK, obj.(*model.ChainInfo))
}

func (svc *Service) ProcessListResults(_ controller.MContext, result interface{}) interface{} {
	return control.PSerializeListChainInfosReply(types.CodeOK, result.([]model.ChainInfo))
}

func (svc *Service) CreateFilter() interface{} {
	return new(model.Filter)
}

func (svc *Service) GetPutRequest() interface{} {
	return new(control.PutChainInfoRequest)
}

func (svc *Service) FillPutFields(c controller.MContext, chainInfo base_service.CRUDEntity, req interface{}) []string {
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
