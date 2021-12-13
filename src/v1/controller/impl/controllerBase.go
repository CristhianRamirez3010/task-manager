package impl

import (
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/gin-gonic/gin"
)

const (
	errDefault       = "Error with controller"
	errConvertStruct = "Error to convert Struct"
)

type controllerBase struct{}

func (_ controllerBase) responseManager(response *responseDto.ResponseDto, context *gin.Context) {
	if response.Error == (errorManagerDto.ErrorManagerDto{}) {
		context.JSON(http.StatusOK, response)
	} else {
		context.JSON(response.Error.Status, response)
	}
}
func (_ controllerBase) loadContextDto(con *gin.Context) *contextDto.ContextDto {
	contextDto := contextDto.ContextDto{
		ClientIP:    con.ClientIP(),
		AccessToken: con.GetHeader("accessId"),
	}
	return &contextDto
}
