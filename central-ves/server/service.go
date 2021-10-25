package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/functional-go"
	"github.com/HyperService-Consortium/go-ves/central-ves/service"
)

type serviceResult struct {
	serviceName string
	functional.DecayResult
}

func (srv *Server) PrepareService() bool {
	for _, serviceResult := range []serviceResult{
		{"chainInfoService", functional.Decay(service.NewChainInfoService(srv.Module))},
		{"userService", functional.Decay(service.NewUserService(srv.Module))},
		{"authService", functional.Decay(service.NewAuthService(srv.Module))},
		{"objectService", functional.Decay(service.NewObjectService(srv.Module))},
	} {
		// build Router failed when requesting service with database, report and return
		if serviceResult.Err != nil {
			srv.Logger.Debug(fmt.Sprintf("get %T service error", serviceResult.First), "error", serviceResult.Err)
			return false
		}
		srv.ServiceProvider.Register(serviceResult.serviceName, serviceResult.First)
	}
	return true
}
