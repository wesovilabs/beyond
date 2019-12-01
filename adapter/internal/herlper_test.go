package internal

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"testing"
)

func Test_astToExpressionEval(t *testing.T) {
	assert := assert.New(t)
	cases:=[]struct{
		expr ast.Expr
		path string
	}{
		{
			expr:&ast.MapType{
				Key:   NewIdentObj("string"),
				Value: NewIdentObj("*int"),
			},
			path: "map[string]*int",
		},
		{
			expr:&ast.Ident{
				Name:`"hello"`,
			},
			path:`"hello"`,
		},
		{
			expr:&ast.Ident{
				Name:`"123.2"`,
			},
			path:`"123.2"`,
		},
		{
			expr:&ast.InterfaceType{},
			path:"interface{}",
		},
		{
			expr:&ast.StarExpr{
				X:NewIdentObj("variable"),
			},
			path:"*variable",
		},
		{
			expr:&ast.StructType{},
			path:"struct{}",
		},
		{
			expr:&ast.ArrayType{
				Elt:&ast.InterfaceType{},
			},
			path:"[]interface{}",
		},
		{
			expr:&ast.Ellipsis{
				Elt:&ast.InterfaceType{},
			},
			path:"[]interface{}",
		},
		{
			expr:&ast.SelectorExpr{
				X:NewIdentObj("x"),
				Sel:NewIdentObj("y"),
			},
			path:"x.y",
		},
	}
	for _,c:=range cases{
		assert.Equal(c.path,astToExpressionEval(c.expr))
	}

}
