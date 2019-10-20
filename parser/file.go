package parser

import (
	"go/ast"
)

type jjj struct {
	path    string
	node    *ast.File
	imports []string
}

func newGoaFile(node *ast.File) []string {
	imports := make([]string, len(node.Imports))
	for _, importSpec := range node.Imports {
		importPath := importSpec.Path.Value
		imports = append(imports, importPath[1:len(importPath)-1])
	}
	return imports
}
