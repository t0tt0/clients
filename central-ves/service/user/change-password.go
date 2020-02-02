package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	"github.com/Myriad-Dreamin/go-ves/central-ves/lib/serial"
	ginhelper "github.com/Myriad-Dreamin/go-ves/central-ves/service/gin-helper"
	"github.com/Myriad-Dreamin/go-ves/central-ves/types"
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
			Code: types.CodeWeakPassword,
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
