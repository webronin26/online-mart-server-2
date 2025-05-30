package router

import (
	"github.com/labstack/echo/v4"
	"online-mart-server-2/pkg/presenter/retailer"
)

func RetailerRoute(engine *echo.Echo, middlewares ...echo.MiddlewareFunc) {

	routers := engine.Group("/retailer", middlewares...)

	// logout router (登出功能)
	routers.DELETE("/logout", retailer.RetailerLogout)
}
