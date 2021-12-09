package impl

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handler"
	"github.com/gin-gonic/gin"
)

type ProjectControllerImpl struct {
	controllerBase
}

func BuildProjectControllerImpl() *ProjectControllerImpl {
	con := contextDto.ContextDto{}
	return &ProjectControllerImpl{
		controllerBase: controllerBase{contextDto: &con},
	}
}

func (p *ProjectControllerImpl) GetProjects(c *gin.Context) {
	p.loadContextDto(c)
	projectHandler := handler.BuildIProjectHandler(p.contextDto)
	p.responseManager(projectHandler.GetProject(), c)
}

func (p *ProjectControllerImpl) NewProject(c *gin.Context) {
	projectHandler := handler.BuildIProjectHandler(p.contextDto)
	p.responseManager(projectHandler.NewProject(), c)
}
