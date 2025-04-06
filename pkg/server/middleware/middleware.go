package middleware

import "github.com/labstack/echo/v4/middleware"

var (
	Recover = middleware.Recover()
	CORS    = middleware.CORS()
	Logger  = middleware.Logger()
)
