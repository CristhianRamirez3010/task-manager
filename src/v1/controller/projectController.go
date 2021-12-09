package controller

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/controller/impl"
	"github.com/gin-gonic/gin"
)

type IProjectController interface {
	GetProjects(c *gin.Context)
	NewProject(c *gin.Context)
}

func BuildIProjectController() IProjectController {
	return impl.BuildProjectControllerImpl()
}
