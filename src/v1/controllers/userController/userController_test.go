package userController

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/dto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handlers/userHandler"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	responseMock = new(responseDto.ResponseDto)
)

type userhandlerMock struct {
	userHandler.IUserHandler
	mock.Mock
}

func (u *userhandlerMock) ValidateLogin(userLogin *useLoginModel.UseLoginModel) *responseDto.ResponseDto {
	return responseMock
}

func (u *userhandlerMock) CreateNewUser(userData *dto.UserdataDto) *responseDto.ResponseDto {
	return responseMock
}

func TestBuildIUserController(t *testing.T) {
	res := BuildIUserController()
	assert.NotNil(t, res)
}

func TestValidateLogin(t *testing.T) {
	controller := new(userControllerImpl)
	gin := gin.Default()
	gin.POST("/", controller.ValidateLogin)
	hts := httptest.NewServer(gin)

	buildIUserHandler = func(con *contextDto.ContextDto) userHandler.IUserHandler {
		return new(userhandlerMock)
	}
	responseManager = func(response *responseDto.ResponseDto) (int, interface{}) {
		return 200, responseMock
	}

	buffer := new(bytes.Buffer)
	userLogin := new(useLoginModel.UseLoginModel)
	userLogin.Email = "test@test.com"
	userLogin.User = "test"
	userLogin.Password = "***"
	json.NewEncoder(buffer).Encode(userLogin)
	res, _ := http.Post(hts.URL, "application/json", buffer)

	march, _ := io.ReadAll(res.Body)
	macthEq, _ := json.Marshal(responseMock)

	assert.Equal(t, march, macthEq)

}

func TestTestValidateLoginWhenParamErr(t *testing.T) {
	controller := new(userControllerImpl)
	gin := gin.Default()
	gin.POST("/", controller.ValidateLogin)
	hts := httptest.NewServer(gin)

	buildIUserHandler = func(con *contextDto.ContextDto) userHandler.IUserHandler {
		return new(userhandlerMock)
	}
	responseManager = func(response *responseDto.ResponseDto) (int, interface{}) {
		return 200, responseMock
	}
	responseErrorManager = func(errDto *errorManagerDto.ErrorManagerDto) *responseDto.ResponseDto {
		return responseMock
	}
	logger = func(errIn, errOut string, status int, errExeption string) *errorManagerDto.ErrorManagerDto {
		return nil
	}

	buffer := new(bytes.Buffer)
	res, _ := http.Post(hts.URL, "application/json", buffer)

	march, _ := io.ReadAll(res.Body)
	marchEq, _ := json.Marshal(responseMock)
	assert.Equal(t, march, marchEq)
	assert.Equal(t, res.StatusCode, 500)
}

func TestCreateNewUser(t *testing.T) {
	controller := new(userControllerImpl)
	gin := gin.Default()
	gin.POST("/", controller.CreateNewUser)
	hts := httptest.NewServer(gin)

	buildIUserHandler = func(con *contextDto.ContextDto) userHandler.IUserHandler {
		return new(userhandlerMock)
	}
	responseManager = func(response *responseDto.ResponseDto) (int, interface{}) {
		return 200, responseMock
	}
	responseErrorManager = func(errDto *errorManagerDto.ErrorManagerDto) *responseDto.ResponseDto {
		return responseMock
	}
	logger = func(errIn, errOut string, status int, errExeption string) *errorManagerDto.ErrorManagerDto {
		return nil
	}

	buffer := new(bytes.Buffer)
	userData := new(dto.UserdataDto)
	json.NewEncoder(buffer).Encode(userData)
	res, _ := http.Post(hts.URL, "application/json", buffer)

	march, _ := io.ReadAll(res.Body)
	marchEq, _ := json.Marshal(responseMock)
	assert.Equal(t, march, marchEq)
	assert.Equal(t, res.StatusCode, 200)
}

func TestCreateNewUserWhenParamsErr(t *testing.T) {
	controller := new(userControllerImpl)
	gin := gin.Default()
	gin.POST("/", controller.CreateNewUser)
	hts := httptest.NewServer(gin)

	buildIUserHandler = func(con *contextDto.ContextDto) userHandler.IUserHandler {
		return new(userhandlerMock)
	}
	responseManager = func(response *responseDto.ResponseDto) (int, interface{}) {
		return 200, responseMock
	}
	responseErrorManager = func(errDto *errorManagerDto.ErrorManagerDto) *responseDto.ResponseDto {
		return responseMock
	}
	logger = func(errIn, errOut string, status int, errExeption string) *errorManagerDto.ErrorManagerDto {
		return nil
	}

	buffer := new(bytes.Buffer)
	res, _ := http.Post(hts.URL, "application/json", buffer)

	march, _ := io.ReadAll(res.Body)
	marchEq, _ := json.Marshal(responseMock)
	assert.Equal(t, march, marchEq)
	assert.Equal(t, res.StatusCode, 500)
}
