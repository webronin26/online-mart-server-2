package router

import (
	"github.com/labstack/echo/v4"
	"online-mart-server-2/pkg/presenter/general"
)

func GeneralRoute(router *echo.Echo) {

	// test API (測試用 API)
	router.GET("/test", general.Test)

	// login API for varies user type (給不一樣身份的用戶登入用的 API)
	router.POST("/login/member", general.MemberLogin)
	router.POST("/login/admin", general.AdminLogin)
	router.POST("/login/retailer", general.RetailerLogin)
}
