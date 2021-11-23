package impl

import (
	"net/http"
	"time"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/exceptios"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/dto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/usePersonalDataModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository"
)

type UserHandler struct {
	documentTypeRepo repository.IUseDocumentRepo
	loginRepo        repository.IUseLoginRepo
	personalDataRepo repository.IUsePersonalDataRepo
}

func BuildUserHandler() *UserHandler {
	return &UserHandler{
		documentTypeRepo: repository.BuildIUseDocumentRepo(),
		loginRepo:        repository.BuildIUseLoginRepo(),
		personalDataRepo: repository.BuildIUsePersonalDataRepo(),
	}
}

func (u *UserHandler) GetDocuments() *responseDto.ResponseDto {
	content, errorDto := u.documentTypeRepo.GetAll()
	if errorDto != nil {
		return &responseDto.ResponseDto{
			Content: content,
			Error:   *errorDto,
		}
	}

	return &responseDto.ResponseDto{
		Content: content,
	}
}

func (u *UserHandler) ValidateLogin(userLogin *useLoginModel.UseLoginModel) *responseDto.ResponseDto {
	if userLogin == nil && (userLogin.Email != "" || userLogin.User != "") && userLogin.Password != "" {
		return &responseDto.ResponseDto{
			Error:   *utils.Logger(exceptios.OBJECT_EMPTY, exceptios.OBJECT_EMPTY, http.StatusPreconditionFailed, ""),
			Message: "The fields are requerid",
		}
	}

	content, errDto := u.loginRepo.FindByEmailAndUser(&userLogin.Email, &userLogin.User)
	if errDto != nil {
		return &responseDto.ResponseDto{
			Error: *errDto,
		}
	}

	if len(content) == 0 || content[0].Password != userLogin.Password {
		return &responseDto.ResponseDto{
			Error:   errorManagerDto.ErrorManagerDto{Status: http.StatusPreconditionRequired},
			Message: "User or Password Invalid",
		}
	}

	return &responseDto.ResponseDto{
		Message: "User Valid",
	}
}

func (u *UserHandler) CreateNewUser(userData *dto.UserdataDto) *responseDto.ResponseDto {

	loginModel := useLoginModel.UseLoginModel{
		Email:        userData.Email,
		User:         userData.User,
		Password:     userData.Password,
		UserRegister: "",
		DateRegister: time.Now(),
	}
	loginResponse, errDto := u.loginRepo.FindByEmailAndUser(&loginModel.Email, &loginModel.User)
	if errDto != nil {
		return &responseDto.ResponseDto{
			Error: *errDto,
		}
	}
	if len(loginResponse) > 0 {
		return &responseDto.ResponseDto{
			Error: *utils.Logger("that user or email exists", "that user or email exists", http.StatusConflict, ""),
		}
	}
	errDto = u.loginRepo.New(&loginModel)
	if errDto != nil {
		return &responseDto.ResponseDto{
			Error: *errDto,
		}
	}
	loginResponse, errDto = u.loginRepo.FindByEmailAndUser(&loginModel.Email, &loginModel.User)
	if errDto != nil {
		return &responseDto.ResponseDto{
			Error: *errDto,
		}
	}
	if len(loginResponse) < 1 {
		return &responseDto.ResponseDto{
			Error: *utils.Logger("LoginList is empty", errDefault, http.StatusInternalServerError, ""),
		}
	}

	personalDataModel := usePersonalDataModel.UsePersonalDataModel{
		Name:           userData.Name,
		Surname:        userData.Surname,
		Identification: userData.Idenfication,
		Phone:          userData.Phone,
		Country:        userData.Country,
		LoginId:        loginResponse[0].Id,
		UserRegister:   "",
		DateRegister:   time.Now(),
	}

	errDto = u.personalDataRepo.New(&personalDataModel)
	if errDto != nil {
		return &responseDto.ResponseDto{
			Error: *errDto,
		}
	}

	return &responseDto.ResponseDto{
		Message: "Register success",
	}
}
