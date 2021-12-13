package impl

import (
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handler"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/proProjectModel"
	"github.com/gin-gonic/gin"
)

type ProjectControllerImpl struct {
	controllerBase
}

func BuildProjectControllerImpl() *ProjectControllerImpl {

	return &ProjectControllerImpl{}
}

func (p *ProjectControllerImpl) GetProjects(c *gin.Context) {
	projectHandler := handler.BuildIProjectHandler(p.loadContextDto(c))
	p.responseManager(projectHandler.GetProject(), c)
}

func (p *ProjectControllerImpl) NewProject(c *gin.Context) {
	projectHandler := handler.BuildIProjectHandler(p.loadContextDto(c))
	var projectModel proProjectModel.ProProjectModel
	err := c.BindJSON(&projectHandler)
	if err != nil {
		p.responseManager(&responseDto.ResponseDto{
			Error: *utils.Logger(errConvertStruct, errDefault, http.StatusInternalServerError, err.Error()),
		}, c)
	}
	p.responseManager(projectHandler.NewProject(&projectModel), c)
}
