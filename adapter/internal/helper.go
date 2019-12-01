package internal

import (
	"fmt"
	"github.com/wesovilabs/beyond/logger"
	"go/ast"
	"reflect"
	"strings"
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
		params := make([]string, len(t.Params.List))
		if len(t.Params.List) > 0 {
			for i, _ := range t.Params.List {
				field := t.Results.List[i]
				params[i] = astToExpression(field.Type, false)
			}
		}
		results := make([]string, len(t.Results.List))
		if len(t.Results.List) > 0 {
			for i, _ := range t.Results.List {
				field := t.Results.List[i]
				results[i] = astToExpression(field.Type, false)
			}
		}

		return fmt.Sprintf("func(%s)%s", strings.Join(params, ","), strings.Join(results, ","))
	default:
		logger.Errorf("Unexpected type %s", reflect.TypeOf(t))
		return ""
	}
}
