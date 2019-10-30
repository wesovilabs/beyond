package internal

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"testing"
)

func TestAssignGoaContext2(t *testing.T) {
	assert := assert.New(t)
	stmt := AssignGoaContext(map[string]string{"context": "_xontext"})
	assert.Len(stmt.Rhs, 1)
	assert.Len(stmt.Lhs, 1)
	assert.Equal(varGoaContext, stmt.Lhs[0].(*ast.Ident).Name)
	r := stmt.Rhs[0].(*ast.CallExpr)
	assert.NotEmpty(r.Args)
}
