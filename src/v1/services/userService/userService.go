package userService

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/exceptions"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/dto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useHistoryTokensModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/usePersonalDataModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository/useHistoryTokensRepo"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository/useLoginRepo"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository/usePersonalDataRepo"
)

type IUserService interface {
	ValidateUserToken() (*usePersonalDataModel.UsePersonalDataModel, *errorManagerDto.ErrorManagerDto)
	ValidateLoginWithModel(userLogin *useLoginModel.UseLoginModel) (*useLoginModel.UseLoginModel, *errorManagerDto.ErrorManagerDto)
	CreateNewToken(loginModel *useLoginModel.UseLoginModel) (string, *errorManagerDto.ErrorManagerDto)
	SaveUserToken(token string, userId int64) *errorManagerDto.ErrorManagerDto
	FindUserWithLoginId(loginId int64) (*usePersonalDataModel.UsePersonalDataModel, *errorManagerDto.ErrorManagerDto)
	CreateNewUser(userData *dto.UserdataDto) *errorManagerDto.ErrorManagerDto
}

func BuildIUserService(contextDto *contextDto.ContextDto) IUserService {
	service := new(userServiceImpl)
	service.contextDto = contextDto
	return service
}

type userServiceImpl struct {
	contextDto *contextDto.ContextDto

	loginRepo        useLoginRepo.IUseLoginRepo
	historyTokenRepo useHistoryTokensRepo.IUseHistoryTokensRepo
	personalDataRepo usePersonalDataRepo.IUsePersonalDataRepo
}

const (
	errUserAndPassInvalid = "User or Password Invalid"
	errEmailOrUserExists  = "that user or email exists"
	errCreationUser       = "Problems with the creation user"
)

var (
	logger = utils.Logger
)

func (u *userServiceImpl) ValidateUserToken() (*usePersonalDataModel.UsePersonalDataModel, *errorManagerDto.ErrorManagerDto) {
	if u.contextDto.AccessToken == "" {
		return nil, utils.Logger(exceptions.ERR_TOKEN_INVALID, exceptions.ERR_TOKEN_INVALID, http.StatusPreconditionFailed, "")
	}

	u.loadHistoryTokenRepo()
	u.loadPersonalDataRepo()
	perosnalDataModel, errDto := u.personalDataRepo.GetDataByToken(u.contextDto.AccessToken)
	if errDto != nil {
		return nil, errDto
	} else if perosnalDataModel.Id < 1 {
		return nil, utils.Logger(exceptions.ERR_TOKEN_INVALID, exceptions.ERR_TOKEN_INVALID, http.StatusPreconditionFailed, "")
	}

	historyTokenModel, errDto := u.historyTokenRepo.GetLastTokenByPersonId(perosnalDataModel.Id)
	if errDto != nil {
		return nil, errDto
	} else if historyTokenModel.Id < 1 {
		return nil, utils.Logger("token don't found in data base", exceptions.ERR_TOKEN_INVALID, http.StatusPreconditionFailed, "")
	} else if historyTokenModel.Token != u.contextDto.AccessToken {
		return nil, utils.Logger(exceptions.ERR_TOKEN_INVALID, exceptions.ERR_TOKEN_INVALID, http.StatusPreconditionFailed, "")
	}

	return perosnalDataModel, nil
}

func (u *userServiceImpl) ValidateLoginWithModel(userLogin *useLoginModel.UseLoginModel) (*useLoginModel.UseLoginModel, *errorManagerDto.ErrorManagerDto) {
	if userLogin == nil && (userLogin.Email != "" || userLogin.User != "") && userLogin.Password != "" {
		return nil, logger(exceptions.OBJECT_EMPTY, exceptions.OBJECT_EMPTY, http.StatusPreconditionFailed, "")
	}
	u.loadLoginRepo()
	content, errDto := u.loginRepo.FindByEmailAndUser(userLogin.Email, userLogin.User)
	if errDto != nil {
		return nil, errDto
	}

	if len(content) == 0 || content[0].Password != userLogin.Password {
		return nil, logger(errUserAndPassInvalid, errUserAndPassInvalid, http.StatusPreconditionRequired, "")
	}

	userLogin = content[0]

	return userLogin, nil
}

func (u *userServiceImpl) CreateNewToken(loginModel *useLoginModel.UseLoginModel) (string, *errorManagerDto.ErrorManagerDto) {
	hasher := md5.New()
	_, err := hasher.Write([]byte(fmt.Sprintf("%d%s%s", time.Now().UnixNano(), loginModel.Email, loginModel.User)))
	if err != nil {
		return "", utils.Logger("Error with the hash method",
			exceptions.ERR_DEFAULT_HANDLER,
			http.StatusInternalServerError,
			err.Error())
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func (u *userServiceImpl) SaveUserToken(token string, userId int64) *errorManagerDto.ErrorManagerDto {
	date := time.Now()
	date.AddDate(0, 0, 1)
	u.loadHistoryTokenRepo()
	historyTokenModel := new(useHistoryTokensModel.UseHistoryToekensModel)
	historyTokenModel.Token = token
	historyTokenModel.Finish = date
	historyTokenModel.LoginId = userId
	historyTokenModel.UserRegister = ""
	historyTokenModel.DateRegister = time.Now()

	return u.historyTokenRepo.NewToken(historyTokenModel)
}

func (u *userServiceImpl) CreateNewUser(userData *dto.UserdataDto) *errorManagerDto.ErrorManagerDto {
	loginModel := new(useLoginModel.UseLoginModel)
	loginModel.Email = userData.Email
	loginModel.User = userData.User
	loginModel.Password = userData.Password
	loginModel.UserRegister = ""
	loginModel.DateRegister = time.Now()

	u.loadLoginRepo()
	if user, errDto := u.loginRepo.FindByEmailAndUser(loginModel.Email, loginModel.User); errDto != nil {
		return errDto
	} else if len(user) > 0 {
		return logger(errEmailOrUserExists, errEmailOrUserExists, http.StatusConflict, "")
	}

	if errDto := u.loginRepo.New(loginModel); errDto != nil {
		return errDto
	} else if loginModel.Id <= 0 {
		return logger(errCreationUser, exceptions.ERR_DEFAULT_SERVICE, http.StatusInternalServerError, "")
	}

	personalDataModel := new(usePersonalDataModel.UsePersonalDataModel)
	personalDataModel.Name = userData.Name
	personalDataModel.Surname = userData.Surname
	personalDataModel.Document = userData.Document
	personalDataModel.Phone = userData.Phone
	personalDataModel.Country = userData.Country
	personalDataModel.DocumentType = userData.DocumentType
	personalDataModel.LoginId = loginModel.Id
	personalDataModel.UserRegister = ""
	personalDataModel.DateRegister = time.Now()

	u.loadPersonalDataRepo()
	if errDto := u.personalDataRepo.New(personalDataModel); errDto != nil {
		return errDto
	} else if personalDataModel.Id <= 0 {
		return logger(errCreationUser, exceptions.ERR_DEFAULT_SERVICE, http.StatusInternalServerError, "")
	}

	return nil
}

func (u *userServiceImpl) FindUserWithLoginId(loginId int64) (*usePersonalDataModel.UsePersonalDataModel, *errorManagerDto.ErrorManagerDto) {
	u.loadPersonalDataRepo()
	return u.personalDataRepo.GetDataByLoginId(loginId)
}

func (u *userServiceImpl) loadLoginRepo() {
	if u.loginRepo == nil {
		u.loginRepo = useLoginRepo.BuildIUseLoginRepo()
	}
}
func (u *userServiceImpl) loadHistoryTokenRepo() {
	if u.historyTokenRepo == nil {
		u.historyTokenRepo = useHistoryTokensRepo.BuildIUseHistoryTokensRepo()
	}
}
func (u *userServiceImpl) loadPersonalDataRepo() {
	if u.personalDataRepo != nil {
		u.personalDataRepo = usePersonalDataRepo.BuildIUsePersonalDataRepo()
	}
}
