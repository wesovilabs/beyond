package ast

import (
	"fmt"
	"github.com/wesovilabs/goa/function"
	"go/ast"
	"go/token"
)

func findAvailableImportName(imports map[string]string, name string) string {
	for _, n := range imports {
		if n == name {
			return findAvailableImportName(imports, fmt.Sprintf("_%s", name))
		}
	}
	return name
}

func ensureImports(currentImports, requiredImports map[string]string, function *function.Function) {
	specs := make([]ast.Spec, 0)
	for path, name := range requiredImports {
		currentName := ensureImport(currentImports, name, path)
		if _, ok := currentImports[path]; !ok {
			if spec := addImportSpec(function, currentName, path); spec != nil {
				specs = append(specs, spec)
			}
		}
		currentImports[path] = currentName
	}
	updateImportSpec(function, specs)
}

func ensureImport(imports map[string]string, name, path string) string {
	for value, name := range imports {
		if value == path {
			return name
		}
	}
	return findAvailableImportName(imports, name)
}

func addImportSpec(function *function.Function, name, path string) *ast.ImportSpec {
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

func updateImportSpec(function *function.Function, specs []ast.Spec) {
	specs = cleanUpImportSpec(function, specs)
	if len(specs) > 0 {
		function.AddImportSpecs(specs)
	}
}

func cleanUpImportSpec(function *function.Function, specs []ast.Spec) []ast.Spec {
	importSpecs := make([]ast.Spec, 0)
	for _, spec := range specs {
		importSpecs = append(importSpecs, spec)
		for _, decl := range function.FileDecls() {
			if genDecl, ok := decl.(*ast.GenDecl); ok {
				for _, genDeclSpec := range genDecl.Specs {
					if s, ok := genDeclSpec.(*ast.ImportSpec); ok {
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

func updateValue(field *ast.Field) string {
	switch f := field.Type.(type) {
	case *ast.Ident:
		return fmt.Sprintf("(%s)", f.Name)
	case *ast.StarExpr:
		return fmt.Sprintf("(*%s.%s)", f.X.(*ast.SelectorExpr).X.(*ast.Ident).String(), f.X.(*ast.SelectorExpr).Sel.String())
	}
	return ""

}
