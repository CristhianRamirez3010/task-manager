package repository

import "github.com/CristhianRamirez3010/task-manager-go/src/v1/repository/impl"

type IUsPrPersonProjectRepo interface{}

func BuildIUsPrPersonProjectRepo() IUsPrPersonProjectRepo {
	return impl.BuildUsPrPersonProjectRepoImpl()
}
