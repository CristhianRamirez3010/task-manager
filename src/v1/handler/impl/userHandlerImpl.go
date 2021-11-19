package impl

import "github.com/CristhianRamirez3010/task-manager-go/src/v1/models"

type UserHandler struct{}

func BuildUserHandler() UserHandler {
	return UserHandler{}
}

func (u UserHandler) GetDocuments() []models.UseDocumentModel {
	return nil
}
