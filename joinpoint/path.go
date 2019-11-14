package joinpoint

import (
	"fmt"
	"github.com/wesovilabs/goa/logger"
	"go/ast"
	"reflect"
	"strings"
)

func buildPath(pkg, objType string, funcDecl *ast.FuncDecl, imports map[string]string) string {
	in := pathForFieldList(funcDecl.Type.Params, imports, true)
	out := pathForFieldList(funcDecl.Type.Results, imports, false)
	result := fmt.Sprintf("%s%s%s", funcDecl.Name.String(), in, out)

	if objType != "" {
		result = fmt.Sprintf("%s.%s", getValue(objType, imports), result)
	}

	if pkg != "" {
		result = fmt.Sprintf("%s.%s", pkg, result)
	}

	return result
}

func exprPath(expr ast.Expr, imports map[string]string) string {
	switch val := expr.(type) {
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", exprPath(val.X, imports), exprPath(val.Sel, imports))
	case *ast.Ident:
		return getValue(val.Name, imports)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StructType:
		return "struct{}"
	case *ast.StarExpr:
		switch t2 := val.X.(type) {
		case *ast.Ident:
			v := getValue(t2.Name, imports)
			return fmt.Sprintf("*%s", v)
		case *ast.SelectorExpr:
			return fmt.Sprintf("*%s.%s", exprPath(t2.X, imports), t2.Sel.Name)
		default:
			logger.Infof("*%s", reflect.TypeOf(t2))
		}
	case *ast.FuncType:
		params := pathForFieldList(val.Params, imports, true)
		result := pathForFieldList(val.Results, imports, false)

		return fmt.Sprintf("func%s%s", params, result)
	}

	return ""
}

// nolint: gocyclo
func pathForSingleFieldList(field *ast.Field, imports map[string]string, forceParen bool) string {
	switch fieldType := field.Type.(type) {
	case *ast.Ident:
		return fieldType.Name
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", exprPath(fieldType.X, imports))
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", exprPath(fieldType.X, imports), fieldType.Sel.Name)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StructType:
		return "struct{}"
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", exprPath(fieldType.Elt, imports))
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", exprPath(fieldType.Key, imports), exprPath(fieldType.Value, imports))
	case *ast.FuncType:
		params := pathForFieldList(fieldType.Params, imports, true)
		result := pathForFieldList(fieldType.Results, imports, forceParen)

		return fmt.Sprintf("func%s%s", params, result)
	case *ast.Ellipsis:
		return fmt.Sprintf("[]%s", exprPath(fieldType.Elt, imports))
	default:
		return ""
	}
}

func getValue(name string, imports map[string]string) string {
	if i, ok := imports[name]; ok {
		return i
	}

	return name
}

func lenFields(fields []*ast.Field) int {
	totalLen := 0

	for _, f := range fields {
		if len(f.Names) == 0 {
			totalLen++
		}

		totalLen += len(f.Names)
	}

	return totalLen
}

func pathForSomeFieldsList(fields []*ast.Field, imports map[string]string, forceParen bool) string {
	values := make([]string, lenFields(fields))
	index := 0

	for _, field := range fields {
		if len(field.Names) > 0 {
			for range field.Names {
				values[index] = pathForSingleFieldList(field, imports, forceParen)
				index++
			}
		} else {
			values[index] = pathForSingleFieldList(field, imports, forceParen)
			index++
		}
	}

	return strings.Join(values, ",")
}

func pathForFieldList(fieldList *ast.FieldList, imports map[string]string, forceParen bool) string {
	var value string

	switch {
	case fieldList == nil || lenFields(fieldList.List) == 0:
		value = ""
	case lenFields(fieldList.List) == 1:
		value = pathForSingleFieldList(fieldList.List[0], imports, forceParen)
	default:
		value = pathForSomeFieldsList(fieldList.List, imports, forceParen)
	}

	if forceParen || (fieldList != nil && lenFields(fieldList.List) > 1) {
		return fmt.Sprintf("(%s)", value)
	}

	return value
}
