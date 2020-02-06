package userservice

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	ginhelper "github.com/Myriad-Dreamin/go-ves/lib/gin-helper"
	"github.com/Myriad-Dreamin/go-ves/lib/serial"
	types2 "github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
)

func (srv *Service) ChangePassword(c controller.MContext) {
	var req control.ChangePasswordRequest
	id, ok := ginhelper.ParseUintAndBind(c, "id", &req)
	if !ok {
		return
	}
	if sug := CheckStrongPassword(req.NewPassword); len(sug) != 0 {
		c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
			Code: types2.CodeWeakPassword,
			Err:  sug,
		})
		return
	}

	user, err := srv.userDB.Query(id)
	if ginhelper.MaybeSelectError(c, user, err) {
		return
	}

	if !ginhelper.AuthenticatePassword(c, user, req.OldPassword) {
		return
	}

	if ginhelper.ResetPassword(c, user, req.NewPassword) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	}
}
