package controllerImpl

import (
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/gin-gonic/gin"
)

const (
	ErrDefault       = "Error with controller"
	ErrConvertStruct = "Error to convert Struct"
)

type ControllerBase struct{}

func (_ ControllerBase) ResponseManager(response *responseDto.ResponseDto, context *gin.Context) {
	if response.Error == (errorManagerDto.ErrorManagerDto{}) {
		context.JSON(http.StatusOK, response)
	} else {
		context.JSON(response.Error.Status, response)
	}
}
func (_ ControllerBase) LoadContextDto(con *gin.Context) *contextDto.ContextDto {
	contextDto := contextDto.ContextDto{
		ClientIP:    con.ClientIP(),
		AccessToken: con.GetHeader("accessId"),
	}
	return &contextDto
}
