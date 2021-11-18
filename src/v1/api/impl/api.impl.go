package impl

import "github.com/gin-gonic/gin"

type Api struct{}

func BuildApi() Api {
	return Api{}
}

func (a *Api) Routing() {
	gin := gin.Default()
}
