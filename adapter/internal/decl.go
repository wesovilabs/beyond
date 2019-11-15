package internal

import (
	"fmt"
	"go/ast"
)

// FuncDecl create a new FuncDecl
func FuncDecl(name string, params, results []*ast.Field, stmts []ast.Stmt) *ast.FuncDecl {
	currentParams := make([]*ast.Field, len(params))
	annonymousCounter := 0

	for index := range params {
		param := *params[index]
		paramName := param.Names[0].Name

		if paramName == "_" {
			paramName = fmt.Sprintf("annonymous%v", annonymousCounter)
			annonymousCounter++
		}

		param.Names[0].Name = paramName
		currentParams[index] = &param
	}

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
