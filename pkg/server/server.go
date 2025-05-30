package server

import (
	"github.com/labstack/echo/v4"
	"online-mart-server-2/config"
	"online-mart-server-2/pkg/server/middleware"
	"online-mart-server-2/pkg/server/router"
)

func Init() {

	e := echo.New()

	// set up middleware
	// 註冊中間件
	e.Use(
		middleware.Recover,
		middleware.CORS,
		middleware.Logger,
	)

	// set up general-route (do not need authorisation or token)
	// general (通用通道)：不用驗證即可通過
	router.GeneralRoute(e)

	// admin route (will certify admin user's token)
	// admin (管理者通道)：需要管理者驗證
	router.AdminRoute(e, middleware.CertifyAdmin)

	// retailer route (will certify retailer user's token)
	// retailer (店家通道)：需要店家驗證
	router.RetailerRoute(e, middleware.CertifyRetailer)

	// member route (will certify member user's token)
	// member (會員通道)：需要會員驗證
	router.MemberRoute(e, middleware.CertifyMember)

	e.Logger.Fatal(e.Start(":" + config.GetSystemConfig().Address))
}
