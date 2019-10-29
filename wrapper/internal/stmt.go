package internal

import (
	"fmt"
	"go/ast"
	"go/token"
)

func AssignValuesFromContextIn(fields []*FieldDef) []ast.Stmt {
	stmts := make([]ast.Stmt, len(fields))
	for index, f := range fields {
		stmts[index] = getFromContext(SelectorInGet, f)
	}
	return stmts
}

func AssignValuesFromContextOut(fields []*FieldDef) []ast.Stmt {
	stmts := make([]ast.Stmt, len(fields))
	for index, field := range fields {
		stmts[index] = getFromContext(SelectorOutGet, field)
	}
	return stmts
}

func getFromContext(selector *ast.SelectorExpr, field *FieldDef) *ast.AssignStmt {
	return &ast.AssignStmt{
		Tok: token.ASSIGN,
		Lhs: []ast.Expr{
			NewIdentObjVar(field.name),
		},
		Rhs: []ast.Expr{
			&ast.TypeAssertExpr{
				X: &ast.CallExpr{
					Fun: selector,
					Args: []ast.Expr{
						&ast.BasicLit{
							Kind:  token.STRING,
							Value: fmt.Sprintf(`"%s"`, field.name),
						},
					},
				},
				Type: field.kind,
			},
		},
	}
}

func ReturnValuesStmt(fields []*FieldDef) ast.Stmt {
	results := make([]ast.Expr, len(fields))
	for index, field := range fields {
		results[index] = NewIdentObjVar(field.name)
	}
	return &ast.ReturnStmt{
		Results: results,
	}
}
