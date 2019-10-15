package ast

import (
	"github.com/wesovilabs/goa/inspector"
	"github.com/wesovilabs/goa/inspector/aspect"
	"github.com/wesovilabs/goa/inspector/ast/internal"
)

// Transform update ast structure with the given data
func Transform(imports map[string]string, function *inspector.Function, aspect *aspect.Aspect) {
	adapter := internal.New(function, aspect, imports)
	adapter.EnsureImports()
	adapter.UpdateFunction()
}
