//go:generate package-attach-to-path -generate_register_map
package objectservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/go-ves/vesx/config"
	"github.com/Myriad-Dreamin/go-ves/vesx/control"
	base_service "github.com/Myriad-Dreamin/go-ves/vesx/lib/base-service"
	"github.com/Myriad-Dreamin/go-ves/vesx/model"
	"github.com/Myriad-Dreamin/go-ves/vesx/types"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	db     *model.ObjectDB
	cfg    *config.ServerConfig
	logger types.Logger
	key    string
}

func (svc *Service) ObjectServiceSignatureXXX() interface{} { return svc }

func NewService(m module.Module) (control.ObjectService, error) {
	var a = new(Service)
	provider := m.Require(config.ModulePath.Provider.Model).(*model.Provider)
	a.logger = m.Require(config.ModulePath.Global.Logger).(types.Logger)
	a.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)
	a.key = "oid"
	a.db = provider.ObjectDB()
	a.CRUDService = base_service.NewCRUDService(a, a.key)
	a.ListService = base_service.NewListService(a, a.db.FilterI)
	return a, nil
}
