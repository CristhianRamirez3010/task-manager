package repository

import "github.com/CristhianRamirez3010/task-manager-go/src/v1/models"

type UseDocumentRepo interface {
	GetAll() []models.UseDocumentModel
}
