package internal

import (
	"fmt"
	"go/ast"
	"go/token"
)

// ArgsToFunctionArgs return the list of statements
func ArgsToFunctionArgs(argType string, name string, fields []*FieldDef) []ast.Stmt {
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
							Value: fmt.Sprintf(`"%s"`, f.Name),
						},
					},
				},
			},
		})
		stmts = append(stmts, argumentToVariable(resultName, fmt.Sprintf("%s%v", argType, index), f))
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

func argumentToVariable(variable string, fieldName string, field *FieldDef) ast.Stmt {
	kind := field.Kind
	if ell, ok := kind.(*ast.Ellipsis); ok {
		kind = &ast.ArrayType{
			Elt: ell.Elt,
		}
	}

	assigmentStmt := &ast.AssignStmt{
		Tok: token.ASSIGN,
		Lhs: []ast.Expr{
			NewIdentObjVar(fieldName),
		},
		Rhs: []ast.Expr{
			&ast.TypeAssertExpr{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   NewIdentObjVar(variable),
						Sel: NewIdent("Value"),
					},
				},
				Type: kind,
			},
		},
	}

	return ifArgumentValueIsNotNil(variable, assigmentStmt)
}

// ReturnValuesStmt return the list of statements
func ReturnValuesStmt(fields []*FieldDef) ast.Stmt {
	results := make([]ast.Expr, len(fields))

	for index, field := range fields {
		results[index] = NewIdentObjVar(field.Name)
	}

	return &ast.ReturnStmt{
		Results: results,
	}
}

// TakeArgs takes the arguments from the method
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

// SetArgs set arguments
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
