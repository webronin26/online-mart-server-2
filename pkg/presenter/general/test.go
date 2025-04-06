package general

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-mart-server-2/pkg/presenter"
	"online-mart-server-2/pkg/usecases/test"
)

func Test(ctx echo.Context) error {

	var result presenter.Result

	output, statusCode, err := test.Exec()
	if err != nil {
		result.Code = statusCode
		result.Data = err.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = statusCode
	result.Data = output

	return ctx.JSON(http.StatusOK, result)
}
