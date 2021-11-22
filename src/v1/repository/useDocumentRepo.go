package repository

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/userDocumentTypeModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository/impl"
)

type IUseDocumentRepo interface {
	GetAll() ([]*userDocumentTypeModel.UseDocumentTypeModel, *errorManagerDto.ErrorManagerDto)
}

func BuildIUseDocumentRepo() IUseDocumentRepo {

	return impl.BuildUseDocumentImpl()
}
