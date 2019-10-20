package parser

import (
	"github.com/wesovilabs/goa/logger"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func NewGoaPackage(path string) (*ast.Package, []string) {
	if _, err := os.Stat(path); err != nil {
		return nil, nil
	}
	logger.Infof("inspecting directory %s", path)
	fileSet := token.NewFileSet()
	if packages, err := parser.ParseDir(fileSet, path, nil, parser.ParseComments); err == nil {
		pkg := applyPkgFilters(packages, excludeTestPackages)
		imports := make([]string, 0)
		index := 0

		for _, file := range pkg.Files {
			for _, importSpec := range file.Imports {
				importPath := importSpec.Path.Value
				importPath = importPath[1 : len(importPath)-1]
				imports = append(imports, importPath)
				index++
			}
		}
		return pkg, imports

	}
	return nil, nil
}
