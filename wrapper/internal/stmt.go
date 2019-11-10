package internal

import (
	"fmt"
	"go/ast"
	"go/token"
)

// AssignValuesFromContextIn return the list of statements
func AssignValuesFromContextIn(fields []*FieldDef) []ast.Stmt {
	stmts := make([]ast.Stmt, len(fields))
	for index, f := range fields {
		stmts[index] = getFromContext("GetInValue", f)
	}

	return stmts
}

// AssignValuesFromContextOut return the list of statements
func AssignValuesFromContextOut(fields []*FieldDef) []ast.Stmt {
	stmts := make([]ast.Stmt, len(fields))
	for index, field := range fields {
		stmts[index] = getFromContext("GetOutValue", field)
	}

	return stmts
}

/**
if goaContext.GetOutValue("result1")!=nil{
                result1 = goaContext.GetOutValue("result1").(error)
        }
*/

func checkIfValueIsNotNil(op string, field *FieldDef, stmt ast.Stmt) ast.Stmt {
	return &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   NewIdentObjVar(varGoaContext),
					Sel: NewIdent(op),
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: fmt.Sprintf(`"%s"`, field.name),
					},
				},
			},
			Op: token.NEQ,
			Y:  NewIdent("nil"),
		},

		Body: &ast.BlockStmt{List: []ast.Stmt{stmt}},
	}
}

func getFromContext(op string, field *FieldDef) ast.Stmt {
	return checkIfValueIsNotNil(op, field, &ast.AssignStmt{
		Tok: token.ASSIGN,
		Lhs: []ast.Expr{
			NewIdentObjVar(field.name),
		},
		Rhs: []ast.Expr{
			&ast.TypeAssertExpr{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   NewIdentObjVar(varGoaContext),
						Sel: NewIdent(op),
					},
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
