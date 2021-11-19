package connections

import (
	"database/sql"
	"log"
	"time"
)

type MySQLConnection struct {
	connectionsString string
	maxOpenDbConn     int
	masIdleDbConn     time.Duration
	maxDbLifetime     time.Duration
}

func BuildMySQLConnection(connectionsString string, maxOpenDbConn int, masIdleDbConn time.Duration, maxDbLifetime time.Duration) MySQLConnection {
	return MySQLConnection{
		connectionsString: connectionsString,
		masIdleDbConn:     masIdleDbConn,
		maxDbLifetime:     maxDbLifetime * time.Minute,
		maxOpenDbConn:     maxOpenDbConn,
	}
}

func (m MySQLConnection) ConnectDBMysql() (*sql.DB, error) {
	if m.connectionsString == "" {
		log.Fatal("not connection string")
		return nil, nil
	}
	db, err := sql.Open("mysql", m.connectionsString)
	if err != nil {
		log.Fatal("error with the connection", err.Error())
		return nil, err
	}
	db.SetMaxOpenConns(m.maxOpenDbConn)
	db.SetConnMaxIdleTime(m.masIdleDbConn)
	db.SetConnMaxLifetime(m.maxDbLifetime)
	return db, nil
}
