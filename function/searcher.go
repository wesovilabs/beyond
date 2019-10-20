package function

import (
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
		if obj.Kind != ast.Fun {
			continue
		}
		funcDecl := obj.Decl.(*ast.FuncDecl)
		path := buildPath(file, funcDecl)
		functions.withFunction(&Function{
			parent: file,
			decl:   funcDecl,
			path:   path,
		})
	}
}
