package internal

import (
	"fmt"
	"github.com/wesovilabs/goa/inspector"
	"go/ast"
	"go/token"
)

var requiredImportsPath = map[string]string{
	"context":                           "context",
	"github.com/wesovilabs/goa/goa":     "goa",
	"github.com/wesovilabs/goa/context": "goaContext",
}

// EnsureImports ensure that required imports are in the ast
func (a *ASTAdapter) EnsureImports() {
	specs := make([]ast.Spec, 0)
	for path, name := range requiredImportsPath {
		currentName := ensureImport(a.imports, name, path)
		if _, ok := a.imports[path]; !ok {
			if spec := addImportSpec(a.function, currentName, path); spec != nil {
				specs = append(specs, spec)
			}
		}
		a.imports[path] = currentName

	}

	updateASTWithImportSpecs(a.function, specs)
}

func findAvailableImportName(imports map[string]string, name string) string {
	for _, n := range imports {
		if n == name {
			return findAvailableImportName(imports, fmt.Sprintf("_%s", name))
		}
	}
	return name
}

func ensureImport(imports map[string]string, name, path string) string {
	for value, name := range imports {
		if value == path {
			return name
		}
	}
	return findAvailableImportName(imports, name)
}

func removeExistingImportSpecs(function *inspector.Function, specs []ast.Spec) []ast.Spec {
	importSpecs := make([]ast.Spec, 0)
	for _, spec := range specs {
		importSpecs = append(importSpecs, spec)
		for _, decl := range function.FileDecls() {
			switch genDecl := decl.(type) {
			case *ast.GenDecl:
				for _, genDeclSpec := range genDecl.Specs {
					switch s := genDeclSpec.(type) {
					case *ast.ImportSpec:
						if s.Path == spec.(*ast.ImportSpec).Path {
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

func updateASTWithImportSpecs(function *inspector.Function, specs []ast.Spec) {
	specs = removeExistingImportSpecs(function, specs)
	if len(specs) > 0 {
		importsSpecDecl := []ast.Decl{&ast.GenDecl{
			Tok:   token.IMPORT,
			Specs: removeExistingImportSpecs(function, specs),
		}}
		function.AddDeclsBefore(importsSpecDecl)
	}
}

func addImportSpec(function *inspector.Function, name, path string) *ast.ImportSpec {
	spec := &ast.ImportSpec{
		Name: ast.NewIdent(name),
		Path: &ast.BasicLit{
			Value: fmt.Sprintf("\"%s\"", path),
			Kind:  token.STRING,
		},
	}
	function.AddImport(spec)
	return spec
}
