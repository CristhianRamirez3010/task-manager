package repository

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/usePersonalDataModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository/impl"
)

type IUsePersonalDataRepo interface {
	GetDataByLoginId(loginId *int64) (*usePersonalDataModel.UsePersonalDataModel, *errorManagerDto.ErrorManagerDto)
	New(personaldata *usePersonalDataModel.UsePersonalDataModel) *errorManagerDto.ErrorManagerDto
}

func BuildIUsePersonalDataRepo() IUsePersonalDataRepo {
	return impl.BuildUsePersonalDataRepoImpl()
}
