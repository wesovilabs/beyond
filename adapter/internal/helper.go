package internal

import (
	"fmt"
	"github.com/wesovilabs/goa/logger"
	"go/ast"
	"reflect"
)

func astToExpression(expr ast.Expr, root bool) string {
	kind := astToExpressionEval(expr)

	if root {
		return fmt.Sprintf("\"%s\"", kind)
	}

	return kind
}

func astToExpressionEval(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", astToExpression(t.Key, false), astToExpression(t.Value, false))
	case *ast.Ident:
		return t.String()
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", astToExpression(t.X, false))
	case *ast.StructType:
		return "struct{}"
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", astToExpression(t.Elt, false))
	case *ast.Ellipsis:
		return fmt.Sprintf("[]%s", astToExpression(t.Elt, false))
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", astToExpression(t.X, false), t.Sel.String())
	case *ast.FuncType:

	default:
		logger.Errorf("Unexpected type %s", reflect.TypeOf(t))
	}

	return ""
}
