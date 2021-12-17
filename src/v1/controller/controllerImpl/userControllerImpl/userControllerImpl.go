package userControllerImpl

import (
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/controller/controllerImpl"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/dto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handler"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	controllerImpl.ControllerBase
}

const (
	errConvertStruct = controllerImpl.ErrConvertStruct
	errDefault       = controllerImpl.ErrDefault
)

func BuildUserControllerImpl() *UserControllerImpl {

	return &UserControllerImpl{}
}

func (u *UserControllerImpl) GetDocuments(c *gin.Context) {
	userHandler := handler.BuildIUserHandler(u.LoadContextDto(c))
	u.ResponseManager(userHandler.GetDocuments(), c)
}

func (u *UserControllerImpl) ValidateLogin(c *gin.Context) {
	userHandler := handler.BuildIUserHandler(u.LoadContextDto(c))
	var userLogin useLoginModel.UseLoginModel
	err := c.BindJSON(&userLogin)
	if err != nil {
		u.ResponseManager(
			&responseDto.ResponseDto{
				Error: *utils.Logger(errConvertStruct, errDefault, http.StatusInternalServerError, err.Error()),
			},
			c)
	} else {
		u.ResponseManager(userHandler.ValidateLogin(&userLogin), c)
	}
}

func (u *UserControllerImpl) CreateNewUser(c *gin.Context) {
	userHandler := handler.BuildIUserHandler(u.LoadContextDto(c))
	var userData dto.UserdataDto
	err := c.BindJSON(&userData)
	if err != nil {
		u.ResponseManager(
			&responseDto.ResponseDto{
				Error: *utils.Logger(errConvertStruct, errDefault, http.StatusInternalServerError, err.Error()),
			},
			c)
	}

	u.ResponseManager(userHandler.CreateNewUser(&userData), c)
}
