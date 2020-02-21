package userservice

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/control/auth"
	"github.com/HyperService-Consortium/go-ves/lib/backend/serial"
	types2 "github.com/HyperService-Consortium/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type PostRequest struct {
}

func (srv *Service) SerializePost(c controller.MContext) interface{} {
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
			Code: types2.CodeAddReadPrivilegeError,
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
			Code: types2.CodeAddWritePrivilegeError,
			Err:  err.Error(),
		}
	} else {
		if !b {
			srv.logger.Debug("add failed")
		}
	}
	return reply
}
