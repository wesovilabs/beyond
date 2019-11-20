package internal

import (
	"fmt"
	"go/ast"
)

// FieldDef struct
type FieldDef struct {
	Name string
	Kind ast.Expr
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
				Name: paramName,
			}
			/**
			if ell, ok := arg.Type.(*ast.Ellipsis); ok {
				fd.kind = &ast.ArrayType{
					Elt: ell.Elt,
				}
			} else {
				fd.kind = arg.Type
			}
			**/
			fd.Kind = arg.Type
			params = append(params, fd)
		}
	}

	return params
}

// Results return the results
func Results(fields []*ast.Field) []*FieldDef {
	results := make([]*FieldDef, 0)
	index := 0

	for i := range fields {
		arg := fields[i]
		if arg.Names != nil {
			for range arg.Names {
				results = append(results, &FieldDef{
					Name: fmt.Sprintf("result%v", index),
					Kind: arg.Type,
				})
				index++
			}
		} else {
			results = append(results, &FieldDef{
				Name: fmt.Sprintf("result%v", index),
				Kind: arg.Type,
			})
			index++
		}
	}

	return results
}
