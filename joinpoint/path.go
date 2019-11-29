package joinpoint

import (
	"fmt"
	"github.com/wesovilabs/beyond/logger"
	"go/ast"
	"reflect"
	"strings"
)

func buildPath(rootPkg string, pkg, objType string, funcDecl *ast.FuncDecl, imports map[string]string) string {
	in := pathForFieldList(rootPkg, pkg, funcDecl.Type.Params, imports, true)
	out := pathForFieldList(rootPkg, pkg, funcDecl.Type.Results, imports, false)
	result := fmt.Sprintf("%s%s%s", funcDecl.Name.String(), in, out)

	if objType != "" {
		result = fmt.Sprintf("%s.%s", getValue(objType, imports), result)
	}

	if pkg != "" {
		result = fmt.Sprintf("%s.%s", pkg, result)
	}

	return result
}

func exprPath(rootPkg string, pkg string, expr ast.Expr, imports map[string]string) string {
	switch val := expr.(type) {
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", exprPath(rootPkg, pkg, val.X, imports), exprPath(rootPkg, pkg, val.Sel, imports))
	case *ast.Ident:
		isObj := val.Obj != nil
		v := getValue(val.Name, imports)

		if !isObj {
			return v
		}

		return fmt.Sprintf("%s/%s.%s", rootPkg, pkg, v)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StructType:
		return "struct{}"
	case *ast.StarExpr:
		return exprPathStarExpr(rootPkg, pkg, val, imports)
	case *ast.FuncType:
		params := pathForFieldList(rootPkg, pkg, val.Params, imports, true)
		result := pathForFieldList(rootPkg, pkg, val.Results, imports, false)

		return fmt.Sprintf("func%s%s", params, result)
	}

	return ""
}

func exprPathStarExpr(rootPkg string, pkg string, val *ast.StarExpr, imports map[string]string) string {
	switch t2 := val.X.(type) {
	case *ast.Ident:
		v := getValue(t2.Name, imports)
		return fmt.Sprintf("*%s", v)
	case *ast.SelectorExpr:
		return fmt.Sprintf("*%s.%s", exprPath(rootPkg, pkg, t2.X, imports), t2.Sel.Name)
	default:
		logger.Infof("*%s", reflect.TypeOf(t2))
	}

	return ""
}

// nolint: gocyclo
func pathForSingleFieldList(rootPkg string, pkg string, field *ast.Field, imports map[string]string,
	forceParen bool) string {
	switch fieldType := field.Type.(type) {
	case *ast.Ident:
		isObj := fieldType.Obj != nil
		v := getValue(fieldType.Name, imports)

		if !isObj {
			return v
		}

		return fmt.Sprintf("%s/%s.%s", rootPkg, pkg, v)
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", exprPath(rootPkg, pkg, fieldType.X, imports))
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", exprPath(rootPkg, pkg, fieldType.X, imports), fieldType.Sel.Name)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StructType:
		return "struct{}"
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", exprPath(rootPkg, pkg, fieldType.Elt, imports))
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", exprPath(rootPkg, pkg, fieldType.Key, imports), exprPath(rootPkg,
			pkg, fieldType.Value, imports))
	case *ast.FuncType:
		params := pathForFieldList(rootPkg, pkg, fieldType.Params, imports, true)
		result := pathForFieldList(rootPkg, pkg, fieldType.Results, imports, forceParen)

		return fmt.Sprintf("func%s%s", params, result)
	case *ast.Ellipsis:
		return fmt.Sprintf("[]%s", exprPath(rootPkg, pkg, fieldType.Elt, imports))
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

func pathForSomeFieldsList(rootPkg string, pkg string, fields []*ast.Field, imports map[string]string,
	forceParen bool) string {
	values := make([]string, lenFields(fields))
	index := 0

	for _, field := range fields {
		if len(field.Names) > 0 {
			for range field.Names {
				values[index] = pathForSingleFieldList(rootPkg, pkg, field, imports, forceParen)
				index++
			}
		} else {
			values[index] = pathForSingleFieldList(rootPkg, pkg, field, imports, forceParen)
			index++
		}
	}

	return strings.Join(values, ",")
}

func pathForFieldList(rootPkg string, pkg string, fieldList *ast.FieldList,
	imports map[string]string, forceParen bool) string {
	var value string

	switch {
	case fieldList == nil || lenFields(fieldList.List) == 0:
		value = ""
	case lenFields(fieldList.List) == 1:
		value = pathForSingleFieldList(rootPkg, pkg, fieldList.List[0], imports, forceParen)
	default:
		value = pathForSomeFieldsList(rootPkg, pkg, fieldList.List, imports, forceParen)
	}

	if forceParen || (fieldList != nil && lenFields(fieldList.List) > 1) {
		return fmt.Sprintf("(%s)", value)
	}

	return value
}
