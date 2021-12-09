package repository

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useHistoryTokensModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository/impl"
)

type IUseHistoryTokensRepo interface {
	NewToken(hToken *useHistoryTokensModel.UseHistoryToekensModel) *errorManagerDto.ErrorManagerDto
	GetLastTokenByPersonId(personId int64) (*useHistoryTokensModel.UseHistoryToekensModel, *errorManagerDto.ErrorManagerDto)
}

func BuildIUseHistoryTokensRepo() IUseHistoryTokensRepo {
	return impl.BuildUseHistoryTokensRepoImpl()
}
