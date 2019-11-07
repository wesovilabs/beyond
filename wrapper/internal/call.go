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

// CallCreateAspect return the call expression
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

// SetCtxIn return the call expression
func SetCtxIn(field *FieldDef) *ast.CallExpr {
	return setContext("SetIn", field.name)
}

// SetCtxOut return the call expression
func SetCtxOut(field *FieldDef) *ast.CallExpr {
	return setContext("SetOut", field.name)
}

func setContext(op string, name string) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   NewIdentObjVar(varGoaContext),
			Sel: NewIdent(op),
		},
		Args: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf(`"%s"`, name),
			},
			NewIdentObjVar(name),
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

	if f.GetRecv() != nil {
		objName := f.GetRecv().List[0].Names[0].String()
		//objType:=f.GetRecv().List[0].Type
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
