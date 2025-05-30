package router

import (
	"github.com/labstack/echo/v4"
	"online-mart-server-2/pkg/presenter/member"
)

func MemberRoute(engine *echo.Echo, middlewares ...echo.MiddlewareFunc) {

	routers := engine.Group("/member", middlewares...)

	// logout router (登出功能)
	routers.DELETE("/logout", member.MemberLogout)
}
