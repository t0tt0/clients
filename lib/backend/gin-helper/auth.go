package ginhelper

import (
	"github.com/HyperService-Consortium/go-ves/lib/backend/serial"
	types2 "github.com/HyperService-Consortium/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
)

func ResetPassword(c controller.MContext, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code: types2.CodeUpdateError,
			Err:  err.Error(),
		})
		return false
	}
	return true
}

func AuthenticatePassword(c controller.MContext, ok bool, err error) bool {
	if err != nil {
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
