package function

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

func buildPath(fileDecl *ast.File, funcDecl *ast.FuncDecl) string {
	in := pathForFieldList(funcDecl.Type.Params, true)
	out := pathForFieldList(funcDecl.Type.Results, false)
	return fmt.Sprintf("%s.%s%s%s", fileDecl.Name.Name, funcDecl.Name.String(), in, out)
}

func exprPath(expr ast.Expr) string {
	switch val := expr.(type) {
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", exprPath(val.X), exprPath(val.Sel))
	case *ast.Ident:
		return val.Name
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StructType:
		return "struct{}"
	case *ast.StarExpr:
		switch t2 := val.X.(type) {
		case *ast.Ident:
			return fmt.Sprintf("*%s", t2.Name)
		case *ast.SelectorExpr:
			return fmt.Sprintf("%s.%s", t2.X, t2.Sel.Name)
		default:
			fmt.Println(reflect.TypeOf(t2))
		}
	default:
		fmt.Println(val)
	}
	return ""
}

// nolint: gocyclo
func pathForSingleFieldList(field *ast.Field, forceParen bool) string {
	switch fieldType := field.Type.(type) {
	case *ast.Ident:
		return fieldType.Name
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", exprPath(fieldType.X))
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", fieldType.X, fieldType.Sel.Name)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StructType:
		return "struct{}"
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", exprPath(fieldType.Elt))
	case *ast.MapType:
		if star, ok := fieldType.Value.(*ast.StarExpr); ok {
			return fmt.Sprintf("map[%s]*%s", exprPath(fieldType.Key), exprPath(star))
		}
		return fmt.Sprintf("map[%s]%s", exprPath(fieldType.Key), exprPath(fieldType.Value))
	case *ast.FuncType:
		params := pathForFieldList(fieldType.Params, true)
		result := pathForFieldList(fieldType.Results, forceParen)
		return fmt.Sprintf("func%s%s", params, result)
	case *ast.Ellipsis:
		return fmt.Sprintf("...%s", exprPath(fieldType.Elt))
	default:
		return ""
	}
}

func pathForSomeFieldsList(fields []*ast.Field, forceParen bool) string {
	values := make([]string, len(fields))
	for index, field := range fields {
		values[index] = pathForSingleFieldList(field, forceParen)
	}
	return strings.Join(values, ",")
}

func pathForFieldList(fieldList *ast.FieldList, forceParen bool) string {
	var value string

	switch {
	case fieldList == nil || len(fieldList.List) == 0:
		value = ""
	case len(fieldList.List) == 1:
		value = pathForSingleFieldList(fieldList.List[0], forceParen)
	default:
		value = pathForSomeFieldsList(fieldList.List, forceParen)
	}

	if forceParen || (fieldList != nil && len(fieldList.List) > 1) {
		return fmt.Sprintf("(%s)", value)
	}
	return value

}
