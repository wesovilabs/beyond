package ast

import (
	"fmt"
	"github.com/wesovilabs/goa/aspect"
	"github.com/wesovilabs/goa/function"
	"go/ast"
	"go/token"
)

type AroundExecutor struct {
	CurrentImports map[string]string
	Function       *function.Function
	Aspect         aspect.Aspect
}

func (e *AroundExecutor) requiredImports() map[string]string {
	return map[string]string{
		"context":                           "context",
		"github.com/wesovilabs/goa/api":     "api",
		"github.com/wesovilabs/goa/context": "goaContext",
	}
}

func (e *AroundExecutor) Execute() {
	ensureImports(e.CurrentImports, e.requiredImports(), e.Function)
	e.createStatements()
}

func (e *AroundExecutor) createStatements() []ast.Stmt {
	statements := make([]ast.Stmt, 1)
	bodyStatements := []ast.Stmt{
		e.createAspectCallStatement(),
	}
	for _, field := range e.Function.ParamsList() {
		bodyStmt := e.updateCtxAttributesStatements(field)
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
				e.callNewContext(),
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
	e.Function.AddStatementsAtBegin(statements)
	return statements
}

func (e *AroundExecutor) callContextBackground() *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent(e.CurrentImports["context"]),
			Sel: ast.NewIdent("Background"),
		},
		Args: []ast.Expr{},
	}
}

func (e *AroundExecutor) callWithFunctionName() *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent(e.CurrentImports["github.com/wesovilabs/goa/context"]),
			Sel: ast.NewIdent("WithName"),
		},
		Args: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("\"%s\"", e.Function.Name()),
			},
		},
	}
}

func (e *AroundExecutor) callWithPkgName() *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent(e.CurrentImports["github.com/wesovilabs/goa/context"]),
			Sel: ast.NewIdent("WithPkg"),
		},
		Args: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("\"%s\"", e.Function.Pkg()),
			},
		},
	}
}

func (e *AroundExecutor) createInputArgCall(field *ast.Field) []ast.Expr {
	expressions := make([]ast.Expr, len(field.Names))
	for index, name := range field.Names {
		expressions[index] = &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent(e.CurrentImports["github.com/wesovilabs/goa/context"]),
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

func (e *AroundExecutor) callWithInput() *ast.CallExpr {
	params := make([]ast.Expr, 0)
	for _, field := range e.Function.ParamsList() {
		params = append(params, e.createInputArgCall(field)...)
	}
	goaCtx := e.CurrentImports["github.com/wesovilabs/goa/context"]
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

func (e *AroundExecutor) createAspectCallStatement() ast.Stmt {
	var aspectCall *ast.CallExpr
	if e.Function.Pkg() != e.Aspect.Pkg() {
		aspectCall = &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				Sel: &ast.Ident{
					Name: e.Aspect.Name(),
				},
				X: &ast.BasicLit{Kind: token.INTERFACE, Value: e.Aspect.Pkg()},
			},
			Args: []ast.Expr{
				ast.NewIdent("goaCtx"),
			},
		}
	} else {
		aspectCall = &ast.CallExpr{
			Fun: ast.NewIdent(e.Aspect.Name()),
			Args: []ast.Expr{
				ast.NewIdent("goaCtx"),
			},
		}
	}
	return &ast.ExprStmt{
		X: aspectCall,
	}
}

func (e *AroundExecutor) updateCtxAttributesStatements(field *ast.Field) []ast.Stmt {
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

func (e *AroundExecutor) callNewContext() *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent(e.CurrentImports["github.com/wesovilabs/goa/context"]),
			Sel: ast.NewIdent("NewAroundContext"),
		},
		Args: []ast.Expr{
			e.callContextBackground(),
			e.callWithFunctionName(),
			e.callWithPkgName(),
			e.callWithInput(),
		},
	}
}
