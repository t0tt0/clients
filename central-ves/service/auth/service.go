//go:generate package-attach-to-path -generate_register_map
package authservice

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/config"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/go-ves/lib/jwt"
	types2 "github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"net/http"
)

type Service struct {
	cfg        *config.ServerConfig
	logger     types2.Logger
	middleware *jwt.Middleware
	enforcer   *model.Enforcer
}

func (svc *Service) AuthSignatureXXX() interface{} { return svc }

type RefreshTokenReply struct {
	Code  types2.CodeRawType `json:"code"`
	Token string             `json:"token"`
}

func (svc *Service) RefreshToken(c controller.MContext) {
	newToken, err := svc.middleware.RefreshToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, RefreshTokenReply{
		Code:  types2.CodeOK,
		Token: newToken,
	})
}

func NewService(m module.Module) (a *Service, err error) {
	a = new(Service)
	a.logger = m.Require(config.ModulePath.Minimum.Global.Logger).(types2.Logger)
	a.cfg = m.Require(config.ModulePath.Minimum.Global.Configuration).(*config.ServerConfig)
	a.enforcer = m.Require(config.ModulePath.Minimum.Provider.Model).(model.Provider).Enforcer()
	a.middleware = m.Require(config.ModulePath.Minimum.Middleware.JWT).(*jwt.Middleware)
	return
}
