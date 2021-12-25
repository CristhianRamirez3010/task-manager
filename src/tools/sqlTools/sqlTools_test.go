package sqlTools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildSqlTools(t *testing.T) {
	res := BuildSqlTools("", "")
	assert.NotNil(t, res)
}

func TestAddSelect(t *testing.T) {
	toolMock := new(SqlTool)
	toolMock.table = "maiu"
	toolMock.fields = []string{"test1", "test2"}
	toolMock.AddSelect()
	assert.NotEmpty(t, toolMock.selectStr)

}

func TestAddSelectWithRef(t *testing.T) {
	toolMock := new(SqlTool)
	toolMock.table = "maiu"
	toolMock.fields = []string{"test1", "test2"}
	toolMock.AddSelectWithRef("maiu")
	assert.NotEmpty(t, toolMock.selectStr)

}

func TestAddInnerJoin(t *testing.T) {
	toolMock := new(SqlTool)
	toolMock.AddInnerJoin("", "", "")
	assert.NotEmpty(t, toolMock.joins)
}

func TestAddWhere(t *testing.T) {
	toolMock := new(SqlTool)
	toolMock.AddWhere("")
	assert.NotEmpty(t, toolMock.where)
}

func TestAddInsert(t *testing.T) {
	toolMock := new(SqlTool)
	toolMock.table = "maiu"
	toolMock.fields = []string{"test1", "test2"}
	toolMock.AddInsert()
	assert.NotEmpty(t, toolMock.insert)
}

func TestBuildQueryInsert(t *testing.T) {
	toolMock := new(SqlTool)
	toolMock.table = "maiu"
	toolMock.fields = []string{"test1", "test2"}
	query := toolMock.AddInsert().BuildQuery()
	assert.NotEmpty(t, query)
}

func TestBuildQuerySelect(t *testing.T) {
	toolMock := new(SqlTool)
	toolMock.table = "maiu"
	toolMock.fields = []string{"test1", "test2"}
	toolMock.joins = []string{"test1", "test2"}
	query := toolMock.AddSelect().BuildQuery()
	assert.NotEmpty(t, query)
}
