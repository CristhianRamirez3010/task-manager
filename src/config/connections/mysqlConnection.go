package connections

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	_ "github.com/go-sql-driver/mysql"
)

const (
	errDefault = "Error with the connection"
)

type MySQLConnection struct {
	connectionsString string
	maxOpenDbConn     int
	masIdleDbConn     time.Duration
	maxDbLifetime     time.Duration
}

var (
	openConnection = sql.Open
	logger         = utils.Logger
)

func BuildMySQLConnection(connectionsString string, maxOpenDbConn int, masIdleDbConn time.Duration, maxDbLifetime time.Duration) *MySQLConnection {
	return &MySQLConnection{
		connectionsString: connectionsString,
		masIdleDbConn:     masIdleDbConn,
		maxDbLifetime:     maxDbLifetime * time.Minute,
		maxOpenDbConn:     maxOpenDbConn,
	}
}

func (m *MySQLConnection) ConnectDBMysql() (*sql.DB, *errorManagerDto.ErrorManagerDto) {
	if m.connectionsString == "" {
		return nil, logger("Not connection string in .env", errDefault, http.StatusInternalServerError, "")
	}
	db, err := openConnection("mysql", m.connectionsString)
	if err != nil {
		return nil, logger("Connection with provider is not working", errDefault, http.StatusInternalServerError, err.Error())
	}
	db.SetMaxOpenConns(m.maxOpenDbConn)
	db.SetConnMaxIdleTime(m.masIdleDbConn)
	db.SetConnMaxLifetime(m.maxDbLifetime)
	return db, nil
}
