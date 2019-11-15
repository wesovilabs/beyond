package adapter

import (
	"fmt"
	"github.com/wesovilabs/goa/adapter/internal"
	"github.com/wesovilabs/goa/advice"
	"github.com/wesovilabs/goa/joinpoint"
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
	argsVariable := "joinPointParams"
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
	argsVariable := "joinPointResults"
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

func adapterFuncDecl(joinPoint *joinpoint.JoinPoint, advices map[string]*advice.Advice) *ast.FuncDecl {
	imports := internal.GetImports(joinPoint.Parent())
	ensureImports(imports, requiredImports, joinPoint)
	recv := joinPoint.GetRecv()
	imports[joinPoint.Pkg()] = ""
	stmts := make([]ast.Stmt, 0)
	stmts = append(stmts, internal.AssignGoaContext(imports))
	stmts = append(stmts, internal.SetUpGoaContext(joinPoint)...)

	for name, d := range advices {
		if importName, found := imports[d.Pkg()]; !found {
			index := strings.LastIndex(d.Pkg(), "/")
			pkgName := findImportName(imports, d.Pkg()[index+1:], d.Pkg())
			addImportSpec(joinPoint, importName, d.Pkg())

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

	params := internal.Params(joinPoint.ParamsList())
	results := internal.Results(joinPoint.ResultsList())

	stmts = append(stmts, wrapBeforeStatements(advices, params)...)

	if recv != nil {
		// Call joinPoint
		stmts = append(stmts, internal.CallMethodAndAssign(recv, joinPoint.Parent().Name.String(), "",
			fmt.Sprintf("%sInternal", joinPoint.Name()), params, results))
	} else {
		// Call joinPoint
		stmts = append(stmts, internal.CallFunctionAndAssign(joinPoint.Parent().Name.String(), "",
			fmt.Sprintf("%sInternal", joinPoint.Name()), params, results))
	}

	if hasAnyReturning(advices) {
		stmts = append(stmts, wrapReturningStatements(advices, results)...)
	}

	stmts = append(stmts, internal.ReturnValuesStmt(results))
	funcDecl := internal.FuncDecl(joinPoint.Name(), joinPoint.ParamsList(),
		joinPoint.ResultsList(), stmts)
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
