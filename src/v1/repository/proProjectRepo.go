package repository

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/proProjectModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/usePersonalDataModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository/impl"
)

type IProProjectRepo interface {
	FindProjectByUser(user *usePersonalDataModel.UsePersonalDataModel) ([]*proProjectModel.ProProjectModel, *errorManagerDto.ErrorManagerDto)
}

func BuildIProProjectRepo() IProProjectRepo {
	return impl.BuildProProjectRepoImpl()
}
