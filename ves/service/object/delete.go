package objectservice

import (
	"github.com/HyperService-Consortium/go-ves/lib/backend/serial"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/ves/model"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
)

func (svc *Service) deleteHook(c controller.MContext, object *model.Object) (canDelete bool) {
	c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
		Code: types.CodeDeleteError,
		Err:  "generated delete api has not been implemented yet",
	})
	return false
}
