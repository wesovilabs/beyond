package function

import (
	"github.com/wesovilabs/goa/logger"
	"go/ast"
)

func GetFunctions(packages map[string]*ast.Package) *Functions {
	functinos := &Functions{}
	for _, pkg := range packages {
		for _, file := range pkg.Files {
			searchFunctions(file, functinos)
		}
	}
	return functinos
}

func searchFunctions(file *ast.File, functions *Functions) {
	for _, obj := range file.Scope.Objects {
		switch decl := obj.Decl.(type) {
		case *ast.FuncDecl:
			path := buildPath(file, decl)
			functions.withFunction(&Function{
				parent: file,
				decl:   decl,
				path:   path,
			})
		default:
			logger.Info("not expected value")
		}
		if obj.Kind != ast.Fun {
			continue
		}

	}
}
