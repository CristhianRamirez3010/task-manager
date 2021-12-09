package api

import (
	"fmt"

	"github.com/CristhianRamirez3010/task-manager-go/src/v1/controller"
	"github.com/gin-gonic/gin"
)

type Api struct{}

func BuildApi() *Api {
	return &Api{}
}

func (a *Api) Routing(gin *gin.Engine) {
	a.userApi(gin)
	a.projectApi(gin)
}

func (a *Api) userApi(gin *gin.Engine) {
	userController := controller.BuildIUserController()
	const userEndpoint = "/v1/user"

	gin.GET(fmt.Sprintf("%s/documents", userEndpoint), userController.GetDocuments)

	gin.POST(fmt.Sprintf("%s/login", userEndpoint), userController.ValidateLogin)
	gin.POST(fmt.Sprintf("%s/create", userEndpoint), userController.CreateNewUser)
}

func (a *Api) projectApi(gin *gin.Engine) {
	projectController := controller.BuildIProjectController()
	const projectEnpoint = "/v1/projects"

	gin.GET(projectEnpoint, projectController.GetProjects)
}
