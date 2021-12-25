package projectHandler

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/proProjectModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/services/projectService"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/services/userService"
)

type IProjectHandler interface {
	GetProject() *responseDto.ResponseDto
	NewProject(projectModel *proProjectModel.ProProjectModel) *responseDto.ResponseDto
}

func BuildIProjectHandler(context *contextDto.ContextDto) IProjectHandler {
	handler := new(projectHandlerImpl)
	handler.contextDto = context
	return handler
}

type projectHandlerImpl struct {
	contextDto *contextDto.ContextDto

	userService    userService.IUserService
	projectService projectService.IProjectService
}

var (
	responseErrorManager = utils.ResponseErrorManager
)

func (p *projectHandlerImpl) GetProject() *responseDto.ResponseDto {
	if responseErr := p.validateUser(); responseErr != nil {
		return responseErr
	}

	p.loadProjectService()
	projectModelList, errDto := p.projectService.GetProjectByContext()
	if errDto != nil {
		return &responseDto.ResponseDto{
			Error: *errDto,
		}
	}
	return &responseDto.ResponseDto{
		Content: projectModelList,
	}
}

func (p *projectHandlerImpl) NewProject(projectModel *proProjectModel.ProProjectModel) *responseDto.ResponseDto {
	if responseErr := p.validateUser(); responseErr != nil {
		return responseErr
	}

	p.loadProjectService()
	if errDto := p.projectService.CreateNewProject(projectModel); errDto != nil {
		return &responseDto.ResponseDto{
			Error: *errDto,
		}
	}

	return &responseDto.ResponseDto{
		Content: projectModel,
	}
}

func (p *projectHandlerImpl) validateUser() *responseDto.ResponseDto {
	p.loadUserService()
	personalDto, errDto := p.userService.ValidateUserToken()
	if errDto != nil {
		return responseErrorManager(errDto)
	}
	p.contextDto.PersonalId = personalDto.Id
	return nil
}

func (p *projectHandlerImpl) loadUserService() {
	if p.userService == nil {
		p.projectService = projectService.BuildIProjectService(p.contextDto)
	}
}

func (p *projectHandlerImpl) loadProjectService() {
	if p.projectService == nil {
		p.projectService = projectService.BuildIProjectService(p.contextDto)
	}
}
