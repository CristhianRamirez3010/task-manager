package impl

import (
	"fmt"
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

type controllerBase struct {
	contextDto *contextDto.ContextDto
}

func (c *controllerBase) responseManager(response *responseDto.ResponseDto, context *gin.Context) {
	fmt.Println(response)
	if response.Error == (errorManagerDto.ErrorManagerDto{}) {
		context.JSON(http.StatusOK, response)
	} else {
		context.JSON(response.Error.Status, response)
	}
}
func (c *controllerBase) loadContextDto(con *gin.Context) {
	c.contextDto.ClientIP = con.ClientIP()
	c.contextDto.AccessToken = con.GetHeader("accessId")
}
