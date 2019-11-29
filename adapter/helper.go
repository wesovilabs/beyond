package adapter

import (
	"fmt"
	"github.com/wesovilabs/beyond/joinpoint"
	"go/ast"
	"go/token"
	"strings"
)

func ensureImports(currentImports, requiredImports map[string]string, function *joinpoint.JoinPoint) {
	specs := make([]ast.Spec, 0)

	for path, name := range requiredImports {
		currentName := findImportName(currentImports, name, path)

		if _, ok := currentImports[path]; !ok {
			if spec := addImportSpec(function, currentName, path); spec != nil {
				specs = append(specs, spec)
			}
		}

		if path != "" {
			currentImports[path] = currentName
		}
	}

	updateImportSpec(function, specs)
}

func findImportName(imports map[string]string, name, path string) string {
	for value, name := range imports {
		if value == path {
			return name
		}
	}

	if name == "" {
		name = path[strings.LastIndex(path, "/")+1:]
	}

	return findAvailableImportName(imports, name)
}

func addImportSpec(function *joinpoint.JoinPoint, name, path string) *ast.ImportSpec {
	if path == "" {
		return nil
	}

	spec := &ast.ImportSpec{
		Name: ast.NewIdent(name),
		Path: &ast.BasicLit{
			Value: fmt.Sprintf("\"%s\"", path),
			Kind:  token.STRING,
		},
	}
	function.AddImportSpec(spec)

	return spec
}

func findAvailableImportName(imports map[string]string, name string) string {
	for _, n := range imports {
		if n == name {
			return findAvailableImportName(imports, fmt.Sprintf("_%s", name))
		}
	}

	return name
}

func updateImportSpec(function *joinpoint.JoinPoint, specs []ast.Spec) {
	specs = cleanUpImportSpec(function, specs)
	if len(specs) > 0 {
		function.AddImportSpecs(specs)
	}
}

func cleanUpImportSpec(function *joinpoint.JoinPoint, specs []ast.Spec) []ast.Spec {
	importSpecs := make([]ast.Spec, 0)

	for _, spec := range specs {
		importSpecs = append(importSpecs, spec)

		for _, decl := range function.FileDecls() {
			if genDecl, ok := decl.(*ast.GenDecl); ok {
				for _, genDeclSpec := range genDecl.Specs {
					if s, ok := genDeclSpec.(*ast.ImportSpec); ok {
						if s.Path == spec.(*ast.ImportSpec).Path {
							if len(importSpecs) == 0 {
								break
							}

							importSpecs = importSpecs[:len(importSpecs)-1]

							break
						}
					}
				}
			}
		}
	}

	return importSpecs
}
