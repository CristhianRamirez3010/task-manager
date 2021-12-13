package handler

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handler/impl"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/proProjectModel"
)

type IProjectHandler interface {
	GetProject() *responseDto.ResponseDto
	NewProject(projectModel *proProjectModel.ProProjectModel) *responseDto.ResponseDto
}

func BuildIProjectHandler(context *contextDto.ContextDto) IProjectHandler {
	return impl.BuildProjectHandlerImpl(context)
}
