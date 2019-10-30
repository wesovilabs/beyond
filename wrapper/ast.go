package wrapper

import (
	"github.com/wesovilabs/goa/aspect"
	"github.com/wesovilabs/goa/function"
)

var requiredImports = map[string]string{
	"context":                               "context",
	"github.com/wesovilabs/goa/api/context": "goaContext",
}

// Wrap function that create the ast for the intercepted function
func Wrap(function *function.Function, definitions map[string]*aspect.Definition) {
	file := function.Parent()

	funcDecl := wrapperFuncDecl(function, definitions)
	file.Decls = append(file.Decls, funcDecl)
	function.RenameToInternal()
}
