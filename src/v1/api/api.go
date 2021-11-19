package api

import (
	"github.com/gin-gonic/gin"
)

type Api struct{}

func BuildApi() Api {
	return Api{}
}

func (a Api) Routing(gin *gin.Engine) {
	a.userApi(gin)
}

func (a Api) userApi(gin *gin.Engine) {
	const userEndpoint = "/v1/user/"
}
