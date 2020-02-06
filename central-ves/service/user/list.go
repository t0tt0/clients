package userservice

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	types2 "github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (srv *Service) ProcessListResults(c controller.MContext, result interface{}) interface{} {
	return control.SerializeListUsersReply(
		types2.CodeOK,
		control.PackSerializeListUserReply(result.([]model.User)))
}
