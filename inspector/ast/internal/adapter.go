package internal

import (
	"github.com/wesovilabs/goa/inspector"
	"github.com/wesovilabs/goa/inspector/aspect"
	"go/ast"
)

// ASTAdapter struct used to modify the ast
type ASTAdapter struct {
	function *inspector.Function
	aspect   *aspect.Aspect
	imports  map[string]string
}

// New reutrns an instace of ASTAdapter
func New(function *inspector.Function, aspect *aspect.Aspect, imports map[string]string) *ASTAdapter {
	return &ASTAdapter{
		function: function,
		aspect:   aspect,
		imports:  imports,
	}
}

func (a *ASTAdapter) importName(path string) string {
	return a.imports[path]
}

func (a *ASTAdapter) AddFunctionStatements(statements []ast.Stmt) {
	a.function.AddStatementsAtBegin(statements)
}

func (a *ASTAdapter) FunctionParamsList() []*ast.Field {
	return a.function.ParamsList()
}

func (a *ASTAdapter) InPackage() bool {
	return a.function.Pkg() == a.aspect.Pkg()
}
