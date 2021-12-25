package projectController

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handlers/projectHandler"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/proProjectModel"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type projectHandlerMock struct {
	projectHandler.IProjectHandler
	mock.Mock
}

func (p *projectHandlerMock) GetProject() *responseDto.ResponseDto {
	return new(responseDto.ResponseDto)
}
func (p *projectHandlerMock) NewProject(projectModel *proProjectModel.ProProjectModel) *responseDto.ResponseDto {
	return new(responseDto.ResponseDto)
}

func TestBuildIProjectControllerNotNil(t *testing.T) {
	obj := BuildIProjectController()
	assert.NotNil(t, obj)
}

func TestGetProjectsReturnMock(t *testing.T) {
	responseMock := new(responseDto.ResponseDto)
	buildIProjectHandler = func(context *contextDto.ContextDto) projectHandler.IProjectHandler {
		return new(projectHandlerMock)
	}
	responseManager = func(response *responseDto.ResponseDto) (int, interface{}) {
		return 200, responseMock
	}

	controller := new(projectControllerImpl)
	gin := gin.Default()
	gin.GET("/", controller.GetProjects)
	hts := httptest.NewServer(gin)
	res, _ := http.Get(hts.URL)

	march, _ := io.ReadAll(res.Body)
	res.Body.Close()
	marchEq, _ := json.Marshal(responseMock)

	assert.Equal(t, res.StatusCode, 200)
	assert.Equal(t, march, marchEq)
}

func TestNewProjectReturnMock(t *testing.T) {
	responseMock := new(responseDto.ResponseDto)
	controller := new(projectControllerImpl)
	buildIProjectHandler = func(context *contextDto.ContextDto) projectHandler.IProjectHandler {
		return new(projectHandlerMock)
	}
	responseManager = func(response *responseDto.ResponseDto) (int, interface{}) {
		return 200, responseMock
	}
	gin := gin.Default()
	gin.POST("/", controller.NewProject)
	hts := httptest.NewServer(gin)
	userMock := proProjectModel.ProProjectModel{
		Name:        "test",
		Description: "test",
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(userMock)
	res, _ := http.Post(hts.URL, "application/json", &buf)

	march, _ := io.ReadAll(res.Body)
	res.Body.Close()
	marchEq, err := json.Marshal(responseMock)

	assert.Equal(t, err, nil)
	assert.Equal(t, res.StatusCode, 200)
	assert.Equal(t, march, marchEq)
}

func TestNewTestNewProjectReturnErrWithStruct(t *testing.T) {
	responseMock := new(responseDto.ResponseDto)
	controller := new(projectControllerImpl)
	buildIProjectHandler = func(context *contextDto.ContextDto) projectHandler.IProjectHandler {
		return new(projectHandlerMock)
	}
	responseManager = func(response *responseDto.ResponseDto) (int, interface{}) {
		return 200, responseMock
	}
	gin := gin.Default()
	gin.POST("/", controller.NewProject)
	hts := httptest.NewServer(gin)
	var buf bytes.Buffer
	res, _ := http.Post(hts.URL, "application/json", &buf)
	assert.Equal(t, res.StatusCode, 500)
}
