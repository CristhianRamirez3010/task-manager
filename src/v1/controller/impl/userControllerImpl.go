package impl

import (
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
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
	response := u.userHandler.GetDocuments()
	if response.Error == (errorManagerDto.ErrorManagerDto{}) {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(response.Error.Status, response)
	}
}
