package controller

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/controller/impl"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetDocuments(c *gin.Context)
	ValidateLogin(c *gin.Context)
	CreateNewUser(c *gin.Context)
}

func BuildIUserController() IUserController {
	return impl.BuildUserControllerImpl()
}
