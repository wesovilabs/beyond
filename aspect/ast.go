package aspect

import (
	"fmt"
	"github.com/wesovilabs/goa/aspect/internal"
	"go/ast"
	"strings"
)

const (
	around definitionKind = iota
	before
	returning
	pkgSeparator = "/"
	apiPath      = "github.com/wesovilabs/goa/api"
)

func GetDefinitions(rootPkg string, packages map[string]*ast.Package) *Definitions {
	defs := &Definitions{
		items: make([]*Definition, 0),
	}
	for _, pkg := range packages {
		for _, file := range pkg.Files {
			searchDefinitions(rootPkg, file, defs)
		}
	}
	return defs
}

func searchDefinitions(rootPkg string, node *ast.File, definitions *Definitions) {
	if funcDecl := containsDefinitions(node); funcDecl != nil {
		for _, stmt := range funcDecl.Body.List {
			if expr, ok := stmt.(*ast.ReturnStmt); ok {
				if callExpr, ok := expr.Results[0].(*ast.CallExpr); ok {
					addDefinition(rootPkg, callExpr, node.Name.Name, definitions, node.Imports)
				}
				return
			}
		}
	}
}

func containsDefinitions(file *ast.File) *ast.FuncDecl {
	for _, importSpec := range file.Imports {
		value := importSpec.Path.Value[1 : len(importSpec.Path.Value)-1]
		if apiPath == value {
			if importSpec.Name != nil {
				return findGoaFunction(file, importSpec.Name.Name)
			}
			lastIndex := strings.LastIndex(value, pkgSeparator)
			return findGoaFunction(file, value[lastIndex+1:])
		}
	}
	return nil
}

var aspectTypes = map[string]definitionKind{
	"WithBefore":    before,
	"WithReturning": returning,
	"WithAround":    around,
}

func addDefinition(rootPkg string, expr *ast.CallExpr, pkg string, definitions *Definitions, importSpecs []*ast.ImportSpec) {
	if selExpr, ok := expr.Fun.(*ast.SelectorExpr); ok {
		if kind, ok := aspectTypes[selExpr.Sel.Name]; ok {
			definition := &Definition{
				kind:    kind,
				pkgPath: rootPkg,
			}
			switch arg := expr.Args[0].(type) {
			case *ast.BasicLit:
				if len(arg.Value) < 2 {
					return
				}
				definition.regExp = internal.NormalizeExpression(arg.Value[1 : len(arg.Value)-1])
			}
			switch arg := expr.Args[1].(type) {
			case *ast.Ident:
				definition.name = arg.Name
			case *ast.SelectorExpr:
				definition.name = arg.Sel.Name
				switch x := arg.X.(type) {
				case *ast.Ident:
					definition.pkgPath = pkgPathForType(x.Name, importSpecs)
				}
			}
			definitions.Add(definition)

		}
		if callExpr, ok := selExpr.X.(*ast.CallExpr); ok {
			addDefinition(rootPkg, callExpr, pkg, definitions, importSpecs)
		}
	}
}

func findGoaFunction(file *ast.File, instanceName string) *ast.FuncDecl {
	for _, obj := range file.Scope.Objects {
		if obj.Kind != ast.Fun {
			continue
		}
		funcDecl := obj.Decl.(*ast.FuncDecl)
		if funcDecl.Type.Results == nil {
			continue
		}
		results := funcDecl.Type.Results.List
		if len(results) != 1 {
			continue
		}
		if expr, ok := results[0].Type.(*ast.StarExpr); ok {
			if expr, ok := expr.X.(*ast.SelectorExpr); ok {
				exprX, ok := expr.X.(*ast.Ident)
				if !ok {
					continue
				}
				if exprX.Name == instanceName && expr.Sel.Name == "Goa" {
					return funcDecl
				}
			}
		}
	}
	return nil
}

func pkgPathForType(name string, importSpecs []*ast.ImportSpec) string {
	value := ""
	for _, importSpec := range importSpecs {
		path := importSpec.Path.Value[1 : len(importSpec.Path.Value)-1]
		if importSpec.Name != nil && importSpec.Name.Name == name {
			return path
		}
		if strings.HasSuffix(path, fmt.Sprintf("/%s", name)) {
			value = path
		}
	}
	return value
}
