package impl

import (
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository"
)

type ProjectHandlerImpl struct {
	contextDto *contextDto.ContextDto

	handlerBase
}

func BuildProjectHandlerImpl(con *contextDto.ContextDto) *ProjectHandlerImpl {
	return &ProjectHandlerImpl{
		contextDto: con,
	}
}

func (p *ProjectHandlerImpl) GetProject() *responseDto.ResponseDto {
	personalDataModel, errDto := p.validateToken(p.contextDto.AccessToken)
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

func (p *ProjectHandlerImpl) NewProject() *responseDto.ResponseDto {
	if p.contextDto.AccessToken == "" {
		return &responseDto.ResponseDto{
			Error: *utils.Logger(errTokenInvalid, errTokenInvalid, http.StatusPreconditionFailed, ""),
		}
	}
	return &responseDto.ResponseDto{}
}
