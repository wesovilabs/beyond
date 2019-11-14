package internal

import (
	"go/ast"
)

// FuncDecl create a new FuncDecl
func FuncDecl(name string, params, results []*ast.Field, stmts []ast.Stmt) *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: NewIdentObj(name),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: params,
			},
			Results: &ast.FieldList{
				List: results,
			},
		},
		Body: &ast.BlockStmt{
			List: stmts,
		},
	}
}
