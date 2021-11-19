package impl

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handler"
	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	userHandler handler.IUserHandler
}

func BuildUserControllerImpl() UserControllerImpl {
	return UserControllerImpl{
		userHandler: handler.BuildIUserHandler(),
	}
}

func (u UserControllerImpl) GetDocuments(c *gin.Context) {

}
