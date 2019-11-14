package adapter

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/joinpoint"
	"github.com/wesovilabs/goa/parser"
	"go/ast"
	"testing"
)

const project = "github.com/wesovilabs/goa/adapter/testdata"

func Test_ensureImports(t *testing.T) {
	assert := assert.New(t)
	packages := parser.
		New("testdata", project).
		Parse("")
	functions := joinpoint.GetJoinPoints("", packages)
	function := functions.List()[0]
	imports := function.ImportSpecs()
	ensureImports(map[string]string{}, map[string]string{}, function)
	assert.Len(function.ImportSpecs(), len(imports))
	ensureImports(map[string]string{}, map[string]string{"pkg/test": "test"}, function)
	assert.Len(function.ImportSpecs(), len(imports)+1)
	ensureImports(map[string]string{}, map[string]string{"pkg/test": "test"}, function)
	assert.Len(function.ImportSpecs(), len(imports)+2)
	ensureImports(map[string]string{"pkg/test": "test"}, map[string]string{"pkg/test": "test"}, function)
	assert.Len(function.ImportSpecs(), len(imports)+2)
}

func Test_findImportName(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("test",
		findImportName(map[string]string{}, "test", "pkg/test"),
	)
	assert.Equal("test",
		findImportName(map[string]string{"pkg/test": "test"}, "test", "pkg/test"),
	)
	assert.Equal("test2",
		findImportName(map[string]string{"pkg/test": "test2"}, "test", "pkg/test"),
	)
	assert.Equal("_test",
		findImportName(map[string]string{"pkg2/test": "test"}, "test", "pkg/test"),
	)

	assert.Equal("test",
		findImportName(map[string]string{"pkg/test": "test"}, "", "pkg/test"),
	)
	assert.Equal("test2",
		findImportName(map[string]string{"pkg/test": "test2"}, "", "pkg/test"),
	)
	assert.Equal("_test",
		findImportName(map[string]string{"pkg2/test": "test"}, "", "pkg/test"),
	)
}

func Test_updateImportSpect(t *testing.T) {
	assert := assert.New(t)
	packages := parser.
		New("testdata", project).
		Parse("")
	functions := joinpoint.GetJoinPoints("", packages)
	function := functions.List()[0]
	currentImporSpecs := len(function.ImportSpecs())
	updateImportSpec(function, []ast.Spec{
		&ast.ImportSpec{
			Name: ast.NewIdent("test"),
		},
	})
	assert.Len(function.ImportSpecs(), currentImporSpecs+1)
}
