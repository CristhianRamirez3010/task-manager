package impl

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/exceptions"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/dto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useHistoryTokensModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/usePersonalDataModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository"
)

type UserHandler struct {
	contextDto *contextDto.ContextDto

	documentTypeRepo repository.IUseDocumentRepo
	loginRepo        repository.IUseLoginRepo
	personalDataRepo repository.IUsePersonalDataRepo
	historyTokeRepo  repository.IUseHistoryTokensRepo
}

func BuildUserHandler(con *contextDto.ContextDto) *UserHandler {
	return &UserHandler{
		contextDto:       con,
		documentTypeRepo: repository.BuildIUseDocumentRepo(),
		loginRepo:        repository.BuildIUseLoginRepo(),
		personalDataRepo: repository.BuildIUsePersonalDataRepo(),
		historyTokeRepo:  repository.BuildIUseHistoryTokensRepo(),
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
			Error:   *utils.Logger(exceptions.OBJECT_EMPTY, exceptions.OBJECT_EMPTY, http.StatusPreconditionFailed, ""),
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

	token, errDto := u.createHashToke(content[0])
	if errDto != nil {
		return &responseDto.ResponseDto{
			Error: *errDto,
		}
	}

	errDto = u.saveUserToken(token, content[0])
	if errDto != nil {
		return &responseDto.ResponseDto{
			Error: *errDto,
		}
	}

	personalData, errDto := u.personalDataRepo.GetDataByLoginId(&content[0].Id)
	if errDto != nil {
		return &responseDto.ResponseDto{
			Error: *errDto,
		}
	}

	return &responseDto.ResponseDto{
		Content: &dto.UserdataDto{
			Name:    personalData.Name,
			Surname: personalData.Surname,
			Country: personalData.Country,
			Phone:   personalData.Phone,
			Token:   token,
			Email:   content[0].Email,
		},
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
		Name:         userData.Name,
		Surname:      userData.Surname,
		Document:     userData.Document,
		Phone:        userData.Phone,
		Country:      userData.Country,
		DocumentType: userData.DocumentType,
		LoginId:      loginResponse[0].Id,
		UserRegister: "",
		DateRegister: time.Now(),
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

func (u *UserHandler) createHashToke(loginModel *useLoginModel.UseLoginModel) (string, *errorManagerDto.ErrorManagerDto) {
	hasher := md5.New()
	_, err := hasher.Write([]byte(fmt.Sprintf("%d%s%s", time.Now().UnixNano(), loginModel.Email, loginModel.User)))
	if err != nil {
		return "", utils.Logger("Error with the hash method", errDefault, http.StatusInternalServerError, err.Error())
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func (u *UserHandler) saveUserToken(token string, user *useLoginModel.UseLoginModel) *errorManagerDto.ErrorManagerDto {
	date := time.Now()
	date.AddDate(0, 0, 1)
	errDto := u.historyTokeRepo.NewToken(&useHistoryTokensModel.UseHistoryToekensModel{
		Token:        token,
		Finish:       date,
		LoginId:      user.Id,
		UserRegister: "",
		DateRegister: time.Now(),
	})
	if errDto != nil {
		return errDto
	}
	return nil
}
