package proProjectRepo

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
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/proProjectModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/usePersonalDataModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/usprPersonProjectModel"
)

type IProProjectRepo interface {
	FindProjectByUserId(userId int64) ([]*proProjectModel.ProProjectModel, *errorManagerDto.ErrorManagerDto)
	New(projectModel *proProjectModel.ProProjectModel) *errorManagerDto.ErrorManagerDto
}

func BuildIProProjectRepo() IProProjectRepo {
	repo := new(proProjectRepoImpl)
	repo.sqlTool = sqlTools.BuildSqlTools(
		proProjectModel.TABLE_NAME,
		proProjectModel.ID,
		proProjectModel.NAME,
		proProjectModel.DESCRIPTION,
		proProjectModel.USER_REGISTER,
		proProjectModel.DATE_REGISTER,
		proProjectModel.USER_UPDATE,
		proProjectModel.DATE_UPDATE,
	)

	return repo
}

type proProjectRepoImpl struct {
	constants constants.Constants
	sqlTool   *sqlTools.SqlTool
}

var (
	buildStrFromArray = utils.BuildStrFromArray
	logger            = utils.Logger
)

func (p *proProjectRepoImpl) FindProjectByUserId(userId int64) ([]*proProjectModel.ProProjectModel, *errorManagerDto.ErrorManagerDto) {
	db, errDto := p.loadConnection()
	if errDto != nil {
		return nil, errDto
	}
	defer db.Close()
	projectRef := "proj"
	persProjRef := "uprpepr"
	personRef := "person"
	query := p.sqlTool.AddSelectWithRef(projectRef).
		AddInnerJoin(usprPersonProjectModel.TABLE_NAME, persProjRef,
			fmt.Sprintf("%s.%s = %s.%s", projectRef, proProjectModel.ID, persProjRef, usprPersonProjectModel.PROJECT_ID)).
		AddInnerJoin(usePersonalDataModel.TABLE_NAME, personRef,
			fmt.Sprintf("%s.%s=%s.%s", personRef, usePersonalDataModel.ID, persProjRef, usprPersonProjectModel.PEROSNALDATA_ID)).
		AddWhere(buildStrFromArray(fmt.Sprintf("%s.%s = ?", personRef, usePersonalDataModel.ID))).
		BuildQuery()
	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, logger("Error with the query (FindProjectByUser,ProjectRepo)", exceptions.ERR_DEFAULT_REPO, http.StatusInternalServerError, err.Error())
	}

	var projects []*proProjectModel.ProProjectModel

	for rows.Next() {
		var project proProjectModel.ProProjectModel
		project.ScanModel(rows)
		projects = append(projects, &project)
	}

	return projects, nil
}

func (p *proProjectRepoImpl) New(projectModel *proProjectModel.ProProjectModel) *errorManagerDto.ErrorManagerDto {
	db, errDto := p.loadConnection()
	if errDto != nil {
		return errDto
	}
	defer db.Close()

	query := p.sqlTool.AddInsert().BuildQuery()
	res, err := db.Exec(query,
		projectModel.Id,
		projectModel.Name,
		projectModel.Description,
		projectModel.UserRegister,
		projectModel.DateRegister,
		projectModel.UserUpdate,
		projectModel.DateUpdate,
	)
	if err != nil {
		return logger("Problems with the insert (New, ProjectRepo)", exceptions.ERR_DEFAULT_REPO, http.StatusInternalServerError, err.Error())
	}

	if projectModel.Id, err = res.LastInsertId(); err != nil {
		projectModel.Id = 0
		return logger("Problems when i tried to get id", exceptions.ERR_DEFAULT_REPO, http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (p *proProjectRepoImpl) loadConnection() (*sql.DB, *errorManagerDto.ErrorManagerDto) {
	return connections.BuildMySQLConnection(
		p.constants.GetMysqlConnectionString(),
		p.constants.GetMaxOpenDbConn(),
		p.constants.GetMasIdleDbConn(),
		p.constants.GetMaxDbLifetime(),
	).ConnectDBMysql()
}
