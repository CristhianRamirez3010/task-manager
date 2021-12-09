package usprPersonProjectModel

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
)

const (
	TABLE_NAME      = "uspr_personproject"
	ID              = "id"
	PEROSNALDATA_ID = "personaldata_id"
	PROJECT_ID      = "project_id"
	MASTER          = "master"
	USER_REGISTER   = "user_register"
	DATE_REGISTER   = "date_register"
	USER_UPDATE     = "user_update"
	DATE_UPDATE     = "date_update"

	errDefault = "Error with the model"
)

type UsPrPersonProject struct {
	Id             int64     `db:"id"`
	PersonalDataId int64     `db:"personaldata_id"`
	Project_id     int64     `db:"project_id"`
	Master         bool      `db:"master"`
	UserRegister   string    `db:"user_register"`
	DateRegister   time.Time `db:"date_register"`
	UserUpdate     string    `db:"user_update"`
	DateUpdate     time.Time `db:"date_update"`
}

func (p *UsPrPersonProject) ScanModel(rows *sql.Rows) *errorManagerDto.ErrorManagerDto {
	err := rows.Scan(
		&p.Id,
		&p.PersonalDataId,
		&p.Project_id,
		&p.Master,
		&p.UserRegister,
		&p.DateRegister,
		&p.UserUpdate,
		&p.DateUpdate,
	)
	if err != nil {
		return utils.Logger("Error with the scan model (UPrPersonProject,ScanModel)", errDefault, http.StatusInternalServerError, err.Error())
	}
	return nil
}
