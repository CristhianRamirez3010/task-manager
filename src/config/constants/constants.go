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

const (
	maxOpenDbConnDefault = 10
	masIdleDbConnDefault = 5
	maxDbLifetimeDefault = 5 * time.Minute
)

var (
	castToInt     = strconv.Atoi
	parseDuration = time.ParseDuration
)

func BuildConstants() *Constants {
	return &Constants{
		port:                  os.Getenv("PORT"),
		mySQLConnectionString: os.Getenv("MYSQLCONNECTIONSTRING"),
		maxOpenDbConn:         os.Getenv("MAXOPENDBCONN"),
		masIdleDbConn:         os.Getenv("MASIDLEDBCONN"),
		maxDbLifetime:         os.Getenv("MAXDBLIFETIME"),
	}
}

func (c *Constants) GetPort() string {
	return c.port
}
func (c *Constants) GetMaxOpenDbConn() int {
	maxOpenDbConn, err := castToInt(c.maxOpenDbConn)
	if err != nil {
		log.Print("not maxOpenDbConn from env (default: 10)")
		maxOpenDbConn = maxOpenDbConnDefault
	}
	return maxOpenDbConn
}
func (c *Constants) GetMasIdleDbConn() time.Duration {
	masIdleDbConn, err := parseDuration(c.masIdleDbConn)
	if err != nil {
		log.Print("not masIdleDbConn from env (default: 5ms)")
		masIdleDbConn = time.Duration(masIdleDbConnDefault)
	}
	return masIdleDbConn
}
func (c *Constants) GetMaxDbLifetime() time.Duration {
	maxDbLifetime, err := parseDuration(c.maxDbLifetime)
	if err != nil {
		log.Print("not maxDbLifetime from env (default: 5m)")
		maxDbLifetime = time.Duration(maxDbLifetimeDefault)
	}
	return maxDbLifetime

}
func (c *Constants) GetMysqlConnectionString() string {
	return c.mySQLConnectionString
}
