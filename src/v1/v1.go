package v1

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/api"
	"github.com/gin-gonic/gin"
)

type IApi interface {
	Routing(g *gin.Engine)
}

func BuildIApi() IApi {
	return api.BuildApi()
}
