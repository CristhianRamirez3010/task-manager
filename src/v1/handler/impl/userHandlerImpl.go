package impl

import (
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository"
)

type UserHandler struct {
	documentTypeRepo repository.IUseDocumentRepo
	loginRepo        repository.IUseLoginRepo
}

func BuildUserHandler() UserHandler {
	return UserHandler{
		documentTypeRepo: repository.BuildIUseDocumentRepo(),
		loginRepo:        repository.BuildIUseLoginRepo(),
	}
}

func (u UserHandler) GetDocuments() responseDto.ResponseDto {
	content, errorDto := u.documentTypeRepo.GetAll()
	if errorDto != nil {
		return responseDto.ResponseDto{
			Content: content,
			Error:   *errorDto,
		}
	}

	return responseDto.ResponseDto{
		Content: content,
	}
}

func (u UserHandler) ValidateLogin(userLogin *useLoginModel.UseLoginModel) responseDto.ResponseDto {
	if userLogin == nil && (userLogin.Email != "" || userLogin.User != "") && userLogin.Password != "" {
		return responseDto.ResponseDto{
			Error:   *utils.Logger("Object is empty", "Object is empty", http.StatusPreconditionRequired, ""),
			Message: "The fields are requerid",
		}
	}

	content, errDto := u.loginRepo.FindByEmailAndUser(userLogin.Email, userLogin.User)
	if errDto != nil {
		return responseDto.ResponseDto{
			Error: *errDto,
		}
	}

	if len(content) == 0 || *&content[0].Password == userLogin.Password {
		return responseDto.ResponseDto{
			Message: "User or Password Invalid",
		}
	}

	return responseDto.ResponseDto{
		Message: "User Valid",
	}
}
