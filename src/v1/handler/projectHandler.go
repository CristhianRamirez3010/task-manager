package handler

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handler/impl"
)

type IProjectHandler interface {
	GetProject() *responseDto.ResponseDto
	NewProject() *responseDto.ResponseDto
}

func BuildIProjectHandler(context *contextDto.ContextDto) IProjectHandler {
	return impl.BuildProjectHandlerImpl(context)
}
