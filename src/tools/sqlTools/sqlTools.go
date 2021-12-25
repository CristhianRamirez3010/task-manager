package sqlTools

import (
	"fmt"
)

type SqlTool struct {
	fields    []string
	table     string
	selectStr string
	joins     []string
	where     string
	insert    string
	addQuery  string
}

func BuildSqlTools(table string, fields ...string) *SqlTool {
	tool := new(SqlTool)
	tool.table = table
	tool.fields = fields
	return tool
}

func (s *SqlTool) AddSelect() *SqlTool {
	return s.AddSelectWithRef("")
}

func (s *SqlTool) AddSelectWithRef(tableRefence string) *SqlTool {
	s.cleanAll()
	s.selectStr = fmt.Sprintf(" select %s from %s %s ",
		s.buildFields(tableRefence),
		s.table,
		tableRefence)
	return s
}

func (s *SqlTool) BuildQuery() string {
	query := ""
	if s.selectStr != "" {
		query = fmt.Sprintf("%s %s", query, s.selectStr)
		for join := range s.joins {
			fmt.Sprintf("%s %s", query, join)
		}
		query = fmt.Sprintf("%s %s;", query, s.where)
	} else if s.insert != "" {
		query = fmt.Sprintf("%s %s;", query, s.insert)
	}
	query = fmt.Sprintf("%s %s", query, s.addQuery)

	return query
}

func (s *SqlTool) AddInnerJoin(table string, tableRefence string, conditions string) *SqlTool {
	s.joins = append(s.joins, fmt.Sprintf("inner join %s %s on %s", table, tableRefence, conditions))
	return s
}

func (s *SqlTool) AddWhere(where string) *SqlTool {
	s.where = fmt.Sprintf(" where %s", where)
	return s
}

func (s *SqlTool) AddInsert() *SqlTool {
	s.cleanAll()
	values := ""
	for range s.fields {
		values = fmt.Sprintf("%s?,", values)
	}
	values = values[:len(values)-1]
	s.insert = fmt.Sprintf("insert into %s(%s) values(%s)", s.table, s.buildFields(""), values)
	return s
}

func (s *SqlTool) buildFields(tableRefence string) string {
	fields := ""
	if tableRefence != "" {
		tableRefence = fmt.Sprintf("%s.", tableRefence)
	}
	for _, field := range s.fields {
		fields = fmt.Sprintf("%s %s%s,", fields, tableRefence, field)
	}
	varFields := fields[:len(fields)-1]
	return varFields
}

func (s *SqlTool) cleanAll() {
	s.selectStr = ""
	s.joins = []string{}
	s.insert = ""
	s.where = ""
}
