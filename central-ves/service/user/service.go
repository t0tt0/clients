//go:generate package-attach-to-path -generate_register_map
package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/go-ves/central-ves/config"
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	base_service "github.com/Myriad-Dreamin/go-ves/central-ves/lib/base-service"
	"github.com/Myriad-Dreamin/go-ves/central-ves/lib/jwt"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/go-ves/central-ves/types"
	"github.com/casbin/casbin/v2"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	userDB     *model.UserDB
	enforcer   *casbin.SyncedEnforcer
	logger     types.Logger
	cfg        *config.ServerConfig
	middleware *jwt.Middleware
}

func (srv *Service) UserServiceSignatureXXX() interface{} { return srv }

func NewService(m module.Module) (a control.UserService, err error) {
	srv := new(Service)
	provider := m.Require(config.ModulePath.Minimum.Provider.Model).(*model.Provider)
	srv.logger = m.Require(config.ModulePath.Minimum.Global.Logger).(types.Logger)
	srv.cfg = m.Require(config.ModulePath.Minimum.Global.Configuration).(*config.ServerConfig)
	srv.middleware = m.Require(config.ModulePath.Minimum.Middleware.JWT).(*jwt.Middleware)
	srv.userDB = provider.UserDB()
	srv.enforcer = provider.Enforcer()
	srv.CRUDService = base_service.NewCRUDService(srv, srv.cfg.BaseParametersConfig.PathPlaceholder.User)
	srv.ListService = base_service.NewListService(srv, srv.userDB.FilterI)
	a = srv
	return
}

/*
type User struct {
}
*/
