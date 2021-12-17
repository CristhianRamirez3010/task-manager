package connections

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/stretchr/testify/assert"
)

func TestBuildMySQLConnection(t *testing.T) {
	as := assert.New(t)
	res := BuildMySQLConnection("", 0, 0, 0)
	as.NotNil(res, "BuildMySQLConnection should not be null")
}

func TestConnectDBMysql(t *testing.T) {
	as := assert.New(t)
	mysql := new(MySQLConnection)
	mysql.connectionsString = "test"
	sqlMock := new(sql.DB)
	openConnection = func(driverName, dataSourceName string) (*sql.DB, error) {
		return sqlMock, nil
	}
	logger = func(errIn, errOut string, status int, errExeption string) *errorManagerDto.ErrorManagerDto {
		return new(errorManagerDto.ErrorManagerDto)
	}
	result, err := mysql.ConnectDBMysql()
	as.Nil(err, "ConnectDBMysql error should be null")
	as.NotNil(result, "ConnectDBMysql object should not becls null")

}

func TestConnectDBMysqlIfConnectionsStringIsNull(t *testing.T) {
	as := assert.New(t)
	mysql := new(MySQLConnection)
	sqlMock := new(sql.DB)
	openConnection = func(driverName, dataSourceName string) (*sql.DB, error) {
		return sqlMock, nil
	}
	logger = func(errIn, errOut string, status int, errExeption string) *errorManagerDto.ErrorManagerDto {
		return new(errorManagerDto.ErrorManagerDto)
	}
	_, err := mysql.ConnectDBMysql()
	as.NotNil(err, "ConnectDBMysql error should not be null")
}

func TestConnectDBMysqlIfOpenConnectionIsError(t *testing.T) {
	as := assert.New(t)
	mysql := new(MySQLConnection)
	mysql.connectionsString = "test"
	openConnection = func(driverName, dataSourceName string) (*sql.DB, error) {
		return nil, fmt.Errorf("err")
	}
	logger = func(errIn, errOut string, status int, errExeption string) *errorManagerDto.ErrorManagerDto {
		return new(errorManagerDto.ErrorManagerDto)
	}
	_, err := mysql.ConnectDBMysql()
	as.NotNil(err, "ConnectDBMysql error should not be null")
}
