package proProjectModel

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
)

const (
	TABLE_NAME    = "pro_projects"
	ID            = "id"
	NAME          = "name"
	DESCRIPTION   = "description"
	USER_REGISTER = "user_register"
	DATE_REGISTER = "date_register"
	USER_UPDATE   = "user_update"
	DATE_UPDATE   = "date_update"

	errDefault = "Error with the model"
)

type ProProjectModel struct {
	Id           int64     `db:"id"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	UserRegister string    `db:"user_register"`
	DateRegister time.Time `db:"date_register"`
	UserUpdate   string    `db:"user_update"`
	DateUpdate   time.Time `db:"date_update"`
}

func (p *ProProjectModel) ScanModel(rows *sql.Rows) *errorManagerDto.ErrorManagerDto {
	err := rows.Scan(
		&p.Id,
		&p.Name,
		&p.Description,
		&p.UserRegister,
		&p.DateRegister,
		&p.UserUpdate,
		&p.DateUpdate,
	)

	if err != nil {
		return utils.Logger("Error with the scan (ProProjectModel, Scan)", errDefault, http.StatusInternalServerError, err.Error())
	}
	return nil
}
