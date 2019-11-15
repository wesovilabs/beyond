package adapter

import (
	"github.com/wesovilabs/goa/advice"
	"github.com/wesovilabs/goa/joinpoint"
)

var requiredImports = map[string]string{
	"context":                               "context",
	"github.com/wesovilabs/goa/api/context": "goaContext",
}

// Adapter function that create the ast for the intercepted function
func Adapter(joinPoint *joinpoint.JoinPoint, advices map[string]*advice.Advice) {
	file := joinPoint.Parent()
	funcDecl := adapterFuncDecl(joinPoint, advices)
	file.Decls = append(file.Decls, funcDecl)

	joinPoint.RenameToInternal()
}
