package router

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/service"
)

type UserRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware
	IDRouter   *UserIDRouter

	Login    *LeafRouter
	Register *LeafRouter
	GetList  *LeafRouter
}

type UserIDRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	ChangePassword *LeafRouter
	Get            *LeafRouter
	Put            *LeafRouter
	Delete         *LeafRouter
}

func BuildUserRouter(parent *RootRouter, serviceProvider *service.Provider) (router *UserRouter) {
	userService := serviceProvider.UserService()
	router = &UserRouter{
		Router:     parent.GetRouter().Extend("user"),
		AuthRouter: parent.GetAuthRouter().Extend("user"),
		Auth:       parent.GetAuth().Copy(),
	}
	router.GetList = router.GET("user-list", userService.ListUsers)
	router.Register = router.POST("/user", userService.Register)
	router.Login = router.POST("/login", userService.Login)

	router.IDRouter = router.IDRouter.subBuild(router, serviceProvider)

	return
}

func (*UserIDRouter) subBuild(parent *UserRouter, serviceProvider *service.Provider) (
	router *UserIDRouter) {

	userService := serviceProvider.UserService()

	router = &UserIDRouter{
		Router:     parent.Group("/user/:id"),
		AuthRouter: parent.AuthRouter.Group("/user/:id"),
		Auth:       parent.Auth.MustGroup("user", "id"),
	}

	router.Get = router.GET("", userService.GetUser)
	router.ChangePassword = router.AuthRouter.PUT("/password", userService.ChangePassword)
	router.Put = router.AuthRouter.PUT("", userService.PutUser)
	router.Delete = router.AuthRouter.DELETE("", userService.Delete)

	return
}

func (s *Provider) UserRouter() *UserRouter {
	return s.userRouter
}
