package adapter

import (
	"github.com/wesovilabs/beyond/advice"
	"github.com/wesovilabs/beyond/joinpoint"
)

var requiredImports = map[string]string{
	"github.com/wesovilabs/beyond/api/context": "beyondContext",
}

// Adapter function that create the ast for the intercepted function
func Adapter(joinPoint *joinpoint.JoinPoint, advices map[string]*advice.Advice) {
	file := joinPoint.Parent()
	funcDecl := adapterFuncDecl(joinPoint, advices)
	file.Decls = append(file.Decls, funcDecl)

	joinPoint.RenameToInternal()
}
