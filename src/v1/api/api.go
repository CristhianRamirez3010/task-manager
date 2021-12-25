package api

import (
	"fmt"

	"github.com/CristhianRamirez3010/task-manager-go/src/v1/controllers/projectController"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/controllers/userController"
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
	controller := userController.BuildIUserController()
	const userEndpoint = "/v1/user"

	gin.POST(fmt.Sprintf("%s/login", userEndpoint), controller.ValidateLogin)
	gin.POST(fmt.Sprintf("%s/create", userEndpoint), controller.CreateNewUser)
}

func (a *Api) projectApi(gin *gin.Engine) {

	const projectEnpoint = "/v1/projects"
	controller := projectController.BuildIProjectController()
	gin.GET(projectEnpoint, controller.GetProjects)
	gin.POST(fmt.Sprintf("%s/new", projectEnpoint), controller.NewProject)
}
