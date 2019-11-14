package internal

import (
	"go/ast"
	"go/token"
)

const (
	varGoaContext = "goaContext"
)

// AssignGoaContext create a new assignment
func AssignGoaContext(imports map[string]string) *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			NewIdentObjVar(varGoaContext),
		},
		Rhs: []ast.Expr{
			CallCreateGoaContext(imports),
		},
		Tok: token.DEFINE,
	}
}

// AssignAspect create a new assignment
func AssignAspect(name, pkg, function string) *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			NewIdentObjVar(name),
		},
		Rhs: []ast.Expr{
			CallCreateAspect(pkg, function),
		},
		Tok: token.DEFINE,
	}
}

// CallFunctionAndAssign create a new assignment
func CallFunctionAndAssign(currentPkg, pkg, name string, params, results []*FieldDef) ast.Stmt {
	if len(results) > 0 {
		outputVariables := make([]ast.Expr, len(results))
		for index, field := range results {
			outputVariables[index] = NewIdentObj(field.name)
		}

		return &ast.AssignStmt{
			Tok: token.DEFINE,
			Lhs: outputVariables,
			Rhs: []ast.Expr{
				CallFunction(currentPkg, pkg, name, params),
			},
		}
	}

	return &ast.ExprStmt{
		X: CallFunction(currentPkg, pkg, name, params),
	}
}

// CallMethodAndAssign create a new assignment
func CallMethodAndAssign(recv *ast.FieldList, currentPkg, pkg, name string, params, results []*FieldDef) ast.Stmt {
	objName := recv.List[0].Names[0].String()

	if len(results) > 0 {
		outputVariables := make([]ast.Expr, len(results))
		for index, field := range results {
			outputVariables[index] = NewIdentObj(field.name)
		}

		return &ast.AssignStmt{
			Tok: token.DEFINE,
			Lhs: outputVariables,
			Rhs: []ast.Expr{
				CallMethod(objName, currentPkg, pkg, name, params),
			},
		}
	}

	return &ast.ExprStmt{
		X: CallMethod(objName, currentPkg, pkg, name, params),
	}
}
