package ginhelper

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/lib/serial"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/go-ves/central-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
)

func ResetPassword(c controller.MContext, obj *model.User, password string) bool {
	_, err := obj.ResetPassword(password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code: types.CodeUpdateError,
			Err:  err.Error(),
		})
		return false
	}
	return true
}

func AuthenticatePassword(c controller.MContext, user *model.User, password string) bool {
	if ok, err := user.AuthenticatePassword(password); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code: types.CodeAuthenticatePasswordError,
			Err:  err.Error(),
		})
		return false
	} else if !ok {
		c.JSON(http.StatusOK, &serial.Response{
			Code: types.CodeUserWrongPassword,
		})
		return false
	}
	return true
}
