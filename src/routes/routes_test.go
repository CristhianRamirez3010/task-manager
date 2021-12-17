package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadRoutes(t *testing.T) {
	as := assert.New(t)
	result := LoadRoutes()
	as.NotNil(result)
}
