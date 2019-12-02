package internal

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"testing"
)

func Test_CallCreateAspect(t *testing.T) {
	assert := assert.New(t)
	expr := CallCreateAspect("mypkg", "myfunc")
	assert.NotNil(expr)
	if ident, ok := expr.(*ast.Ident); ok {
		assert.Equal("mypkg.myfunc", ident.Name)
	} else {
		assert.Fail("Unexpected type")
	}

	expr = CallCreateAspect("", "myfunc")
	assert.NotNil(expr)
	if ident, ok := expr.(*ast.Ident); ok {
		assert.Equal("myfunc", ident.Name)
	} else {
		assert.Fail("Unexpected type")
	}

}

func Test_prepareArgs(t *testing.T) {
	assert := assert.New(t)
	res := prepareArgs([]*FieldDef{
		{
			Name: "param",
			Kind: &ast.Ellipsis{},
		},
	}, true)
	assert.Len(res, 1)
	assert.Equal("param...", res[0].(*ast.Ident).Name)
}
