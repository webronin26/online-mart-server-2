package general

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-mart-server-2/pkg/entities"
	"online-mart-server-2/pkg/presenter"
	"online-mart-server-2/pkg/usecases/login"
)

type AdminLoginParam struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func AdminLogin(ctx echo.Context) error {

	var result presenter.Result

	var param AdminLoginParam
	if err := ctx.Bind(&param); err != nil {
		result.Code = presenter.StatusParamError
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if param.Account == "" || param.Password == "" {
		result.Code = presenter.StatusParamError
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input login.Input
	input.Account = param.Account
	input.Password = param.Password
	input.UserKind = entities.AdminUser

	output, serverStatus, loginErr := login.Exec(input)
	if loginErr != nil {
		result.Code = serverStatus
		result.Data = loginErr.Error()
		return ctx.JSON(http.StatusNotFound, result)
	}

	result.Data = output
	result.Code = serverStatus

	return ctx.JSON(http.StatusOK, result)
}
