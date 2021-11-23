package impl

import (
	"database/sql"
	"fmt"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/connections"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/constants"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
)

const (
	errDefault = "Error with the repository"
)

type repositoryBase struct {
	Fields    []string
	Table     string
	Constants constants.Constants
}

func (r repositoryBase) loadConnection() (*sql.DB, *errorManagerDto.ErrorManagerDto) {
	return connections.BuildMySQLConnection(
		r.Constants.GetMysqlConnectionString(),
		r.Constants.GetMaxOpenDbConn(),
		r.Constants.GetMasIdleDbConn(),
		r.Constants.GetMaxDbLifetime(),
	).ConnectDBMysql()
}

func (r repositoryBase) selectAll(where string) string {
	if where != "" {
		where = fmt.Sprintf(" where %s", where)
	}
	selectStr := ""
	for _, field := range r.Fields {
		selectStr = fmt.Sprintf("%s %s,", selectStr, field)
	}
	selectStr = selectStr[:len(selectStr)-1]
	return fmt.Sprintf("select %s from %s %s;",
		selectStr,
		r.Table,
		where)
}
