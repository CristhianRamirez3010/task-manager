package usePersonalDataRepo

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/connections"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/constants"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/exceptions"
	"github.com/CristhianRamirez3010/task-manager-go/src/tools/sqlTools"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useHistoryTokensModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/usePersonalDataModel"
)

type IUsePersonalDataRepo interface {
	GetDataByToken(token string) (*usePersonalDataModel.UsePersonalDataModel, *errorManagerDto.ErrorManagerDto)
	GetDataByLoginId(loginId int64) (*usePersonalDataModel.UsePersonalDataModel, *errorManagerDto.ErrorManagerDto)
	New(personaldata *usePersonalDataModel.UsePersonalDataModel) *errorManagerDto.ErrorManagerDto
}

func BuildIUsePersonalDataRepo() IUsePersonalDataRepo {
	repo := new(usePersonalDataRepoImpl)
	repo.constants = new(constants.Constants)
	repo.sqlTools = sqlTools.BuildSqlTools(
		usePersonalDataModel.TABLE_NAME,
		usePersonalDataModel.ID,
		usePersonalDataModel.NAME,
		usePersonalDataModel.SURNAME,
		usePersonalDataModel.DOCUMENT,
		usePersonalDataModel.PHONE,
		usePersonalDataModel.COUNTRY,
		usePersonalDataModel.DOCUMENTTYPE,
		usePersonalDataModel.LOGIN_ID,
		usePersonalDataModel.USER_REGISTER,
		usePersonalDataModel.DATE_REGISTER,
		usePersonalDataModel.USER_UPDATE,
		usePersonalDataModel.DATE_UPDATE,
	)
	return repo
}

type usePersonalDataRepoImpl struct {
	constants *constants.Constants
	sqlTools  *sqlTools.SqlTool
}

var (
	buildMySQLConnection = connections.BuildMySQLConnection
	logger               = utils.Logger
)

func (p *usePersonalDataRepoImpl) GetDataByLoginId(loginId int64) (*usePersonalDataModel.UsePersonalDataModel, *errorManagerDto.ErrorManagerDto) {

	db, errDto := p.loadConnection()
	if errDto != nil {
		return nil, errDto
	}
	defer db.Close()

	query := p.sqlTools.AddSelect().AddWhere(
		fmt.Sprintf("%s=?", usePersonalDataModel.LOGIN_ID),
	).BuildQuery()
	rows, err := db.Query(query, loginId)
	if err != nil {
		return nil, utils.Logger("Error with the query", exceptions.ERR_DEFAULT_REPO, http.StatusInternalServerError, err.Error())
	}
	usePersonalData := new(usePersonalDataModel.UsePersonalDataModel)
	for rows.Next() {
		usePersonalData.ScanModel(rows)
	}
	return usePersonalData, nil
}

func (p *usePersonalDataRepoImpl) GetDataByToken(token string) (*usePersonalDataModel.UsePersonalDataModel, *errorManagerDto.ErrorManagerDto) {
	db, errDto := p.loadConnection()
	if errDto != nil {
		return nil, errDto
	}
	defer db.Close()

	personRef := "person"
	loginRef := "login"
	tokenRef := "histTok"

	query := p.sqlTools.AddSelectWithRef(personRef).
		AddInnerJoin(useLoginModel.TABLE_NAME, loginRef,
			fmt.Sprintf(" %s.%s=%s.%s ", loginRef, useLoginModel.ID, personRef, usePersonalDataModel.LOGIN_ID)).
		AddInnerJoin(useHistoryTokensModel.TABLE_NAME, tokenRef,
			fmt.Sprintf(" %s.%s=%s.%s ", tokenRef, useHistoryTokensModel.LOGIN_ID, loginRef, useLoginModel.ID)).
		AddWhere(fmt.Sprintf(" %s.%s=? ", tokenRef, useHistoryTokensModel.TOKEN)).
		BuildQuery()

	rows, err := db.Query(query, token)
	if err != nil {
		return nil, utils.Logger("Error with the query (GetDataByToken,PersonalDataRepo)", exceptions.ERR_DEFAULT_REPO, http.StatusInternalServerError, err.Error())
	}
	personalData := new(usePersonalDataModel.UsePersonalDataModel)
	if rows.Next() {
		errDto = personalData.ScanModel(rows)
		if errDto != nil {
			return nil, errDto
		}
	}

	return personalData, nil
}

func (p *usePersonalDataRepoImpl) New(personaldata *usePersonalDataModel.UsePersonalDataModel) *errorManagerDto.ErrorManagerDto {
	db, errDto := p.loadConnection()
	if errDto != nil {
		return errDto
	}
	defer db.Close()

	query := p.sqlTools.AddInsert().BuildQuery()

	res, err := db.Exec(query,
		personaldata.Id,
		personaldata.Name,
		personaldata.Surname,
		personaldata.Document,
		personaldata.Phone,
		personaldata.Country,
		personaldata.DocumentType,
		personaldata.LoginId,
		personaldata.UserRegister,
		personaldata.DateRegister,
		personaldata.UserUpdate,
		personaldata.DateUpdate,
	)
	if err != nil {
		return logger("Error with the insert in PersonalDataRepo (New())", exceptions.ERR_DEFAULT_REPO, http.StatusInternalServerError, err.Error())
	}
	if personaldata.Id, err = res.LastInsertId(); err != nil {
		personaldata.Id = 0
		return logger("Problems when i tried to get id", exceptions.ERR_DEFAULT_REPO, http.StatusInternalServerError, errDto.Message)
	}
	return nil
}

func (h *usePersonalDataRepoImpl) loadConnection() (*sql.DB, *errorManagerDto.ErrorManagerDto) {
	return buildMySQLConnection(
		h.constants.GetMysqlConnectionString(),
		h.constants.GetMaxOpenDbConn(),
		h.constants.GetMasIdleDbConn(),
		h.constants.GetMaxDbLifetime(),
	).ConnectDBMysql()
}
