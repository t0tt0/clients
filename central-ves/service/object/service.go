//go:generate package-attach-to-path -generate_register_map
package objectservice

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/config"
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	base_service "github.com/Myriad-Dreamin/go-ves/lib/backend/base-service"
	types2 "github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	db     model.ObjectDB
	cfg    *config.ServerConfig
	logger types2.Logger
	key    string
}

func (svc *Service) ObjectServiceSignatureXXX() interface{} { return svc }

func NewService(m module.Module) (control.ObjectService, error) {
	var a = new(Service)
	provider := m.Require(config.ModulePath.Minimum.Provider.Model).(model.Provider)
	a.logger = m.Require(config.ModulePath.Minimum.Global.Logger).(types2.Logger)
	a.cfg = m.Require(config.ModulePath.Minimum.Global.Configuration).(*config.ServerConfig)
	a.key = "oid"
	a.db = provider.ObjectDB()
	a.CRUDService = base_service.NewCRUDService(a, provider.ObjectDB().GetTraits(), a.key)
	a.ListService = base_service.NewListService(a, a.db.FilterI)
	return a, nil
}
