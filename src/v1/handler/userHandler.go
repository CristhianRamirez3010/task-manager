package handler

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/dto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handler/impl"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
)

type IUserHandler interface {
	GetDocuments() *responseDto.ResponseDto
	ValidateLogin(userLogin *useLoginModel.UseLoginModel) *responseDto.ResponseDto
	CreateNewUser(userData *dto.UserdataDto) *responseDto.ResponseDto
}

func BuildIUserHandler() IUserHandler {
	return impl.BuildUserHandler()
}
