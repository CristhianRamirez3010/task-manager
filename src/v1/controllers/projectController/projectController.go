package projectController

import (
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/constants"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/exceptions"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handlers/projectHandler"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/proProjectModel"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type IProjectController interface {
	GetProjects(c *gin.Context)
	NewProject(c *gin.Context)
}

func BuildIProjectController() IProjectController {
	return new(projectControllerImpl)
}

type projectControllerImpl struct{}

var (
	buildIProjectHandler = projectHandler.BuildIProjectHandler
	responseManager      = utils.ResponseManager
)

func (p *projectControllerImpl) GetProjects(c *gin.Context) {
	contextDto := new(contextDto.ContextDto)
	contextDto.ClientIP = c.ClientIP()
	contextDto.AccessToken = c.GetHeader(constants.ACCESS_ID)
	projectHandler := buildIProjectHandler(contextDto)
	c.JSON(responseManager(projectHandler.GetProject()))
}

func (p *projectControllerImpl) NewProject(c *gin.Context) {
	contextDto := new(contextDto.ContextDto)
	contextDto.ClientIP = c.ClientIP()
	contextDto.AccessToken = c.GetHeader(constants.ACCESS_ID)
	projectHandler := buildIProjectHandler(contextDto)
	var projectModel proProjectModel.ProProjectModel

	if err := binding.JSON.Bind(c.Request, &projectModel); err != nil {
		c.JSON(500, &responseDto.ResponseDto{
			Error: *utils.Logger(exceptions.ERR_CONVERT_STRUCT, exceptions.ERR_DEFAULT_CONTROLLER, http.StatusInternalServerError, err.Error()),
		})
	} else {
		c.JSON(responseManager(projectHandler.NewProject(&projectModel)))
	}
}
