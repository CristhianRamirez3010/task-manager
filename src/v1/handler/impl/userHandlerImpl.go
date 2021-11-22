package impl

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository"
)

type UserHandler struct {
	documentTypeRepo repository.IUseDocumentRepo
}

func BuildUserHandler() UserHandler {
	return UserHandler{
		documentTypeRepo: repository.BuildIUseDocumentRepo(),
	}
}

func (u UserHandler) GetDocuments() responseDto.DesponseDto {
	content, errorDto := u.documentTypeRepo.GetAll()
	if errorDto != nil {
		return responseDto.DesponseDto{
			Content: content,
			Error:   *errorDto,
		}
	}

	return responseDto.DesponseDto{
		Content: content,
	}
}
