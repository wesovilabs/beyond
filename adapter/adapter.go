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

	stmts = append(stmts, internal.ArgsToFunctionArgs("param", argsVariable, params)...)

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
		stmts = append(stmts, internal.ArgsToFunctionArgs("result", argsVariable, results)...)
	}

	return stmts
}

func adapterFuncDecl(joinPoint *joinpoint.JoinPoint, advices map[string]*advice.Advice) *ast.FuncDecl {
	imports := internal.GetImports(joinPoint.Parent())

	for i := range advices {
		advice := advices[i]
		for j := range advice.Imports() {
			importPath := advice.Imports()[j]
			if importPath != joinPoint.PkgPath() && imports[importPath] == "" {
				lastIndex := strings.LastIndex(importPath, "/")
				requiredImports[importPath] = findImportName(requiredImports, importPath[lastIndex+1:], importPath)
			}
		}
	}

	delete(requiredImports, joinPoint.PkgPath())
	ensureImports(imports, requiredImports, joinPoint)
	recv := joinPoint.GetRecv()
	imports[joinPoint.Pkg()] = ""
	stmts := make([]ast.Stmt, 0)
	stmts = append(stmts, internal.AssignGoaContext(imports))
	stmts = append(stmts, internal.SetUpGoaContext(joinPoint)...)

	for name, advice := range advices {
		stmts = append(stmts, applyAdvices(name, advice, imports, joinPoint)...)
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

func applyAdvices(name string, advice *advice.Advice, imports map[string]string,
	joinPoint *joinpoint.JoinPoint) []ast.Stmt {
	stmts := make([]ast.Stmt, 0)

	if importName, found := imports[advice.Pkg()]; !found {
		index := strings.LastIndex(advice.Pkg(), "/")
		pkgName := findImportName(imports, advice.Pkg()[index+1:], advice.Pkg())

		if joinPoint.PkgPath() != advice.Pkg() {
			if i := addImportSpec(joinPoint, importName, advice.Pkg()); i != nil {
				imports[advice.Pkg()] = pkgName
			}

			stmts = append(stmts, internal.AssignAspect(name, pkgName, advice.GetAdviceCall(joinPoint.PkgPath(), imports)))
		} else {
			imports[advice.Pkg()] = pkgName
			stmts = append(stmts, internal.AssignAspect(name, "", advice.GetAdviceCall(joinPoint.PkgPath(), imports)))
		}
	} else {
		if importName == "" {
			index := strings.LastIndex(advice.Pkg(), "/")
			importName = advice.Pkg()[index+1:]
		}

		stmts = append(stmts, internal.AssignAspect(name, importName, advice.GetAdviceCall(joinPoint.PkgPath(), imports)))
	}

	return stmts
}

func setArgsValues(name string, argsType string, params []*internal.FieldDef) []ast.Stmt {
	stmts := make([]ast.Stmt, len(params)+2)
	stmts[0] = internal.TakeArgs(name, argsType)

	for index, param := range params {
		paramName := param.Name
		if argsType == "Params" {
			paramName = fmt.Sprintf("param%v", index)
		}

		stmts[index+1] = &ast.ExprStmt{
			X: internal.SetArgValue(name, param, paramName),
		}
	}

	stmts[len(params)+1] = internal.SetArgs(fmt.Sprintf("Set%s", argsType), name)

	return stmts
}
