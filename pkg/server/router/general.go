package router

import (
	"github.com/labstack/echo/v4"
	"online-mart-server-2/pkg/presenter/general"
)

func GeneralRoute(router *echo.Echo) {
	router.GET("/test", general.Test)
}
