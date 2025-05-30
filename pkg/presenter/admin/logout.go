package admin

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-mart-server-2/pkg/entities"
	"online-mart-server-2/pkg/presenter"
	"online-mart-server-2/pkg/usecases/logout"
)

func AdminLogout(ctx echo.Context) error {

	user := ctx.Get("user").(*entities.User)

	var result presenter.Result
	var input logout.Input
	input.UserId = int64(user.ID)

	if serverStatus, err := logout.Exec(input); err != nil {
		result.Code = serverStatus
		result.Data = err.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = presenter.StatusSuccess

	return ctx.JSON(http.StatusOK, result)
}
