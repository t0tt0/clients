package userservice

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/control"
	"github.com/HyperService-Consortium/go-ves/central-ves/model"
	ginhelper "github.com/HyperService-Consortium/go-ves/lib/backend/gin-helper"
	"github.com/HyperService-Consortium/go-ves/lib/backend/serial"
	types2 "github.com/HyperService-Consortium/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/rbac"
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
	aff, err := srv.userDB.Register(user)
	if err != nil {
		//fmt.Println(reflect.TypeOf(err))
		if ginhelper.CheckInsertError(c, err) {
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code: types2.CodeInsertError,
			Err:  err.Error(),
		})
		return
	} else if aff == 0 {
		c.JSON(http.StatusOK, &serial.ErrorSerializer{
			Code: types2.CodeInsertError,
			Err:  "existed",
		})
		return
	}
	c.JSON(http.StatusOK, srv.AfterPost(control.SerializeRegisterReply(types2.CodeOK, user)))

	_, err = rbac.AddGroupingPolicy("user:"+strconv.Itoa(int(user.ID)), "norm")
	if err != nil {
		srv.logger.Debug("update group error", "error", err)
	}
}
