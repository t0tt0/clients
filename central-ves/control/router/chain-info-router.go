package router

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/service"
)

type ChainInfoRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware
	IDRouter   *ChainInfoIDRouter

	Post    *LeafRouter
	GetList *LeafRouter
}

type ChainInfoIDRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	Get    *LeafRouter
	Put    *LeafRouter
	Delete *LeafRouter
}

func BuildChainInfoRouter(parent H, serviceProvider *service.Provider) (router *ChainInfoRouter) {
	chainInfoService := serviceProvider.ChainInfoService()
	router = &ChainInfoRouter{
		Router:     parent.GetRouter().Extend("chain_info"),
		AuthRouter: parent.GetAuthRouter().Extend("chain_info"),
		Auth:       parent.GetAuth().Copy(),
	}
	router.GetList = router.GET("chain_info-list", chainInfoService.ListChainInfos)
	router.Post = router.AuthRouter.POST("/chain_info", chainInfoService.PostChainInfo)

	router.IDRouter = router.IDRouter.subBuild(router, serviceProvider)

	return
}

func (*ChainInfoIDRouter) subBuild(parent *ChainInfoRouter, serviceProvider *service.Provider) (
	router *ChainInfoIDRouter) {

	chainInfoService := serviceProvider.ChainInfoService()

	router = &ChainInfoIDRouter{
		Router:     parent.Group("/chain_info/:cid"),
		AuthRouter: parent.AuthRouter.Group("/chain_info/:cid"),
		Auth:       parent.Auth.MustGroup("chain_info", "cid"),
	}

	router.Get = router.GET("", chainInfoService.GetChainInfo)
	router.Put = router.AuthRouter.PUT("", chainInfoService.PutChainInfo)
	router.Delete = router.AuthRouter.DELETE("", chainInfoService.Delete)
	return
}

func (s *Provider) ChainInfoRouter() *ChainInfoRouter {
	return s.chainInfoRouter
}
