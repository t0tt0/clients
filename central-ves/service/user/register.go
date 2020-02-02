package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/rbac"
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	"github.com/Myriad-Dreamin/go-ves/central-ves/lib/serial"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	ginhelper "github.com/Myriad-Dreamin/go-ves/central-ves/service/gin-helper"
	"github.com/Myriad-Dreamin/go-ves/central-ves/types"
	"net/http"
	"strconv"
)

func (srv *Service) Register(c controller.MContext) {
	var req = new(control.RegisterRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	//if sug := CheckStrongPassword(req.Password); len(sug) != 0 {
	//	c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
	//		Code:  types.CodeWeakPassword,
	//		Error: sug,
	//	})
	//	return
	//}

	var user = new(model.User)
	user.Name = req.Name
	user.Password = req.Password

	// check default value
	aff, err := user.Register()
	if err != nil {
		//fmt.Println(reflect.TypeOf(err))
		if ginhelper.CheckInsertError(c, err) {
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code: types.CodeInsertError,
			Err:  err.Error(),
		})
		return
	} else if aff == 0 {
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code: types.CodeInsertError,
			Err:  "existed",
		})
		return
	}
	c.JSON(http.StatusOK, srv.AfterPost(control.SerializeRegisterReply(types.CodeOK, user)))

	_, err = rbac.AddGroupingPolicy("user:"+strconv.Itoa(int(user.ID)), "norm")
	if err != nil {
		srv.logger.Debug("update group error", "error", err)
	}
}