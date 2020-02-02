package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/go-ves/central-ves/control/auth"
	base_service "github.com/Myriad-Dreamin/go-ves/central-ves/lib/base-service"
	"github.com/Myriad-Dreamin/go-ves/central-ves/lib/serial"
	"github.com/Myriad-Dreamin/go-ves/central-ves/types"
)

type PostRequest struct {
}

func (srv *Service) SerializePost(c controller.MContext) base_service.CRUDEntity {
	panic("abort")
}

type PostReplyI interface {
	GetID() uint
}

func (srv *Service) AfterPost(reply PostReplyI) interface{} {
	if b, err := auth.UserEntity.AddReadPolicy(srv.enforcer, auth.UserEntity.CreateObj(reply.GetID()), reply.GetID()); err != nil {
		if !b {
			srv.logger.Debug("add failed")
		}
		return serial.ErrorSerializer{
			Code: types.CodeAddReadPrivilegeError,
			Err:  err.Error(),
		}
	} else {
		if !b {
			srv.logger.Debug("add failed")
		}
	}

	if b, err := auth.UserEntity.AddWritePolicy(srv.enforcer, auth.UserEntity.CreateObj(reply.GetID()), reply.GetID()); err != nil {
		if !b {
			srv.logger.Debug("add failed")
		}
		return serial.ErrorSerializer{
			Code: types.CodeAddWritePrivilegeError,
			Err:  err.Error(),
		}
	} else {
		if !b {
			srv.logger.Debug("add failed")
		}
	}
	return reply
}
