package userservice

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (srv *Service) deleteHook(c controller.MContext, object *model.User) (canDelete bool) {
	return true
}
