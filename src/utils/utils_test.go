package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildStrFromArray(t *testing.T) {
	as := assert.New(t)
	res := BuildStrFromArray([]string{
		" Test ",
		" Array ",
	})
	as.NotNil(res, "BuildStrFromArray should not return nil")
	as.Equal(*res, " Test  Array ", "Error with the test")
}

func TestLogger(t *testing.T) {
	as := assert.New(t)
	result := Logger("Test", "Test", 500, "Test")
	as.NotNil(result)
	as.NotEqual(result.Message, "", "Message should not be empty")
	as.NotEqual(result.Status, 0, "Status should not be empty")
}
