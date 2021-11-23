package impl

import (
	"fmt"
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/gin-gonic/gin"
)

const (
	errDefault       = "Error with controller"
	errConvertStruct = "Error to convert Struct"
)

type controllerBase struct{}

func (c *controllerBase) ResponseManager(response *responseDto.ResponseDto, context *gin.Context) {
	fmt.Println(response)
	if response.Error == (errorManagerDto.ErrorManagerDto{}) {
		context.JSON(http.StatusOK, response)
	} else {
		context.JSON(response.Error.Status, response)
	}
}
