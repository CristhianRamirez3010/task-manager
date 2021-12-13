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

func (r *repositoryBase) addSelect() *string {
	return r.addSelectWithRef("")
}

func (r *repositoryBase) addInnerJoin(table string, tableRefence string, conditions string) *string {
	join := fmt.Sprintf("inner join %s %s on %s", table, tableRefence, conditions)
	return &join
}

func (r *repositoryBase) addWhere(where string) *string {
	varWhere := fmt.Sprintf(" where %s", where)
	return &varWhere
}

func (r *repositoryBase) addSelectWithRef(tableRefence string) *string {
	varSelect := fmt.Sprintf(" select %s from %s %s ",
		*r.buildFields(tableRefence),
		r.Table,
		tableRefence)
	return &varSelect
}

func (r *repositoryBase) addInsert() *string {
	values := ""

	for range r.Fields {
		values = fmt.Sprintf("%s?,", values)
	}
	values = values[:len(values)-1]
	varSelect := fmt.Sprintf("insert into %s(%s) values(%s)", r.Table, *r.buildFields(""), values)
	return &varSelect
}

func (r *repositoryBase) buildFields(tableRefence string) *string {
	fields := ""
	if tableRefence != "" {
		tableRefence = fmt.Sprintf("%s.", tableRefence)
	}
	for _, field := range r.Fields {
		fields = fmt.Sprintf("%s %s%s,", fields, tableRefence, field)
	}
	varFields := fields[:len(fields)-1]
	return &varFields
}

func (r *repositoryBase) buildQuery(segments []*string) *string {
	query := ""
	for _, segment := range segments {
		query = fmt.Sprintf("%s %s", query, *segment)
	}
	return &query
}

func (r *repositoryBase) addMySqlLastInsertId() *string {
	query := "select LAST_INSERT_ID()"
	return &query
}
