package userservice

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/control"
	"github.com/HyperService-Consortium/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (srv *Service) fillPutFields(c controller.MContext, user *model.User, req *control.PutUserRequest) (fields []string) {
	return nil
}
