package userservice

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/control"
	"github.com/HyperService-Consortium/go-ves/central-ves/model"
	types2 "github.com/HyperService-Consortium/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (srv *Service) CreateEntity(id uint) interface{} {
	return &model.User{ID: id}
}

func (srv *Service) GetEntity(id uint) (interface{}, error) {
	return srv.userDB.ID(id)
}

func (srv *Service) CreateFilter() interface{} {
	return new(model.Filter)
}

func (srv *Service) ResponsePost(obj interface{}) interface{} {
	panic("abort")
	//return UserToPostReply(obj.(*model.User))
}

func (srv *Service) DeleteHook(c controller.MContext, obj interface{}) bool {
	return srv.deleteHook(c, obj.(*model.User))
}

func (srv *Service) ResponseGet(_ controller.MContext, obj interface{}) interface{} {
	return control.SerializeGetUserReply(types2.CodeOK, obj.(*model.User))
}

func (srv *Service) ResponseInspect(_ controller.MContext, obj interface{}) interface{} {
	return control.SerializeInspectUserReply(types2.CodeOK, obj.(*model.User))
}

func (srv *Service) GetPutRequest() interface{} {
	return new(control.PutUserRequest)
}

func (srv *Service) FillPutFields(c controller.MContext, user interface{}, req interface{}) []string {
	return srv.fillPutFields(c, user.(*model.User), req.(*control.PutUserRequest))
}

func (srv *Service) ListUsers(c controller.MContext) {
	srv.List(c)
	return
}

func (srv *Service) InspectUser(c controller.MContext) {
	srv.Inspect(c)
	return
}

func (srv *Service) GetUser(c controller.MContext) {
	srv.Get(c)
	return
}

func (srv *Service) PutUser(c controller.MContext) {
	srv.Put(c)
	return
}
