package useHistoryTokensRepo

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

type IUseHistoryTokensRepo interface {
	NewToken(hToken *useHistoryTokensModel.UseHistoryToekensModel) *errorManagerDto.ErrorManagerDto
	GetLastTokenByPersonId(personId int64) (*useHistoryTokensModel.UseHistoryToekensModel, *errorManagerDto.ErrorManagerDto)
}

func BuildIUseHistoryTokensRepo() IUseHistoryTokensRepo {
	repo := new(useHistoryTokensRepoImpl)
	repo.constants = constants.BuildConstants()
	repo.sqlTools = sqlTools.BuildSqlTools(
		useHistoryTokensModel.TABLE_NAME,
		useHistoryTokensModel.ID,
		useHistoryTokensModel.TOKEN,
		useHistoryTokensModel.FINISH,
		useHistoryTokensModel.LOGIN_ID,
		useHistoryTokensModel.USER_REGISTER,
		useHistoryTokensModel.DATE_REGISTER,
		useHistoryTokensModel.USER_UPDATE,
		useHistoryTokensModel.DATE_UPDATE,
	)
	return repo
}

type useHistoryTokensRepoImpl struct {
	constants *constants.Constants
	sqlTools  *sqlTools.SqlTool
}

var (
	logger               = utils.Logger
	buildStrFromArray    = utils.BuildStrFromArray
	buildMySQLConnection = connections.BuildMySQLConnection
)

func (h *useHistoryTokensRepoImpl) NewToken(hToken *useHistoryTokensModel.UseHistoryToekensModel) *errorManagerDto.ErrorManagerDto {
	db, errDto := h.loadConnection()
	if errDto != nil {
		return errDto
	}
	defer db.Close()

	query := h.sqlTools.AddInsert().BuildQuery()
	res, err := db.Exec(query,
		0,
		hToken.Token,
		hToken.Finish,
		hToken.LoginId,
		hToken.UserRegister,
		hToken.DateRegister,
		hToken.UserUpdate,
		hToken.DateUpdate,
	)
	if err != nil {
		return utils.Logger("Error with the insert (NewToken,HistoryTokeRepo)", exceptions.ERR_DEFAULT_REPO, http.StatusInternalServerError, err.Error())
	}

	if hToken.Id, err = res.LastInsertId(); err != nil {
		hToken.Id = 0
		return logger("Problems when i tried to get id (NewToken,HistoryTokeRepo)", exceptions.ERR_DEFAULT_REPO, http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (h *useHistoryTokensRepoImpl) GetLastTokenByPersonId(personId int64) (*useHistoryTokensModel.UseHistoryToekensModel, *errorManagerDto.ErrorManagerDto) {
	db, errDto := h.loadConnection()
	if errDto != nil {
		return nil, errDto
	}
	defer db.Close()

	tokenRef := "token"
	personRef := "person"
	loginRef := "login"
	query := h.sqlTools.AddSelect().
		AddInnerJoin(useLoginModel.TABLE_NAME, loginRef,
			fmt.Sprintf(" %s.%s = %s.%s ", loginRef, useLoginModel.ID, tokenRef, useHistoryTokensModel.LOGIN_ID)).
		AddInnerJoin(usePersonalDataModel.TABLE_NAME, personRef,
			fmt.Sprintf(" %s.%s = %s.%s ", personRef, usePersonalDataModel.LOGIN_ID, loginRef, useHistoryTokensModel.ID)).
		AddWhere(buildStrFromArray(
			fmt.Sprintf(" %s.%s=? ", personRef, usePersonalDataModel.ID),
			fmt.Sprintf(" order by %s.%s  desc ", tokenRef, useHistoryTokensModel.ID),
		)).BuildQuery()

	rows, err := db.Query(query, personId)
	if err != nil {
		return nil, logger("Error with teh query (GetLastTokenByPersonId,historyTokeRepo)", exceptions.ERR_DEFAULT_REPO, http.StatusInternalServerError, err.Error())
	}

	var tokenModel useHistoryTokensModel.UseHistoryToekensModel
	if rows.Next() {
		errDto = tokenModel.ScanModel(rows)
		if errDto != nil {
			return nil, errDto
		}
	}
	return &tokenModel, nil
}

func (h *useHistoryTokensRepoImpl) loadConnection() (*sql.DB, *errorManagerDto.ErrorManagerDto) {
	return buildMySQLConnection(
		h.constants.GetMysqlConnectionString(),
		h.constants.GetMaxOpenDbConn(),
		h.constants.GetMasIdleDbConn(),
		h.constants.GetMaxDbLifetime(),
	).ConnectDBMysql()
}
