package internal

import (
	"fmt"
	"github.com/wesovilabs/beyond/joinpoint"
	"go/ast"
	"go/token"
)

const (
	beyondCtx    = "github.com/wesovilabs/beyond/api/context"
	opBefore     = "Before"
	opReturning  = "Returning"
	opNewContext = "NewContext"
)

// CallAspectBefore reutrn the call expression
func CallAspectBefore(name string) *ast.CallExpr {
	return beyondInterceptor(name, opBefore)
}

// CallAspectReturning reutrn the call expression
func CallAspectReturning(name string) *ast.CallExpr {
	return beyondInterceptor(name, opReturning)
}

// CallCreateBeyondContext reutrn the call expression
func CallCreateBeyondContext(imports map[string]string) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   NewIdent(imports[beyondCtx]),
			Sel: NewIdent(opNewContext),
		},
		Args: []ast.Expr{},
	}
}

func beyondInterceptor(name string, operation string) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   NewIdentObjVar(name),
			Sel: NewIdent(operation),
		},
		Args: []ast.Expr{
			NewIdentObj(varBeyondContext),
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
func SetArgValue(name string, field *FieldDef, paramName string) ast.Expr {
	kind := astToExpression(field.Kind, true)

	callExpr := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   NewIdentObjVar(name),
			Sel: NewIdent("SetWithType"),
		},
		Args: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf(`"%s"`, field.Name),
			},
			NewIdentObjVar(paramName),
			NewIdentObj(kind),
		},
	}

	return callExpr
}

func prepareArgs(fields []*FieldDef, withName bool) []ast.Expr {
	args := make([]ast.Expr, len(fields))

	for index, field := range fields {
		var param string
		if withName {
			param = field.Name
		} else {
			param = fmt.Sprintf("param%v", index)
		}

		switch field.Kind.(type) {
		case *ast.Ellipsis:
			args[index] = NewIdentObj(param + "...")
		default:
			args[index] = NewIdentObj(param)
		}
	}

	return args
}

// CallFunction return the call expression
func CallFunction(currentPkg, pkg, name string, fields []*FieldDef) *ast.CallExpr {
	argsWithName := pkg == "beyondContext" && (name == "WithPkg" || name == "WithName" || name == "WithType")
	args := prepareArgs(fields, argsWithName)

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
	args := prepareArgs(fields, false)

	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   NewIdentObjVar(objName),
			Sel: NewIdent(name),
		},
		Args: args,
	}
}

// SetUpBeyondContext return the list of required statements
func SetUpBeyondContext(f *joinpoint.JoinPoint) []ast.Stmt {
	stmts := make([]ast.Stmt, 2)
	stmts[0] = &ast.ExprStmt{
		X: CallFunction("", varBeyondContext, "WithPkg", []*FieldDef{
			{
				Name: fmt.Sprintf(`"%s"`, f.Pkg()),
				Kind: NewIdent(f.Pkg()),
			},
		}),
	}
	stmts[1] = &ast.ExprStmt{
		X: CallFunction("", varBeyondContext, "WithName", []*FieldDef{
			{
				Name: fmt.Sprintf(`"%s"`, f.Name()),
				Kind: NewIdent(f.Name()),
			},
		}),
	}

	if f.GetRecv() != nil {
		objName := f.GetRecv().List[0].Names[0].String()

		stmts = append(stmts, &ast.ExprStmt{
			X: CallFunction("", varBeyondContext, "WithType", []*FieldDef{
				{
					Name: objName,
					Kind: NewIdent(f.Name()),
				},
			}),
		})
	}

	return stmts
}
