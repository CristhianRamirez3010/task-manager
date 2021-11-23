package api

import (
	"fmt"

	"github.com/CristhianRamirez3010/task-manager-go/src/v1/controller"
	"github.com/gin-gonic/gin"
)

type Api struct {
	userController controller.IUserController
}

func BuildApi() *Api {
	return &Api{
		userController: controller.BuildIUserController(),
	}
}

func (a *Api) Routing(gin *gin.Engine) {
	a.userApi(gin)
}

func (a *Api) userApi(gin *gin.Engine) {
	const userEndpoint = "/v1/user"
	gin.GET(fmt.Sprintf("%s/documents", userEndpoint), a.userController.GetDocuments)

	gin.POST(fmt.Sprintf("%s/login", userEndpoint), a.userController.ValidateLogin)
	gin.POST(fmt.Sprintf("%s/create", userEndpoint), a.userController.CreateNewUser)
}
