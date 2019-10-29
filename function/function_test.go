package function

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestFunction(t *testing.T) {
	assert := assert.New(t)
	file, err := parser.ParseFile(&token.FileSet{}, "testdata/package/sample.go", nil, parser.ParseComments)
	assert.Nil(err)
	assert.NotNil(file)
	functions := &Functions{}
	searchFunctions("testdata/package", file, functions)
	assert.Equal(2, len(functions.List()))
	function := functions.List()[0]
	assert.Equal("testdata/package", function.Pkg())
	assert.Equal("sample", function.Name())
	assert.Equal("_package.sample(Person)string", function.Path())
	function = functions.List()[1]
	assert.Equal("testdata/package", function.Pkg())
	assert.Equal("sample2", function.Name())
	assert.Equal("_package.sample2(*other.Other)func(map[string]interface{})", function.Path())
	assert.Equal(1, len(function.ImportSpecs()))
	assert.Equal("\"github.com/wesovilabs/goa/function/testdata/other\"", function.ImportSpecs()[0].Path.Value)
	stmt := &ast.EmptyStmt{}
	function.AddStatementsAtBegin([]ast.Stmt{stmt})
	assert.Equal(1, len(function.ParamsList()))
	assert.Equal(1, len(function.ImportSpecs()))
	assert.Equal(3, len(function.FileDecls()))
	assert.Equal(file, function.Parent())

}
