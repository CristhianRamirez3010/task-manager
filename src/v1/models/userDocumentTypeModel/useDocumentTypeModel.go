package userDocumentTypeModel

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
)

const (
	TABLE_NAME    = "use_documenttype"
	ID            = "id"
	NAME          = "name"
	DESCRIPTION   = "description"
	USER_REGISTER = "user_register"
	DATE_REGISTER = "date_register"
	USER_UPDATE   = "user_update"
	DATE_UPDATE   = "date_update"

	errDefault = "Error with the model"
)

type UseDocumentTypeModel struct {
	Id           int64     `db:"id"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	UserRegister string    `db:"user_register"`
	DateRegister time.Time `db:"date_register"`
	UserUpdate   string    `db:"user_update"`
	DateUpdate   time.Time `db:"date_update"`
}

func (d *UseDocumentTypeModel) ScanModel(rows *sql.Rows) *errorManagerDto.ErrorManagerDto {
	err := rows.Scan(
		&d.Id,
		&d.Name,
		&d.Description,
		&d.UserRegister,
		&d.DateRegister,
		&d.UserUpdate,
		&d.UserUpdate,
	)
	if err != nil {
		return utils.Logger("The tranform in the model failed", errDefault, http.StatusOK, err.Error())
	}
	return nil
}