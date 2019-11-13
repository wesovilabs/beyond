package wrapper

import (
	"fmt"
	"github.com/wesovilabs/goa/advice"
	"github.com/wesovilabs/goa/function"
	"github.com/wesovilabs/goa/wrapper/internal"
	"go/ast"
	"strings"
)

func hasAnyReturning(definitions map[string]*advice.Advice) bool {
	for _, d := range definitions {
		if d.HasReturning() {
			return true
		}
	}

	return false
}

func wrapBeforeStatements(definitions map[string]*advice.Advice, params []*internal.FieldDef) []ast.Stmt {
	argsVariable := "functionParams"
	stmts := make([]ast.Stmt, 0)
	// set values to context
	stmts = append(stmts, setArgsValues(argsVariable, "Params", params)...)
	// call aspects
	for name, d := range definitions {
		if d.HasBefore() {
			stmts = append(stmts, &ast.ExprStmt{
				X: internal.CallAspectBefore(name),
			})
		}
	}

	stmts = append(stmts, internal.ArgsToFunctionArgs(argsVariable, params)...)

	return stmts
}
func wrapReturningStatements(definitions map[string]*advice.Advice, results []*internal.FieldDef) []ast.Stmt {
	argsVariable := "functionResults"
	stmts := make([]ast.Stmt, 0)

	if len(results) > 0 {
		stmts = append(stmts, setArgsValues(argsVariable, "Results", results)...)
	}

	for name, d := range definitions {
		if d.HasReturning() {
			stmts = append(stmts, &ast.ExprStmt{
				X: internal.CallAspectReturning(name),
			})
		}
	}

	if len(results) > 0 {
		stmts = append(stmts, internal.ArgsToFunctionArgs(argsVariable, results)...)
	}

	return stmts
}

func wrapperFuncDecl(function *function.Function, definitions map[string]*advice.Advice) *ast.FuncDecl {
	imports := internal.GetImports(function.Parent())
	ensureImports(imports, requiredImports, function)
	recv := function.GetRecv()
	imports[function.Pkg()] = ""
	stmts := make([]ast.Stmt, 0)
	stmts = append(stmts, internal.AssignGoaContext(imports))
	stmts = append(stmts, internal.SetUpGoaContext(function)...)

	for name, d := range definitions {
		if importName, found := imports[d.Pkg()]; !found {
			index := strings.LastIndex(d.Pkg(), "/")
			pkgName := findImportName(imports, d.Pkg()[index+1:], d.Pkg())
			addImportSpec(function, importName, d.Pkg())

			imports[d.Pkg()] = pkgName
			stmts = append(stmts, internal.AssignAspect(name, pkgName, d.Name()))
		} else {
			if importName == "" {
				index := strings.LastIndex(d.Pkg(), "/")
				importName = d.Pkg()[index+1:]
			}

			stmts = append(stmts, internal.AssignAspect(name, importName, d.Name()))
		}
	}

	params := internal.Params(function.ParamsList())
	results := internal.Results(function.ResultsList())

	stmts = append(stmts, wrapBeforeStatements(definitions, params)...)

	if recv != nil {
		// Call function
		stmts = append(stmts, internal.CallMethodAndAssign(recv, function.Parent().Name.String(), "",
			fmt.Sprintf("%sInternal", function.Name()), params, results))
	} else {
		// Call function
		stmts = append(stmts, internal.CallFunctionAndAssign(function.Parent().Name.String(), "",
			fmt.Sprintf("%sInternal", function.Name()), params, results))
	}

	if hasAnyReturning(definitions) {
		stmts = append(stmts, wrapReturningStatements(definitions, results)...)
	}

	stmts = append(stmts, internal.ReturnValuesStmt(results))
	funcDecl := internal.FuncDecl(function.Name(), function.ParamsList(),
		function.ResultsList(), stmts)
	funcDecl.Recv = recv

	return funcDecl
}

func setArgsValues(name string, argsType string, params []*internal.FieldDef) []ast.Stmt {
	stmts := make([]ast.Stmt, len(params)+2)
	stmts[0] = internal.TakeArgs(name, argsType)

	for index, param := range params {
		stmts[index+1] = &ast.ExprStmt{
			X: internal.SetArgValue(name, param),
		}
	}

	stmts[len(params)+1] = internal.SetArgs(fmt.Sprintf("Set%s", argsType), name)

	return stmts
}
