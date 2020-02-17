package chainInfoservice

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (svc *Service) deleteHook(c controller.MContext, chainInfo *model.ChainInfo) (canDelete bool) {

	return true
}
