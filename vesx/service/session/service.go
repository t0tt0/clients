//go:generate package-attach-to-path -generate_register_map
package objectservice

import (
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/vesx/config"
	"github.com/Myriad-Dreamin/go-ves/vesx/control"
	"github.com/Myriad-Dreamin/go-ves/vesx/model"
	"github.com/Myriad-Dreamin/go-ves/vesx/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Service struct {
	db     *model.SessionDB
	signer uiptypes.Signer
	cfg    *config.ServerConfig
	logger types.Logger
	key    string
}

func (svc *Service) SessionServiceSignatureXXX() interface{} { return svc }

func NewService(m module.Module) (control.SessionService, error) {
	var a = new(Service)
	provider := m.Require(config.ModulePath.Provider.Model).(*model.Provider)
	a.signer = m.Require(config.ModulePath.Global.Signer).(uiptypes.Signer)
	a.logger = m.Require(config.ModulePath.Global.Logger).(types.Logger)
	a.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)
	a.key = "sid"
	a.db = provider.SessionDB()
	return a, nil
}
