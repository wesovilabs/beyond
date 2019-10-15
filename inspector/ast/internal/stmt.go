package internal

import (
	"fmt"
	"go/ast"
	"go/token"
)

func (a *ASTAdapter) callContextBackground() *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent(a.importName("context")),
			Sel: ast.NewIdent("Background"),
		},
		Args: []ast.Expr{},
	}
}

func (a *ASTAdapter) callWithFunctionName() *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent(a.importName("github.com/wesovilabs/goa/context")),
			Sel: ast.NewIdent("WithName"),
		},
		Args: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("\"%s\"", a.function.Name()),
			},
		},
	}
}

func (a *ASTAdapter) callWithPkgName() *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent(a.importName("github.com/wesovilabs/goa/context")),
			Sel: ast.NewIdent("WithPkg"),
		},
		Args: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("\"%s\"", a.function.Pkg()),
			},
		},
	}
}

func (a *ASTAdapter) createInputArgCall(field *ast.Field) []ast.Expr {
	expressions := make([]ast.Expr, len(field.Names))
	for index, name := range field.Names {
		expressions[index] = &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent(a.importName("github.com/wesovilabs/goa/context")),
				Sel: ast.NewIdent("NewArg"),
			},
			Args: []ast.Expr{
				ast.NewIdent(fmt.Sprintf("\"%s\"", name.Obj.Name)),
				ast.NewIdent(name.Obj.Name),
			},
		}
	}
	return expressions
}

func (a *ASTAdapter) callWithInput() *ast.CallExpr {
	params := make([]ast.Expr, 0)
	for _, field := range a.FunctionParamsList() {
		params = append(params, a.createInputArgCall(field)...)
	}
	goaCtx := a.importName("github.com/wesovilabs/goa/context")
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent(goaCtx),
			Sel: ast.NewIdent("WithInput"),
		},
		Args: []ast.Expr{
			&ast.CompositeLit{
				Type: &ast.SelectorExpr{
					X:   ast.NewIdent(goaCtx),
					Sel: ast.NewIdent("Input"),
				},
				Elts: params,
			},
		},
	}
}

func (a *ASTAdapter) callNewContext() *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent(a.importName("github.com/wesovilabs/goa/context")),
			Sel: ast.NewIdent("New"),
		},
		Args: []ast.Expr{
			a.callContextBackground(),
			a.callWithFunctionName(),
			a.callWithPkgName(),
			a.callWithInput(),
		},
	}
}

func (a *ASTAdapter) updateCtxAttributesStatements(field *ast.Field) []ast.Stmt {
	stmts := make([]ast.Stmt, len(field.Names)*2)
	for index, name := range field.Names {
		stmts[(index*2)+0] = &ast.AssignStmt{
			Tok: token.DEFINE,
			Lhs: []ast.Expr{
				&ast.Ident{
					Name: fmt.Sprintf("%sUnTyped", name.Name),
					Obj: &ast.Object{
						Name: fmt.Sprintf("%sUnTyped", name.Name),
						Kind: ast.Var,
					},
				},
			},
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   ast.NewIdent("goaCtx.In()"),
						Sel: ast.NewIdent("Get"),
					},
					Args: []ast.Expr{
						ast.NewIdent(fmt.Sprintf("\"%s\"", name.Name)),
					},
				},
			},
		}
		stmts[(index*2)+1] = &ast.AssignStmt{
			Tok: token.ASSIGN,
			Lhs: []ast.Expr{
				&ast.Ident{
					Name: name.Name,
					Obj: &ast.Object{
						Name: name.Name,
						Kind: ast.Var,
					},
				},
			},
			Rhs: []ast.Expr{
				&ast.SelectorExpr{
					X: ast.NewIdent(fmt.Sprintf("%sUnTyped", name)),
					Sel: &ast.Ident{

						Name: updateValue(field),
					},
				},
			},
		}
	}
	return stmts
}

func updateValue(field *ast.Field) string {
	switch f := field.Type.(type) {
	case *ast.Ident:
		return fmt.Sprintf("(%s)", f.Name)
	case *ast.StarExpr:
		return fmt.Sprintf("(*%s.%s)", f.X.(*ast.SelectorExpr).X.(*ast.Ident).String(), f.X.(*ast.SelectorExpr).Sel.String())
	}
	return ""

}

// UpdateFunction update function ast attributes
func (a *ASTAdapter) UpdateFunction() {
	statements := a.createStatements()
	a.AddFunctionStatements(statements)
}

func (a *ASTAdapter) createStatements() []ast.Stmt {
	statements := make([]ast.Stmt, 1)
	bodyStatements := []ast.Stmt{
		a.createAspectCallStatement(),
	}
	for _, field := range a.FunctionParamsList() {
		bodyStmt := a.updateCtxAttributesStatements(field)
		bodyStatements = append(bodyStatements, bodyStmt...)
	}
	statements[0] = &ast.IfStmt{
		Init: &ast.AssignStmt{
			Tok: token.DEFINE,
			Lhs: []ast.Expr{
				&ast.Ident{
					Name: "goaCtx",
					Obj: &ast.Object{
						Name: "goaCtx",
						Kind: ast.Var,
					},
				},
				&ast.Ident{
					Name: "err",
					Obj: &ast.Object{
						Name: "err",
						Kind: ast.Var,
					},
				},
			},
			Rhs: []ast.Expr{
				a.callNewContext(),
			},
		},
		Cond: &ast.BinaryExpr{
			X:  ast.NewIdent("err"),
			Y:  ast.NewIdent("nil"),
			Op: token.EQL,
		},
		Body: &ast.BlockStmt{
			List: bodyStatements,
		},
	}
	return statements
}

func (a *ASTAdapter) createAspectCallStatement() ast.Stmt {
	var aspectCall *ast.CallExpr
	if !a.InPackage() {
		aspectCall = &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				Sel: &ast.Ident{
					Name: a.aspect.Name(),
				},
				X: &ast.BasicLit{Kind: token.INTERFACE, Value: a.aspect.Pkg()},
			},
			Args: []ast.Expr{
				ast.NewIdent("goaCtx"),
			},
		}
	} else {
		aspectCall = &ast.CallExpr{
			Fun: ast.NewIdent(a.aspect.Name()),
			Args: []ast.Expr{
				ast.NewIdent("goaCtx"),
			},
		}
	}
	return &ast.ExprStmt{
		X: aspectCall,
	}
}
