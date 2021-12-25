package useLoginRepo

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
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
)

type IUseLoginRepo interface {
	FindByEmailAndUser(email string, user string) ([]*useLoginModel.UseLoginModel, *errorManagerDto.ErrorManagerDto)
	New(useLogin *useLoginModel.UseLoginModel) *errorManagerDto.ErrorManagerDto
}

func BuildIUseLoginRepo() IUseLoginRepo {
	repo := new(useLoginRepoImpl)
	repo.constants = new(constants.Constants)
	repo.sqlTools = sqlTools.BuildSqlTools(
		useLoginModel.TABLE_NAME,
		useLoginModel.ID,
		useLoginModel.EMAIL,
		useLoginModel.USER,
		useLoginModel.PASSOWRD,
		useLoginModel.USER_REGISTER,
		useLoginModel.DATE_REGISTER,
		useLoginModel.USER_UPDATE,
		useLoginModel.DATE_UPDATE,
	)
	return repo
}

type useLoginRepoImpl struct {
	constants *constants.Constants
	sqlTools  *sqlTools.SqlTool
}

var (
	buildMySQLConnection = connections.BuildMySQLConnection
	logger               = utils.Logger
)

func (u *useLoginRepoImpl) FindByEmailAndUser(email string, user string) ([]*useLoginModel.UseLoginModel, *errorManagerDto.ErrorManagerDto) {
	var usersLogin []*useLoginModel.UseLoginModel
	db, errDto := u.loadConnection()
	if errDto != nil {
		return nil, errDto
	}
	defer db.Close()

	query := u.sqlTools.AddSelect().AddWhere(
		fmt.Sprintf("%s=? or %s=?", useLoginModel.EMAIL, useLoginModel.USER),
	).BuildQuery()

	rows, err := db.Query(query, email, user)
	if err != nil {
		return nil, logger("Error with the query in UseLoginRepo", exceptions.ERR_DEFAULT_REPO, http.StatusInternalServerError, err.Error())
	}

	for rows.Next() {
		userLogin := useLoginModel.UseLoginModel{}
		userLogin.ScanModel(rows)
		usersLogin = append(usersLogin, &userLogin)
	}

	return usersLogin, nil
}

func (u *useLoginRepoImpl) New(useLogin *useLoginModel.UseLoginModel) *errorManagerDto.ErrorManagerDto {
	db, errDto := u.loadConnection()
	if errDto != nil {
		return errDto
	}
	defer db.Close()

	query := u.sqlTools.AddInsert().BuildQuery()
	res, err := db.Exec(query,
		useLogin.Id,
		useLogin.Email,
		useLogin.User,
		useLogin.Password,
		useLogin.UserRegister,
		useLogin.DateRegister,
		useLogin.UserUpdate,
		useLogin.DateUpdate,
	)
	if err != nil {
		return logger("Error with the insert in LoginRepo(new())", exceptions.ERR_DEFAULT_REPO, http.StatusInternalServerError, err.Error())
	}

	if useLogin.Id, err = res.LastInsertId(); err != nil {
		useLogin.Id = 0
		return logger("Problems when i tried to get id", exceptions.ERR_DEFAULT_REPO, http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (h *useLoginRepoImpl) loadConnection() (*sql.DB, *errorManagerDto.ErrorManagerDto) {
	return buildMySQLConnection(
		h.constants.GetMysqlConnectionString(),
		h.constants.GetMaxOpenDbConn(),
		h.constants.GetMasIdleDbConn(),
		h.constants.GetMaxDbLifetime(),
	).ConnectDBMysql()
}
