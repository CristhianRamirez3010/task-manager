package projectService

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/proProjectModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository/proProjectRepo"
)

type IProjectService interface {
	GetProjectByContext() ([]*proProjectModel.ProProjectModel, *errorManagerDto.ErrorManagerDto)
	CreateNewProject(projectModel *proProjectModel.ProProjectModel) *errorManagerDto.ErrorManagerDto
}

func BuildIProjectService(contextDto *contextDto.ContextDto) IProjectService {
	service := new(projectServiceImpl)
	service.contextDto = contextDto
	return service
}

type projectServiceImpl struct {
	contextDto *contextDto.ContextDto

	projectRepo proProjectRepo.IProProjectRepo
}

var (
	buildIProProjectRepo = proProjectRepo.BuildIProProjectRepo
)

func (p *projectServiceImpl) GetProjectByContext() ([]*proProjectModel.ProProjectModel, *errorManagerDto.ErrorManagerDto) {
	p.projectRepo = buildIProProjectRepo()
	projectModelList, errDto := p.projectRepo.FindProjectByUserId(p.contextDto.PersonalId)
	if errDto != nil {
		return nil, errDto
	}
	return projectModelList, nil
}

func (p *projectServiceImpl) CreateNewProject(projectModel *proProjectModel.ProProjectModel) *errorManagerDto.ErrorManagerDto {
	p.projectRepo = buildIProProjectRepo()

	if errDto := p.projectRepo.New(projectModel); errDto != nil {
		return errDto
	}

	return nil
}
