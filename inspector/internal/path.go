package internal

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

// BuildPath return the signature for the input
func BuildPath(fileDecl *ast.File, funcDecl *ast.FuncDecl) string {
	in := pathForFieldList(funcDecl.Type.Params, true)
	out := pathForFieldList(funcDecl.Type.Results, false)
	return fmt.Sprintf("%s.%s%s%s", fileDecl.Name.Name, funcDecl.Name.String(), in, out)
}

func pathForFieldList(fieldList *ast.FieldList, forceParents bool) string {
	if fieldList == nil {
		return ""
	}
	switch len(fieldList.List) {
	case 0:
		if forceParents {
			return "()"
		}
		return ""
	case 1:
		var value string

		switch t := fieldList.List[0].Type.(type) {
		case *ast.Ident:
			value = t.Name
		case *ast.StarExpr:
			switch t2 := t.X.(type) {
			case *ast.SelectorExpr:
				value = fmt.Sprintf("%s.%s", t2.X, t2.Sel.Name)
			default:
				fmt.Println(reflect.TypeOf(t2))
			}
		case *ast.SelectorExpr:
			value = fmt.Sprintf("%s.%s", t.X, t.Sel.Name)
		}
		if forceParents {
			value = "(" + value + ")"
		}
		return value
	default:
		values := make([]string, len(fieldList.List))
		for index, field := range fieldList.List {
			// Tocar aqui para (val1,val2 string)
			switch ft := field.Type.(type) {
			case *ast.Ident:
				values[index] = ft.Name
			case *ast.FuncType:
				v := "func("
				for index, p2 := range ft.Params.List {
					fmt.Println(index)
					fmt.Println(p2.Names)

					switch ft2 := p2.Type.(type) {
					case *ast.Ellipsis:
						fmt.Println("---")
						elt := ft2.Elt.(*ast.InterfaceType)
						fmt.Printf("%#v\n", elt.Methods.NumFields())
						for _, m := range elt.Methods.List {

							fmt.Println(m.Names)
						}
						fmt.Println("---")
						if index < len(ft.Params.List)-1 {
							v += ","
						}
					}

				}
				values[index] = v
			}
			if ident, ok := field.Type.(*ast.Ident); ok {
				values[index] = ident.Name
			} else {
				fmt.Println(reflect.TypeOf(field.Type))
			}
		}
		return "(" + strings.Join(values, ",") + ")"
	}
	return ""
}
