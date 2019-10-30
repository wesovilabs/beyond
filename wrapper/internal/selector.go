package internal

import (
	"go/ast"
)

const (
	contextIn  = "In"
	contextOut = "Out"
	get        = "Get"
	set        = "Set"
	goaCtx     = "github.com/wesovilabs/goa/api/context"
)

var (
	selectorOutGet = selectorContextOperation(contextOut, get)
	selectorOutSet = selectorContextOperation(contextOut, set)
	selectorInGet  = selectorContextOperation(contextIn, get)
	selectorInSet  = selectorContextOperation(contextIn, set)
)

func selectorContextOperation(name string, op string) *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   NewIdentObjVar(varGoaContext),
				Sel: NewIdent(name),
			},
		},
		Sel: NewIdent(op),
	}
}
