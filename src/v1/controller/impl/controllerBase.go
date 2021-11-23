package impl

import (
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/gin-gonic/gin"
)

const (
	errDefault = "Error with controller"
)

type controllerBase struct{}

func (c controllerBase) ResponseManager(response responseDto.ResponseDto, context *gin.Context) {
	if response.Error == (errorManagerDto.ErrorManagerDto{}) {
		context.JSON(http.StatusOK, response)
	} else {
		context.JSON(response.Error.Status, response)
	}
}
