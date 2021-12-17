package projectControllerImpl

import (
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/controller/controllerImpl"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handler"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/proProjectModel"
	"github.com/gin-gonic/gin"
)

type ProjectControllerImpl struct {
	controllerImpl.ControllerBase
}

const (
	errConvertStruct = controllerImpl.ErrConvertStruct
	errDefault       = controllerImpl.ErrDefault
)

var (
	buildIProjectHandler = handler.BuildIProjectHandler
)

func BuildProjectControllerImpl() *ProjectControllerImpl {

	return new(ProjectControllerImpl)
}

func (p *ProjectControllerImpl) GetProjects(c *gin.Context) {
	projectHandler := buildIProjectHandler(p.LoadContextDto(c))
	p.ResponseManager(projectHandler.GetProject(), c)
}

func (p *ProjectControllerImpl) NewProject(c *gin.Context) {
	projectHandler := buildIProjectHandler(p.LoadContextDto(c))
	var projectModel proProjectModel.ProProjectModel
	err := c.BindJSON(&projectHandler)
	if err != nil {
		p.ResponseManager(&responseDto.ResponseDto{
			Error: *utils.Logger(errConvertStruct, errDefault, http.StatusInternalServerError, err.Error()),
		}, c)
	}
	p.ResponseManager(projectHandler.NewProject(&projectModel), c)
}
