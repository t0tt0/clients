//go:generate package-attach-to-path -generate_register_map
package chainInfoservice

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/config"
	"github.com/HyperService-Consortium/go-ves/central-ves/control"
	"github.com/HyperService-Consortium/go-ves/central-ves/model"
	base_service "github.com/HyperService-Consortium/go-ves/lib/backend/base-service"
	types2 "github.com/HyperService-Consortium/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	db     model.ChainInfoDB
	cfg    *config.ServerConfig
	logger types2.Logger
	key    string
}

func (svc *Service) ChainInfoServiceSignatureXXX() interface{} { return svc }

func NewService(m module.Module) (control.ChainInfoService, error) {
	var a = new(Service)
	provider := m.Require(config.ModulePath.Minimum.Provider.Model).(model.Provider)
	a.logger = m.Require(config.ModulePath.Minimum.Global.Logger).(types2.Logger)
	a.cfg = m.Require(config.ModulePath.Minimum.Global.Configuration).(*config.ServerConfig)
	a.key = "cid"
	a.db = provider.ChainInfoDB()
	a.CRUDService = base_service.NewCRUDService(a, provider.ChainInfoDB().GetTraits(), a.key)
	a.ListService = base_service.NewListService(a, a.db.FilterI)
	return a, nil
}
