package impl

import (
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/dto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handler"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	userHandler handler.IUserHandler
	controllerBase
}

func BuildUserControllerImpl() *UserControllerImpl {
	return &UserControllerImpl{
		userHandler: handler.BuildIUserHandler(),
	}
}

func (u *UserControllerImpl) GetDocuments(c *gin.Context) {
	u.ResponseManager(u.userHandler.GetDocuments(), c)
}

func (u *UserControllerImpl) ValidateLogin(c *gin.Context) {
	var userLogin useLoginModel.UseLoginModel
	err := c.BindJSON(&userLogin)
	if err != nil {
		u.ResponseManager(
			&responseDto.ResponseDto{
				Error: *utils.Logger(errConvertStruct, errDefault, http.StatusInternalServerError, err.Error()),
			},
			c)
	} else {
		u.ResponseManager(u.userHandler.ValidateLogin(&userLogin), c)
	}
}

func (u *UserControllerImpl) CreateNewUser(c *gin.Context) {
	var userData dto.UserdataDto
	err := c.BindJSON(&userData)
	if err != nil {
		u.ResponseManager(
			&responseDto.ResponseDto{
				Error: *utils.Logger(errConvertStruct, errDefault, http.StatusInternalServerError, err.Error()),
			},
			c)
	}

	u.ResponseManager(u.userHandler.CreateNewUser(&userData), c)
}
