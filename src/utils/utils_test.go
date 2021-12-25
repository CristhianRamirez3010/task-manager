package utils

import (
	"testing"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/stretchr/testify/assert"
)

func TestBuildStrFromArray(t *testing.T) {
	as := assert.New(t)
	res := BuildStrFromArray(
		" Test ",
		" Array ",
	)
	as.NotNil(res, "BuildStrFromArray should not return nil")
	as.Equal(res, " Test  Array ", "Error with the test")
}

func TestLogger(t *testing.T) {
	as := assert.New(t)
	result := Logger("Test", "Test", 500, "Test")
	as.NotNil(result)
	as.NotEqual(result.Message, "", "Message should not be empty")
	as.NotEqual(result.Status, 0, "Status should not be empty")
}

func TestResponseManager(t *testing.T) {
	code, _ := ResponseManager(new(responseDto.ResponseDto))
	assert.Equal(t, code, 200)
}

func TestResponseManagerIsErr(t *testing.T) {
	resp := new(responseDto.ResponseDto)
	resp.Error = errorManagerDto.ErrorManagerDto{
		Status:  500,
		Message: "test",
	}
	code, result := ResponseManager(resp)
	assert.Equal(t, code, 500)
	assert.Equal(t, result, resp)
}
