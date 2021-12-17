package constants

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var constantsMock = new(Constants)

func TestBuildConstants(t *testing.T) {
	as := assert.New(t)
	res := BuildConstants()
	as.NotNil(res)
}

func TestGetPort(t *testing.T) {
	as := assert.New(t)
	port := ":9090"
	constantsMock.port = port
	as.Equal(constantsMock.GetPort(), port)
}
func TestGetMaxOpenDbConn(t *testing.T) {
	as := assert.New(t)
	maxOpenDbConn := "10"
	constantsMock.maxOpenDbConn = maxOpenDbConn
	conv, _ := strconv.Atoi(maxOpenDbConn)
	as.Equal(constantsMock.GetMaxOpenDbConn(), conv)
}

func TestGetMaxOpenDbConnIfEnvIsEmpty(t *testing.T) {
	as := assert.New(t)
	constantsMock.maxOpenDbConn = ""
	as.Equal(constantsMock.GetMaxOpenDbConn(), maxOpenDbConnDefault)
}

func TestGetMasIdleDbConn(t *testing.T) {
	as := assert.New(t)
	masIdleDbConn := "5ms"
	constantsMock.masIdleDbConn = masIdleDbConn
	conv, _ := time.ParseDuration(masIdleDbConn)
	as.Equal(constantsMock.GetMasIdleDbConn(), conv)
}

func TestGetMasIdleDbConnIdEnvIsEmpty(t *testing.T) {
	as := assert.New(t)
	constantsMock.masIdleDbConn = ""
	as.Equal(constantsMock.GetMasIdleDbConn(), time.Duration(masIdleDbConnDefault))
}
func TestGetMaxDbLifetime(t *testing.T) {
	as := assert.New(t)
	maxDbLifetime := "5m"
	constantsMock.maxDbLifetime = maxDbLifetime
	conv, _ := time.ParseDuration(maxDbLifetime)
	as.Equal(constantsMock.GetMaxDbLifetime(), conv)
}
func TestGetMaxDbLifetimeIdEnvIsEmpty(t *testing.T) {
	as := assert.New(t)
	constantsMock.maxDbLifetime = ""
	as.Equal(constantsMock.GetMaxDbLifetime(), time.Duration(maxDbLifetimeDefault))
}

func TestGetMysqlConnectionString(t *testing.T) {
	as := assert.New(t)
	mySQLConnectionString := "some connection string"
	constantsMock.mySQLConnectionString = mySQLConnectionString
	as.Equal(constantsMock.GetMysqlConnectionString(), mySQLConnectionString)
}
