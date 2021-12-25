package userController

import (
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/constants"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/exceptions"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/dto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handlers/userHandler"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type IUserController interface {
	ValidateLogin(c *gin.Context)
	CreateNewUser(c *gin.Context)
}

func BuildIUserController() IUserController {
	return new(userControllerImpl)
}

type userControllerImpl struct{}

var (
	buildIUserHandler = userHandler.BuildIUserHandler
	responseManager   = utils.ResponseManager
)

func (u *userControllerImpl) ValidateLogin(c *gin.Context) {
	contextDto := new(contextDto.ContextDto)
	contextDto.ClientIP = c.ClientIP()
	contextDto.AccessToken = c.GetHeader(constants.ACCESS_ID)
	userHandler := buildIUserHandler(contextDto)
	var userLogin useLoginModel.UseLoginModel

	if err := binding.JSON.Bind(c.Request, &userLogin); err != nil {
		c.JSON(http.StatusInternalServerError,
			&responseDto.ResponseDto{
				Error: *utils.Logger(exceptions.ERR_CONVERT_STRUCT,
					exceptions.ERR_DEFAULT_CONTROLLER,
					http.StatusInternalServerError,
					err.Error()),
			})
	} else {
		c.JSON(responseManager(userHandler.ValidateLogin(&userLogin)))
	}
}

func (u *userControllerImpl) CreateNewUser(c *gin.Context) {
	contextDto := new(contextDto.ContextDto)
	contextDto.ClientIP = c.ClientIP()
	contextDto.AccessToken = c.GetHeader(constants.ACCESS_ID)
	userHandler := buildIUserHandler(contextDto)
	var userData dto.UserdataDto

	if err := binding.JSON.Bind(c.Request, &userData); err != nil {
		c.JSON(http.StatusInternalServerError,
			&responseDto.ResponseDto{
				Error: *utils.Logger(exceptions.ERR_CONVERT_STRUCT,
					exceptions.ERR_DEFAULT_CONTROLLER,
					http.StatusInternalServerError,
					err.Error()),
			})
	}

	c.JSON(responseManager(userHandler.CreateNewUser(&userData)))
}
