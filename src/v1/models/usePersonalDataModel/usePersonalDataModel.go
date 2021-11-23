package usePersonalDataModel

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
)

const (
	TABLE_NAME     = "use_personaldata"
	ID             = "id"
	NAME           = "name"
	SURNAME        = "surname"
	IDENTIFICATION = "identification"
	PHONE          = "phone"
	COUNTRY        = "country"
	LOGIN_ID       = "login_id"
	USER_REGISTER  = "user_register"
	DATE_REGISTER  = "date_register"
	USER_UPDATE    = "user_update"
	DATE_UPDATE    = "date_update"

	errDefault = "Error with de model"
)

type UsePersonalDataModel struct {
	Id             int64     `db:"id"`
	Name           string    `db:"name"`
	Surname        string    `db:"surname"`
	Identification string    `db:"identification"`
	Phone          string    `db:"phone"`
	Country        string    `db:"country"`
	LoginId        int64     `db:"login_id"`
	UserRegister   string    `db:"user_register"`
	DateRegister   time.Time `db:"date_register"`
	UserUpdate     string    `db:"user_update"`
	DateUpdate     time.Time `db:"date_update"`
}

func (p *UsePersonalDataModel) ScanModel(rows *sql.Rows) *errorManagerDto.ErrorManagerDto {
	err := rows.Scan(
		&p.Id,
		&p.Name,
		&p.Surname,
		&p.Identification,
		&p.Phone,
		&p.Country,
		&p.LoginId,
		&p.UserRegister,
		&p.DateRegister,
		&p.UserUpdate,
		&p.UserUpdate,
	)
	if err != nil {
		return utils.Logger("The tranform in the model failed", errDefault, http.StatusOK, err.Error())
	}
	return nil
}
