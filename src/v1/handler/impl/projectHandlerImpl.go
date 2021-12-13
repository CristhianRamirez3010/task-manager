package impl

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/proProjectModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository"
)

type ProjectHandlerImpl struct {
	handlerBase
}

func BuildProjectHandlerImpl(con *contextDto.ContextDto) *ProjectHandlerImpl {
	return &ProjectHandlerImpl{
		handlerBase: handlerBase{contextDto: con},
	}
}

func (p *ProjectHandlerImpl) GetProject() *responseDto.ResponseDto {
	personalDataModel, errDto := p.validateToken()
	if errDto != nil {
		return &responseDto.ResponseDto{
			Error: *errDto,
		}
	}
	projectRepo := repository.BuildIProProjectRepo()
	projectModelList, errDto := projectRepo.FindProjectByUser(personalDataModel)
	if errDto != nil {
		return &responseDto.ResponseDto{
			Error: *errDto,
		}
	}
	return &responseDto.ResponseDto{
		Content: projectModelList,
	}
}

func (p *ProjectHandlerImpl) NewProject(projectModel *proProjectModel.ProProjectModel) *responseDto.ResponseDto {
	_, errDto := p.validateToken()
	if errDto != nil {
		return &responseDto.ResponseDto{
			Error: *errDto,
		}
	}
	projectRepo := repository.BuildIProProjectRepo()

	errDto = projectRepo.New(projectModel)
	if errDto != nil {
		return &responseDto.ResponseDto{
			Error: *errDto,
		}
	}

	return &responseDto.ResponseDto{
		Content: projectModel,
	}
}
