package repository

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository/impl"
)

type IUseLoginRepo interface {
	FindByEmailAndUser(email *string, user *string) ([]*useLoginModel.UseLoginModel, *errorManagerDto.ErrorManagerDto)
	New(useLogin *useLoginModel.UseLoginModel) *errorManagerDto.ErrorManagerDto
}

func BuildIUseLoginRepo() IUseLoginRepo {
	return impl.BuildUseLoginRepoImpl()
}
