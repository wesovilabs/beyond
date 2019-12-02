package internal

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"testing"
)

func Test_Results(t *testing.T) {
	assert := assert.New(t)

	fielDefs := Results([]*ast.Field{
		{
			Names: []*ast.Ident{
				NewIdentObj("err"),
			},
			Type: NewIdentObj("error"),
		},
		{
			Names: []*ast.Ident{
				NewIdentObj("name"),
				NewIdentObj("country"),
			},
			Type: NewIdentObj("string"),
		},
	})
	assert.Len(fielDefs, 3)
	assert.Equal("result0", fielDefs[0].Name)
	assert.Equal("error", fielDefs[0].Kind.(*ast.Ident).String())
	assert.Equal("result1", fielDefs[1].Name)
	assert.Equal("string", fielDefs[1].Kind.(*ast.Ident).String())
	assert.Equal("result2", fielDefs[2].Name)
	assert.Equal("string", fielDefs[2].Kind.(*ast.Ident).String())
}
