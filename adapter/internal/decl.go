package internal

import (
	"fmt"
	"go/ast"
)

// FuncDecl create a new FuncDecl
func FuncDecl(name string, params, results []*ast.Field, stmts []ast.Stmt) *ast.FuncDecl {
	currentParams := make([]*ast.Field, 0)
	paramIndex := 0

	for index := range params {
		param := *params[index]
		newParam := &ast.Field{
			Names:   make([]*ast.Ident, len(param.Names)),
			Doc:     param.Doc,
			Type:    param.Type,
			Tag:     param.Tag,
			Comment: param.Comment,
		}

		for pIndex := range param.Names {
			paramName := fmt.Sprintf("param%v", paramIndex)
			newParam.Names[pIndex] = NewIdent(paramName)
			paramIndex++
		}

		currentParams = append(currentParams, newParam)
	}

	return &ast.FuncDecl{
		Name: NewIdentObj(name),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: currentParams,
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
