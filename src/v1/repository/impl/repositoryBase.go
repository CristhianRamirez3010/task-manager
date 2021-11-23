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
	Constants *constants.Constants
}

func (r repositoryBase) loadConnection() (*sql.DB, *errorManagerDto.ErrorManagerDto) {
	return connections.BuildMySQLConnection(
		r.Constants.GetMysqlConnectionString(),
		r.Constants.GetMaxOpenDbConn(),
		r.Constants.GetMasIdleDbConn(),
		r.Constants.GetMaxDbLifetime(),
	).ConnectDBMysql()
}

func (r *repositoryBase) selectAll(where string) string {
	if where != "" {
		where = fmt.Sprintf(" where %s", where)
	}

	return fmt.Sprintf("select %s from %s %s;",
		r.buildFields(),
		r.Table,
		where)
}

func (r *repositoryBase) insertAll() string {
	values := ""

	for range r.Fields {
		values = fmt.Sprintf("%s?,", values)
	}
	values = values[:len(values)-1]

	return fmt.Sprintf("insert into %s(%s) values(%s)", r.Table, r.buildFields(), values)
}

func (r *repositoryBase) buildFields() string {
	fields := ""

	for _, field := range r.Fields {
		fields = fmt.Sprintf("%s %s,", fields, field)
	}
	return fields[:len(fields)-1]
}
