package advice

import (
	assert2 "github.com/stretchr/testify/assert"
	"go/ast"
	"testing"
)

func Test_unsupportedType(t *testing.T) {
	if res := unsupportedType("sample"); res != "" {
		t.Fatalf("unexpected value")
	}
}

func Test_pkgPathForType(t *testing.T) {
	assert := assert2.New(t)
	res := pkgPathForType("myname", []*ast.ImportSpec{
		{
			Name: ast.NewIdent("other"),
			Path: &ast.BasicLit{
				Value: "\"project/pkg/myname\"",
			},
		},
	})
	assert.Equal("project/pkg/myname", res)
}
