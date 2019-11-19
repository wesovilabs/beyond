package internal

import (
	"fmt"
	"go/ast"
)

// FieldDef struct
type FieldDef struct {
	name string
	kind ast.Expr
}

// Params return the params
func Params(fields []*ast.Field) []*FieldDef {
	params := make([]*FieldDef, 0)
	annonymousCounter := 0

	for _, arg := range fields {
		for _, argName := range arg.Names {
			paramName := argName.String()
			if argName.String() == "_" {
				paramName = fmt.Sprintf("annonymous%v", annonymousCounter)
				annonymousCounter++
			}

			fd := &FieldDef{
				name: paramName,
			}

			if ell, ok := arg.Type.(*ast.Ellipsis); ok {
				fd.kind = &ast.ArrayType{
					Elt: ell.Elt,
				}
			} else {
				fd.kind = arg.Type
			}

			params = append(params, fd)
		}
	}

	return params
}

// Results return the results
func Results(fields []*ast.Field) []*FieldDef {
	results := make([]*FieldDef, 0)
	for index, arg := range fields {
		results = append(results, &FieldDef{
			name: fmt.Sprintf("result%v", index),
			kind: arg.Type,
		})
	}

	return results
}
