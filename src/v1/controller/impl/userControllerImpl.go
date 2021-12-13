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
	controllerBase
}

func BuildUserControllerImpl() *UserControllerImpl {

	return &UserControllerImpl{}
}

func (u *UserControllerImpl) GetDocuments(c *gin.Context) {
	userHandler := handler.BuildIUserHandler(u.loadContextDto(c))
	u.responseManager(userHandler.GetDocuments(), c)
}

func (u *UserControllerImpl) ValidateLogin(c *gin.Context) {
	userHandler := handler.BuildIUserHandler(u.loadContextDto(c))
	var userLogin useLoginModel.UseLoginModel
	err := c.BindJSON(&userLogin)
	if err != nil {
		u.responseManager(
			&responseDto.ResponseDto{
				Error: *utils.Logger(errConvertStruct, errDefault, http.StatusInternalServerError, err.Error()),
			},
			c)
	} else {
		u.responseManager(userHandler.ValidateLogin(&userLogin), c)
	}
}

func (u *UserControllerImpl) CreateNewUser(c *gin.Context) {
	userHandler := handler.BuildIUserHandler(u.loadContextDto(c))
	var userData dto.UserdataDto
	err := c.BindJSON(&userData)
	if err != nil {
		u.responseManager(
			&responseDto.ResponseDto{
				Error: *utils.Logger(errConvertStruct, errDefault, http.StatusInternalServerError, err.Error()),
			},
			c)
	}

	u.responseManager(userHandler.CreateNewUser(&userData), c)
}
