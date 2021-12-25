package repository

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/constants"
	"github.com/CristhianRamirez3010/task-manager-go/src/tools/sqlTools"
)

type IUsPrPersonProjectRepo interface{}

func BuildIUsPrPersonProjectRepo() IUsPrPersonProjectRepo {
	repo := new(usPrPersonProjectRepoImpl)
	repo.constants = new(constants.Constants)

	return repo
}

type usPrPersonProjectRepoImpl struct {
	constants *constants.Constants
	sqlTools  sqlTools.SqlTool
}
