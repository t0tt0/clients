package ginhelper

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/go-ves/lib/serial"
	types2 "github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
)

func ResetPassword(c controller.MContext, obj *model.User, password string) bool {
	_, err := obj.ResetPassword(password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code: types2.CodeUpdateError,
			Err:  err.Error(),
		})
		return false
	}
	return true
}

func AuthenticatePassword(c controller.MContext, user *model.User, password string) bool {
	if ok, err := user.AuthenticatePassword(password); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code: types2.CodeAuthenticatePasswordError,
			Err:  err.Error(),
		})
		return false
	} else if !ok {
		c.JSON(http.StatusOK, &serial.Response{
			Code: types2.CodeUserWrongPassword,
		})
		return false
	}
	return true
}
