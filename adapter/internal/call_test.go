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
