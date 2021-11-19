package handler

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handler/impl"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models"
)

type IUserHandler interface {
	GetDocuments() []models.UseDocumentModel
}

func BuildIUserHandler() IUserHandler {
	return impl.BuildUserHandler()
}
