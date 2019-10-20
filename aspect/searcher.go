package aspect

import (
	"github.com/wesovilabs/goa/logger"
	"go/ast"
	"regexp"
	"strings"
)

const (
	apiPath      = "github.com/wesovilabs/goa/api"
	aroundFn     = "WithAround"
	pkgSeparator = "/"
)

func GetAspects(packages map[string]*ast.Package) *Aspects {
	aspects := &Aspects{
		aroundList: make([]*AroundAspect, 0),
	}
	for _, pkg := range packages {
		for _, file := range pkg.Files {
			searchAspects(file, aspects)
		}
	}
	return aspects
}

func searchAspects(node *ast.File, aspects *Aspects) *Aspects {
	if name := findApiImport(node, apiPath); name != "" {
		funcDecl := findGoaFunction(node, name)
		if funcDecl != nil {
			takeAspectsFromFunction(funcDecl, node.Name.Name, aspects)
		}
	}
	return nil
}

func findRegisteredAspects(expr *ast.CallExpr, pkg string, aspects *Aspects) {
	if selExpr, ok := expr.Fun.(*ast.SelectorExpr); ok {
		if selExpr.Sel.Name == aroundFn {
			pattern := expr.Args[0].(*ast.BasicLit).Value
			regExp, err := regexp.Compile(pattern[1 : len(pattern)-1])
			if err != nil {
				logger.Errorf("invalid regExp: %s ", pattern)
				return
			}
			aspect := &AroundAspect{
				regExp: regExp,
			}
			switch arg := expr.Args[1].(type) {
			case *ast.SelectorExpr:
				ident := arg.X.(*ast.Ident)
				aspect.with(ident.Name, arg.Sel.Name)
			case *ast.Ident:
				aspect.with(pkg, arg.Name)
			}
			aspects.WithAround(aspect)
		}
		if callExpr, ok := selExpr.X.(*ast.CallExpr); ok {
			findRegisteredAspects(callExpr, pkg, aspects)
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

func findApiImport(file *ast.File, path string) string {
	for _, importSpec := range file.Imports {
		value := importSpec.Path.Value[1 : len(importSpec.Path.Value)-1]
		if path == value {
			if importSpec.Name != nil {
				return importSpec.Name.Name
			}
			lastIndex := strings.LastIndex(value, pkgSeparator)
			return value[lastIndex+1:]
		}
	}
	return ""
}

func takeAspectsFromFunction(funcDecl *ast.FuncDecl, pkgPath string, aspects *Aspects) {
	for _, stmt := range funcDecl.Body.List {
		if expr, ok := stmt.(*ast.ReturnStmt); ok {
			if callExpr, ok := expr.Results[0].(*ast.CallExpr); ok {
				findRegisteredAspects(callExpr, pkgPath, aspects)
			} else {
				logger.Infof("this format is not supported yed!")
			}
			return
		}
	}
}
