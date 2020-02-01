//go:generate package-attach-to-path -generate_register_map
package chainInfoservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/go-ves/central-ves/config"
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	base_service "github.com/Myriad-Dreamin/go-ves/central-ves/lib/base-service"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/go-ves/central-ves/types"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	db     *model.ChainInfoDB
	cfg    *config.ServerConfig
	logger types.Logger
	key    string
}

func (svc *Service) ChainInfoServiceSignatureXXX() interface{} { return svc }

func NewService(m module.Module) (control.ChainInfoService, error) {
	var a = new(Service)
	provider := m.Require(config.ModulePath.Provider.Model).(*model.Provider)
	a.logger = m.Require(config.ModulePath.Global.Logger).(types.Logger)
	a.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)
	a.key = "cid"
	a.db = provider.ChainInfoDB()
	a.CRUDService = base_service.NewCRUDService(a, a.key)
	a.ListService = base_service.NewListService(a, a.db.FilterI)
	return a, nil
}
