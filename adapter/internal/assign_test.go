package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/token"
	"testing"
)

type callMethodAndAssignRequest struct {
	recv                  *ast.FieldList
	currentPkg, pkg, name string
	params, results       []*FieldDef
}

func createCallMethodAndAssignRequest() *callMethodAndAssignRequest {
	return &callMethodAndAssignRequest{
		recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{NewIdentObj("param")},
				},
			},
		},
		currentPkg: "storage",
		pkg:        "mypkg",
		name:       "myfunction",
		params: []*FieldDef{
			{
				Name: "param1",
				Kind: &ast.Ident{},
			},
			{
				Name: "param2",
				Kind: &ast.Ident{},
			},
		},
		results: []*FieldDef{
			{
				Name: "output",
				Kind: &ast.Ident{},
			},
		},
	}
}

func checkCallExpr(assert *assert.Assertions, req *callMethodAndAssignRequest, right *ast.CallExpr) {
	fmt.Printf("%#v\n", right.Fun)
	sel, ok := right.Fun.(*ast.SelectorExpr)
	if !ok {
		assert.Fail("Unexpected type for statement")
	}
	fmt.Printf("%#v\n", sel.X)
	selX, ok := sel.X.(*ast.Ident)
	if !ok {
		assert.Fail("Unexpected type for statement")
	}
	assert.Equal(req.recv.List[0].Names[0].Name, selX.Name)
	fmt.Printf("%#v\n", sel.Sel)
	assert.Equal(req.name, sel.Sel.Name)

	fmt.Printf("%#v\n", right.Args)
	assert.Len(right.Args, 2)

	fmt.Printf("%#v\n", right.Args[0])
	param1, ok := right.Args[0].(*ast.Ident)
	if !ok {
		assert.Fail("Unexpected type for statement")
	}
	assert.Equal("param0", param1.Name)
	fmt.Printf("%#v\n", right.Args[1])
	param2, ok := right.Args[1].(*ast.Ident)
	if !ok {
		assert.Fail("Unexpected type for statement")
	}
	assert.Equal("param1", param2.Name)
}

func Test_CallMethodAndAssign(t *testing.T) {
	assert := assert.New(t)
	req := createCallMethodAndAssignRequest()
	stmt := CallMethodAndAssign(req.recv, req.currentPkg, req.pkg, req.name, req.params, req.results)
	fmt.Printf("%#v\n", stmt)
	if assign, ok := stmt.(*ast.AssignStmt); ok {
		assert.Len(assign.Lhs, 1)
		assert.Len(assign.Rhs, 1)
		assert.Equal(token.DEFINE, assign.Tok)
		fmt.Printf("%#v\n", assign.Lhs[0])
		left, ok := assign.Lhs[0].(*ast.Ident)
		if !ok {
			assert.Fail("Unexpected type for statement")
		}

		assert.Equal(req.results[0].Name, left.Name)
		fmt.Printf("%#v\n", assign.Rhs[0])
		right, ok := assign.Rhs[0].(*ast.CallExpr)
		checkCallExpr(assert, req, right)
		return
		if !ok {
			assert.Fail("Unexpected type for statement")
		}

		return
	}
	t.Fatal("Unexpected statement type")
}

func Test_CallMethodAndAssignNoResults(t *testing.T) {
	assert := assert.New(t)
	req := createCallMethodAndAssignRequest()
	stmt := CallMethodAndAssign(req.recv, req.currentPkg, req.pkg, req.name, req.params, nil)
	fmt.Printf("%#v\n", stmt)
	if exprStmt, ok := stmt.(*ast.ExprStmt); ok {
		fmt.Printf("%#v\n", exprStmt.X)
		if callExpr, ok := exprStmt.X.(*ast.CallExpr); ok {
			checkCallExpr(assert, req, callExpr)
			return
		}
	}
	t.Fatal("Unexpected statement type")
}
