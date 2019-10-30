package function

import (
	"github.com/wesovilabs/goa/parser"
	"go/ast"
)

// GetFunctions return the functions
func GetFunctions(packages map[string]*parser.Package) *Functions {
	functions := &Functions{}
	for _, pkg := range packages {
		for _, file := range pkg.Node().Files {
			searchFunctions(pkg.Path(), file, functions)
		}
	}
	return functions
}

func searchFunctions(pkg string, file *ast.File, functions *Functions) {
	for _, obj := range file.Scope.Objects {
		if decl, ok := obj.Decl.(*ast.FuncDecl); ok {
			path := buildPath(file, decl)
			functions.AddFunction(&Function{
				parent: file,
				decl:   decl,
				path:   path,
				pkg:    pkg,
			})
		}
		if obj.Kind != ast.Fun {
			continue
		}

	}
}
