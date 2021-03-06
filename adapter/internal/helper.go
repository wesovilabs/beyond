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

//nolint: gocyclo
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
		return astFuncToExpressionPath(t)
	default:
		logger.Errorf("Unexpected type %s", reflect.TypeOf(t))
		return ""
	}
}

func astFuncToExpressionPath(t *ast.FuncType) string {
	var params, results []string

	if t.Params != nil && t.Params.List != nil {
		params = make([]string, len(t.Params.List))

		for i := range t.Params.List {
			field := t.Params.List[i]
			params[i] = astToExpression(field.Type, false)
		}
	}

	if t.Results != nil && t.Results.List != nil {
		results = make([]string, len(t.Results.List))

		for i := range t.Results.List {
			field := t.Results.List[i]
			results[i] = astToExpression(field.Type, false)
		}
	}

	return fmt.Sprintf("func(%s)%s", strings.Join(params, ","), strings.Join(results, ","))
}
