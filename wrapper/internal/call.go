package internal

import (
	"fmt"
	"github.com/wesovilabs/goa/function"
	"go/ast"
	"go/token"
)

const (
	opBefore     = "Before"
	opReturning  = "Returning"
	opNewContext = "NewContext"
)

func CallAspectBefore(name string) *ast.CallExpr {
	return goaInterceptor(name, opBefore)
}
func CallAspectReturning(name string) *ast.CallExpr {
	return goaInterceptor(name, opReturning)
}

func CallCreateGoaContext(imports map[string]string) *ast.CallExpr {

	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   NewIdent(imports[goaCtx]),
			Sel: NewIdent(opNewContext),
		},
		Args: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   NewIdent(imports["context"]),
					Sel: NewIdent("Background"),
				},
			},
		},
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

func CallCreateAspect(pkg, name string) *ast.CallExpr {
	if pkg != "" {
		return &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   NewIdent(pkg),
				Sel: NewIdent(name),
			},
		}
	}
	return &ast.CallExpr{
		Fun: NewIdent(name),
	}
}

func SetCtxIn(field *FieldDef) *ast.CallExpr {
	return setContext(SelectorInSet, field.name)
}

func GetCtxIn(field *FieldDef) *ast.CallExpr {
	return getContext(SelectorInGet, field.name)
}

func SetCtxOut(field *FieldDef) *ast.CallExpr {
	return setContext(SelectorOutSet, field.name)
}

func setContext(selector *ast.SelectorExpr, name string) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: selector,
		Args: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf(`"%s"`, name),
			},
			NewIdentObjVar(name),
		},
	}
}

func getContext(selector *ast.SelectorExpr, name string) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: selector,
		Args: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf(`"%s"`, name),
			},
			NewIdentObjVar(name),
		},
	}
}

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

func SetUpGoaContext(f *function.Function) []ast.Stmt {
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
	return stmts
}
