package router

import (
	"github.com/labstack/echo/v4"
	"online-mart-server-2/pkg/presenter/admin"
)

func AdminRoute(engine *echo.Echo, middlewares ...echo.MiddlewareFunc) {

	routers := engine.Group("/admin", middlewares...)

	// logout router (登出功能)
	routers.DELETE("/logout", admin.AdminLogout)
}
