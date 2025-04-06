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

	e.Logger.Fatal(e.Start(":" + config.GetSystemConfig().Address))
}
