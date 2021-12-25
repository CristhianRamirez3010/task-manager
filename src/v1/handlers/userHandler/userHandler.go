package userHandler

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/dto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/service/userService"
)

type IUserHandler interface {
	ValidateLogin(userLogin *useLoginModel.UseLoginModel) *responseDto.ResponseDto
	CreateNewUser(userData *dto.UserdataDto) *responseDto.ResponseDto
}

func BuildIUserHandler(con *contextDto.ContextDto) IUserHandler {
	handler := new(userHandlerImpl)
	handler.contextDto = con
	return handler
}

type userHandlerImpl struct {
	contextDto *contextDto.ContextDto

	userService userService.IUserService
}

var (
	responseErrorManager          = utils.ResponseErrorManager
	responseMessageManager        = utils.ResponseMessageManager
	responseMssgAndContentManager = utils.ResponseMssgAndContentManager
)

func (u *userHandlerImpl) ValidateLogin(userLogin *useLoginModel.UseLoginModel) *responseDto.ResponseDto {
	u.loadUserService()
	if _, errDto := u.userService.ValidateLoginWithModel(userLogin); errDto != nil {
		return responseErrorManager(errDto)
	}

	token, errDto := u.userService.CreateNewToken(userLogin)
	if errDto != nil {
		return responseErrorManager(errDto)
	}

	errDto = u.userService.SaveUserToken(token, userLogin.Id)
	if errDto != nil {
		return responseErrorManager(errDto)
	}

	personalData, errDto := u.userService.FindUserWithLoginId(userLogin.Id)
	if errDto != nil {
		return responseErrorManager(errDto)
	}

	userDto := new(dto.UserdataDto)
	userDto.Name = personalData.Name
	userDto.Surname = personalData.Surname
	userDto.Country = personalData.Country
	userDto.Phone = personalData.Phone
	userDto.Token = token
	userDto.Email = userLogin.Email
	return responseMssgAndContentManager("User Valid", userDto)
}

func (u *userHandlerImpl) CreateNewUser(userData *dto.UserdataDto) *responseDto.ResponseDto {
	u.loadUserService()
	if errDto := u.userService.CreateNewUser(userData); errDto != nil {
		return responseErrorManager(errDto)
	}

	return responseMessageManager("Register success")
}

func (u *userHandlerImpl) loadUserService() {
	if u.userService == nil {
		u.userService = userService.BuildIUserService(u.contextDto)
	}
}
