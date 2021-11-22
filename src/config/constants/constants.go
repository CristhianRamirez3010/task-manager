package constants

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Constants struct {
	port                  string
	mySQLConnectionString string
	maxOpenDbConn         string
	masIdleDbConn         string
	maxDbLifetime         string
}

func BuildConstants() Constants {
	return Constants{
		port:                  os.Getenv("PORT"),
		mySQLConnectionString: os.Getenv("MYSQLCONNECTIONSTRING"),
		maxOpenDbConn:         os.Getenv("MAXOPENDBCONN"),
		masIdleDbConn:         os.Getenv("MASIDLEDBCONN"),
		maxDbLifetime:         os.Getenv("MAXDBLIFETIME"),
	}
}

func (c Constants) GetPort() string {
	return c.port
}
func (c Constants) GetMaxOpenDbConn() int {
	maxOpenDbConn, err := strconv.Atoi(c.maxOpenDbConn)
	if err != nil {
		log.Print("not maxOpenDbConn from env (default: 10)")
		maxOpenDbConn = 10
	}
	return maxOpenDbConn
}
func (c Constants) GetMasIdleDbConn() time.Duration {
	masIdleDbConn, err := time.ParseDuration(c.masIdleDbConn)
	if err != nil {
		log.Print("not masIdleDbConn from env (default: 5ms)")
		masIdleDbConn = 5
	}
	return masIdleDbConn
}
func (c Constants) GetMaxDbLifetime() time.Duration {
	maxDbLifetime, err := time.ParseDuration(c.maxDbLifetime)
	if err != nil {
		log.Print("not maxDbLifetime from env (default: 5m)")
		maxDbLifetime = 5 * time.Minute
	}
	return maxDbLifetime

}
func (c Constants) GetMysqlConnectionString() string {
	return c.mySQLConnectionString
}
