package useHistoryTokensModel

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
)

const (
	TABLE_NAME    = "use_historytokens"
	ID            = "id"
	TOKEN         = "token"
	FINISH        = "finish"
	LOGIN_ID      = "login_id"
	USER_REGISTER = "user_register"
	DATE_REGISTER = "date_register"
	USER_UPDATE   = "user_update"
	DATE_UPDATE   = "date_update"

	errDefault = "Error with de model"
)

type UseHistoryToekensModel struct {
	Id           int64     `db:"id"`
	Token        string    `db:"token"`
	Finish       time.Time `db:"finish"`
	LoginId      int64     `db:"login_id"`
	UserRegister string    `db:"user_register"`
	DateRegister time.Time `db:"date_register"`
	UserUpdate   string    `db:"user_update"`
	DateUpdate   time.Time `db:"date_update"`

	LoginModel interface{}
}

func (h *UseHistoryToekensModel) ScanModel(rows *sql.Rows) *errorManagerDto.ErrorManagerDto {
	err := rows.Scan(
		&h.Id,
		&h.Token,
		&h.Finish,
		&h.LoginId,
		&h.UserRegister,
		&h.DateRegister,
		&h.UserUpdate,
		&h.DateUpdate,
	)
	if err != nil {
		return utils.Logger("Error with te transfrom in (ScanModel, useHistoryTokenModel)", errDefault, http.StatusInternalServerError, err.Error())
	}

	return nil
}
