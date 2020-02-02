package userservice

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	"github.com/Myriad-Dreamin/go-ves/central-ves/lib/serial"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	ginhelper "github.com/Myriad-Dreamin/go-ves/central-ves/service/gin-helper"
	"github.com/Myriad-Dreamin/go-ves/central-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
	"strconv"
	"time"
)

func (srv *Service) Login(c controller.MContext) {
	var req = new(control.LoginRequest)

	if !ginhelper.BindRequest(c, req) {
		return
	}

	var user *model.User
	var err error
	if req.Id != 0 {
		user, err = srv.userDB.Query(req.Id)
	} else if len(req.Name) != 0 {
		user, err = srv.userDB.QueryName(req.Name)
	} else {
		c.JSON(http.StatusOK, &serial.Response{
			Code: types.CodeUserIDMissing,
		})
		return
	}
	if ginhelper.MaybeSelectError(c, user, err) {
		return
	}

	if !ginhelper.AuthenticatePassword(c, user, req.Password) {
		return
	}

	if token, refreshToken, err := srv.middleware.GenerateTokenWithRefreshToken(&types.CustomFields{UID: int64(user.ID)}); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code: types.CodeAuthGenerateTokenError,
			Err:  err.Error(),
		})
		return
	} else {
		user.LastLogin = time.Now()

		var identities []string
		for tst := range types.Groups {
			if srv.enforcer.HasGroupingPolicy("user:"+strconv.Itoa(int(user.ID)), types.Groups[tst]) {
				identities = append(identities, types.Groups[tst])
			}
		}

		c.JSON(http.StatusOK, control.SerializeLoginReply(types.CodeOK, user, identities, token, refreshToken))

		aff, err := user.UpdateFields([]string{"last_login"})
		if err != nil || aff == 0 {
			srv.logger.Debug("update last login failed", "error", err, "affected", aff)
		}

		return
	}
}
