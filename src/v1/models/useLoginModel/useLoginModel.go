package useLoginModel

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
)

const (
	TABLE_NAME    = "use_login"
	ID            = "id"
	EMAIL         = "email"
	USER          = "user"
	PASSOWRD      = "password"
	USER_REGISTER = "user_register"
	DATE_REGISTER = "date_register"
	USER_UPDATE   = "user_update"
	DATE_UPDATE   = "date_update"

	errDefault = "Error with de model"
)

type UseLoginModel struct {
	Id           int64     `db:"id"`
	Email        string    `db:"email"`
	User         string    `db:"user"`
	Password     string    `db:"password"`
	UserRegister string    `db:"user_register"`
	DateRegister time.Time `db:"date_register"`
	UserUpdate   string    `db:"user_update"`
	DateUpdate   time.Time `db:"date_update"`

	HistoryTokesModelList interface{}
}

func (l *UseLoginModel) ScanModel(rows *sql.Rows) *errorManagerDto.ErrorManagerDto {
	err := rows.Scan(
		&l.Id,
		&l.Email,
		&l.User,
		&l.Password,
		&l.UserRegister,
		&l.DateRegister,
		&l.UserUpdate,
		&l.UserUpdate,
	)
	if err != nil {
		return utils.Logger("The tranform in the model failed", errDefault, http.StatusOK, err.Error())
	}
	return nil
}
