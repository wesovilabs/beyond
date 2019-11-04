package function

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	goaParser "github.com/wesovilabs/goa/parser"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

const project = "github.com/wesovilabs/goa/function/testdata"

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
	assert.Equal("testdata/package.sample(Person)string", function.Path())
	function = functions.List()[1]
	assert.Equal("testdata/package", function.Pkg())
	assert.Equal("sample2", function.Name())
	assert.Equal("testdata/package.sample2(*github.com/wesovilabs/goa/function/testdata/other.Other)func(map[string]interface{})", function.Path())
	assert.Equal(1, len(function.ImportSpecs()))
	assert.Equal("\"github.com/wesovilabs/goa/function/testdata/other\"", function.ImportSpecs()[0].Path.Value)
	stmt := &ast.EmptyStmt{}
	function.AddStatementsAtBegin([]ast.Stmt{stmt})
	assert.Equal(1, len(function.ParamsList()))
	assert.Equal(1, len(function.ImportSpecs()))
	assert.Equal(3, len(function.FileDecls()))
	assert.Equal(file, function.Parent())

}

func TestFunction_AddImportSpecs(t *testing.T) {
	project := "testdata"
	assert := assert.New(t)
	packages := goaParser.
		New("testdata", project).
		Parse("")
	functions := GetFunctions(packages)
	assert.NotNil(functions)
	assert.Len(functions.List(), 23)
	function := functions.List()[0]
	currentImports := len(function.ImportSpecs())
	function.AddImportSpec(&ast.ImportSpec{
		Name: ast.NewIdent("test"),
	},
	)
	assert.Equal(currentImports+1, len(function.ImportSpecs()))
}

func TestFunction_RenameToInternal(t *testing.T) {
	assert := assert.New(t)
	packages := goaParser.
		New("testdata", project).
		Parse("")
	functions := GetFunctions(packages)
	assert.NotNil(functions)
	assert.Len(functions.List(), 25)
	function := functions.List()[0]
	name := function.Name()
	function.RenameToInternal()
	assert.Equal(fmt.Sprintf("%sInternal", name), function.Name())
}
