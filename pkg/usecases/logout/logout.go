package logout

import (
	"errors"
	"github.com/jinzhu/gorm"
	"online-mart-server-2/pkg/entities"
	"online-mart-server-2/pkg/presenter"
	"online-mart-server-2/pkg/store"
)

type Input struct {
	UserId int64
}

func Exec(input Input) (presenter.StatusCode, error) {

	update := store.DB.Model(entities.User{}).
		Where("id = ?", input.UserId).
		Update("token", gorm.Expr("NULL"))
	if err := update.Error; err != nil {
		return presenter.StatusSQLError, errors.New("update error, logout error")
	}

	return presenter.StatusSuccess, nil
}
