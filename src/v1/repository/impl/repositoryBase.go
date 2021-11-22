package impl

import (
	"database/sql"
	"fmt"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/connections"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/constants"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
)

type RepositoryBase struct {
	fields    []string
	table     string
	Constants constants.Constants
}

func (r RepositoryBase) loadConnection() (*sql.DB, *errorManagerDto.ErrorManagerDto) {
	return connections.BuildMySQLConnection(
		r.Constants.GetMysqlConnectionString(),
		r.Constants.GetMaxOpenDbConn(),
		r.Constants.GetMasIdleDbConn(),
		r.Constants.GetMaxDbLifetime(),
	).ConnectDBMysql()
}

func (r RepositoryBase) selectAll(where string) string {
	if where != "" {
		where = fmt.Sprintf(" where %s", where)
	}
	selectStr := "select "
	for field := range r.fields {
		selectStr = fmt.Sprintf("%s %s,", selectStr, field)
	}
	return fmt.Sprintf("select %s from %s %s;",
		selectStr,
		r.table,
		where)
}
