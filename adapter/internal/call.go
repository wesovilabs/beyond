package internal

import (
	"fmt"
	"github.com/wesovilabs/goa/joinpoint"
	"go/ast"
	"go/token"
)

const (
	opBefore     = "Before"
	opReturning  = "Returning"
	opNewContext = "NewContext"
)

// CallAspectBefore reutrn the call expression
func CallAspectBefore(name string) *ast.CallExpr {
	return goaInterceptor(name, opBefore)
}

// CallAspectReturning reutrn the call expression
func CallAspectReturning(name string) *ast.CallExpr {
	return goaInterceptor(name, opReturning)
}

// CallCreateGoaContext reutrn the call expression
func CallCreateGoaContext(imports map[string]string) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   NewIdent(imports[goaCtx]),
			Sel: NewIdent(opNewContext),
		},
		Args: []ast.Expr{},
	}
}

func goaInterceptor(name string, operation string) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   NewIdentObjVar(name),
			Sel: NewIdent(operation),
		},
		Args: []ast.Expr{
			NewIdentObj(varGoaContext),
		},
	}
}

// CallCreateAspect return the call expression
func CallCreateAspect(pkg, name string) ast.Expr {
	if pkg != "" {
		return NewIdent(fmt.Sprintf("%s.%s", pkg, name))
	}

	return NewIdent(name)
}

// SetArgValue set value to context
func SetArgValue(argsType string, field *FieldDef) ast.Expr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   NewIdentObjVar(argsType),
			Sel: NewIdent("Set"),
		},
		Args: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf(`"%s"`, field.name),
			},
			NewIdentObjVar(field.name),
		},
	}
}

// CallFunction return the call expression
func CallFunction(currentPkg, pkg, name string, fields []*FieldDef) *ast.CallExpr {
	args := make([]ast.Expr, len(fields))
	for index, field := range fields {
		args[index] = NewIdentObj(field.name)
	}

	if currentPkg == pkg || pkg == "" {
		return &ast.CallExpr{
			Fun:  NewIdent(name),
			Args: args,
		}
	}

	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   NewIdentObj(pkg),
			Sel: NewIdent(name),
		},
		Args: args,
	}
}

// CallMethod return the call expression
func CallMethod(objName string, currentPkg, pkg, name string, fields []*FieldDef) ast.Expr {
	args := make([]ast.Expr, len(fields))
	for index, field := range fields {
		args[index] = NewIdentObj(field.name)
	}

	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   NewIdentObjVar(objName),
			Sel: NewIdent(name),
		},
		Args: args,
	}
}

// SetUpGoaContext return the list of required statements
func SetUpGoaContext(f *joinpoint.JoinPoint) []ast.Stmt {
	stmts := make([]ast.Stmt, 2)
	stmts[0] = &ast.ExprStmt{
		X: CallFunction("", varGoaContext, "WithPkg", []*FieldDef{
			{
				name: fmt.Sprintf(`"%s"`, f.Pkg()),
				kind: NewIdent(f.Pkg()),
			},
		}),
	}
	stmts[1] = &ast.ExprStmt{
		X: CallFunction("", varGoaContext, "WithName", []*FieldDef{
			{
				name: fmt.Sprintf(`"%s"`, f.Name()),
				kind: NewIdent(f.Name()),
			},
		}),
	}

	if f.GetRecv() != nil {
		objName := f.GetRecv().List[0].Names[0].String()
		// objType:=f.GetRecv().List[0].Type
		stmts = append(stmts, &ast.ExprStmt{
			X: CallFunction("", varGoaContext, "WithType", []*FieldDef{
				{
					name: objName,
					kind: NewIdent(f.Name()),
				},
			}),
		})
	}

	return stmts
}
