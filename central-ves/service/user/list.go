package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/go-ves/central-ves/types"
)

func (srv *Service) ProcessListResults(c controller.MContext, result interface{}) interface{} {
	return control.SerializeListUsersReply(
		types.CodeOK,
		control.PackSerializeListUserReply(result.([]model.User)))
}
