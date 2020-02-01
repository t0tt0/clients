package chainInfoservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
)

func (svc *Service) fillPutFields(
	c controller.MContext, chainInfo *model.ChainInfo,
	req *control.PutChainInfoRequest) (fields []string) {
	return nil
}
