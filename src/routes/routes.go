package routes

import (
	"net/http"

	v1 "github.com/CristhianRamirez3010/task-manager-go/src/v1"
	"github.com/gin-gonic/gin"
)

func LoadRoutes() http.Handler {
	gin := gin.Default()
	v1.BuildIApi().Routing(gin)
	return gin
}
