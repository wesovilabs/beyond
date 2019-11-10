package internal

import (
	"fmt"
	"go/ast"
	"go/token"
)

// ArgsToFunctionArgs return the list of statements
func ArgsToFunctionArgs(name string, fields []*FieldDef) []ast.Stmt {
	stmts := make([]ast.Stmt, 0)

	for index, f := range fields {
		resultName := fmt.Sprintf("%s%v", name, index)

		stmts = append(stmts, &ast.AssignStmt{
			Lhs: []ast.Expr{
				NewIdentObj(resultName),
			},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   NewIdentObjVar(name),
						Sel: NewIdent("Get"),
					},
					Args: []ast.Expr{
						&ast.BasicLit{
							Kind:  token.STRING,
							Value: fmt.Sprintf(`"%s"`, f.name),
						},
					},
				},
			},
		})
		stmts = append(stmts, argumentToVariable(resultName, f))
	}

	return stmts
}

func ifArgumentValueIsNotNil(variable string, stmt ast.Stmt) ast.Stmt {
	return &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   NewIdentObjVar(variable),
					Sel: NewIdent("Value"),
				},
			},
			Op: token.NEQ,
			Y:  NewIdent("nil"),
		},
		Body: &ast.BlockStmt{List: []ast.Stmt{stmt}},
	}
}

func argumentToVariable(variable string, field *FieldDef) ast.Stmt {
	return ifArgumentValueIsNotNil(variable, &ast.AssignStmt{
		Tok: token.ASSIGN,
		Lhs: []ast.Expr{
			NewIdentObjVar(field.name),
		},
		Rhs: []ast.Expr{
			&ast.TypeAssertExpr{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   NewIdentObjVar(variable),
						Sel: NewIdent("Value"),
					},
				},
				Type: field.kind,
			},
		},
	})
}

// ReturnValuesStmt return the list of statements
func ReturnValuesStmt(fields []*FieldDef) ast.Stmt {
	results := make([]ast.Expr, len(fields))

	for index, field := range fields {
		results[index] = NewIdentObjVar(field.name)
	}

	return &ast.ReturnStmt{
		Results: results,
	}
}

func TakeArgs(name string, method string) ast.Stmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			NewIdentObj(name),
		},
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   NewIdentObjVar(varGoaContext),
					Sel: NewIdentObj(method),
				},
			},
		},
		Tok: token.DEFINE,
	}
}

func SetArgs(method string, name string) ast.Stmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   NewIdentObjVar(varGoaContext),
				Sel: NewIdentObj(method),
			},
			Args: []ast.Expr{
				NewIdentObj(name),
			},
		},
	}
}
