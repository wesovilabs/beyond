package wrapper

import (
	"fmt"
	"github.com/wesovilabs/goa/aspect"
	"github.com/wesovilabs/goa/function"
	"github.com/wesovilabs/goa/wrapper/internal"
	"go/ast"
	"strings"
)

const (
	functionSuffix = ""
	resultPrefix   = "out"
)

func hasAnyBefore(definitions map[string]*aspect.Definition) bool {
	for _, d := range definitions {
		if d.HasBefore() {
			return true
		}
	}
	return false
}
func hasAnyReturning(definitions map[string]*aspect.Definition) bool {
	for _, d := range definitions {
		if d.HasReturning() {
			return true
		}
	}
	return false
}

func wrapperFuncDecl(function *function.Function, definitions map[string]*aspect.Definition) *ast.FuncDecl {
	imports := internal.GetImports(function.Parent())
	ensureImports(imports, requiredImports, function)
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
	if hasAnyBefore(definitions) {
		// set values to context
		stmts = append(stmts, setValuesToContextIn(params)...)
		// call aspects
		for name, d := range definitions {
			if d.HasBefore() {
				stmts = append(stmts, &ast.ExprStmt{
					X: internal.CallAspectBefore(name),
				})
			}
		}
		stmts = append(stmts, internal.AssignValuesFromContextIn(params)...)
	}
	// Call function
	stmts = append(stmts, internal.CallFunctionAndAssign(function.Parent().Name.String(), "", fmt.Sprintf("%sInternal", function.Name()), params, results))

	if hasAnyReturning(definitions) {
		if len(results) > 0 {
			stmts = append(stmts, setValuesToContextOut(results)...)
		}
		for name, d := range definitions {
			if d.HasReturning() {
				stmts = append(stmts, &ast.ExprStmt{
					X: internal.CallAspectReturning(name),
				})
			}
		}
		if len(results) > 0 {
			stmts = append(stmts, internal.AssignValuesFromContextOut(results)...)
		}
	}
	stmts = append(stmts, internal.ReturnValuesStmt(results))
	return internal.FuncDecl(fmt.Sprintf("%s%s", function.Name(), functionSuffix), function.ParamsList(), function.ResultsList(), stmts)
}

func setValuesToContextIn(params []*internal.FieldDef) []ast.Stmt {
	statements := make([]ast.Stmt, len(params))
	for index, param := range params {
		statements[index] = &ast.ExprStmt{
			X: internal.SetCtxIn(param),
		}
	}
	return statements
}

func setValuesToContextOut(results []*internal.FieldDef) []ast.Stmt {
	stmts := make([]ast.Stmt, len(results))
	for index, result := range results {
		stmts[index] = &ast.ExprStmt{
			X: internal.SetCtxOut(result),
		}
	}
	return stmts
}
