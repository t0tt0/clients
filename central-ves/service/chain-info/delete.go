package chainInfoservice

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/lib/serial"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/go-ves/central-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
)

func (svc *Service) deleteHook(c controller.MContext, chainInfo *model.ChainInfo) (canDelete bool) {
	c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
		Code: types.CodeDeleteError,
		Err:  "generated delete api has not been implemented yet",
	})
	return false
}
