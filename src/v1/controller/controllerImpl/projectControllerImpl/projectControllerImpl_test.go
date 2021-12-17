package projectControllerImpl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ProjectControllerImplMock = new(ProjectControllerImpl)
)

func TestBuildProjectControllerImpl(t *testing.T) {
	as := assert.New(t)
	res := BuildProjectControllerImpl()
	as.NotNil(res)
}

func TestGetProjects(t *testing.T) {
	//
}

func TestNewProject(t *testing.T) {
	//
}
