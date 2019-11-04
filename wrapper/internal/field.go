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

	for _, arg := range fields {
		for _, argName := range arg.Names {
			params = append(params, &FieldDef{
				name: argName.String(),
				kind: arg.Type,
			})
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
