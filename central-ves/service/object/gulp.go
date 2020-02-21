package objectservice

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/control"
	"github.com/HyperService-Consortium/go-ves/central-ves/model"
	types2 "github.com/HyperService-Consortium/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (svc *Service) CreateEntity(id uint) interface{} {
	return &model.Object{ID: id}
}

func (svc *Service) GetEntity(id uint) (interface{}, error) {
	return svc.db.ID(id)
}

func (svc *Service) ResponsePost(obj interface{}) interface{} {
	return svc.AfterPost(control.SerializePostObjectReply(types2.CodeOK, obj.(*model.Object)))
}

func (svc *Service) DeleteHook(c controller.MContext, obj interface{}) bool {
	return svc.deleteHook(c, obj.(*model.Object))
}

func (svc *Service) ResponseGet(_ controller.MContext, obj interface{}) interface{} {
	return control.SerializeGetObjectReply(types2.CodeOK, obj.(*model.Object))
}

func (svc *Service) ResponseInspect(_ controller.MContext, obj interface{}) interface{} {
	return control.SerializeInspectObjectReply(types2.CodeOK, obj.(*model.Object))
}

func (svc *Service) ProcessListResults(_ controller.MContext, result interface{}) interface{} {
	return control.PSerializeListObjectsReply(types2.CodeOK, result.([]model.Object))
}

func (svc *Service) CreateFilter() interface{} {
	return new(model.Filter)
}

func (svc *Service) GetPutRequest() interface{} {
	return new(control.PutObjectRequest)
}

func (svc *Service) FillPutFields(c controller.MContext, object interface{}, req interface{}) []string {
	return svc.fillPutFields(c, object.(*model.Object), req.(*control.PutObjectRequest))
}

func (svc *Service) ListObjects(c controller.MContext) {
	svc.List(c)
	return
}

func (svc *Service) GetObject(c controller.MContext) {
	svc.Get(c)
	return
}

func (svc *Service) PutObject(c controller.MContext) {
	svc.Put(c)
	return
}

func (svc *Service) PostObject(c controller.MContext) {
	svc.Post(c)
	return
}

func (svc *Service) InspectObject(c controller.MContext) {
	svc.Inspect(c)
	return
}
