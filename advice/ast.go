package advice

import (
	"fmt"
	"github.com/wesovilabs/goa/advice/internal"
	"github.com/wesovilabs/goa/logger"
	"github.com/wesovilabs/goa/parser"
	"go/ast"
	"go/token"
	"reflect"
	"strings"
)

const (
	around adviceType = iota
	before
	returning
	pkgSeparator = "/"
	apiPath      = "github.com/wesovilabs/goa/api"
)

// GetAdvices return the list of definitions (aspects)
func GetAdvices(rootPkg string, packages map[string]*parser.Package) *Advices {
	defs := &Advices{
		items: make([]*Advice, 0),
	}

	for _, pkg := range packages {
		for _, file := range pkg.Node().Files {
			searchAdvices(rootPkg, file, defs)
		}
	}

	return defs
}

func searchAdvices(rootPkg string, node *ast.File, definitions *Advices) {
	if funcDecl := containsAdvices(node); funcDecl != nil {
		for _, stmt := range funcDecl.Body.List {
			if expr, ok := stmt.(*ast.ReturnStmt); ok {
				if callExpr, ok := expr.Results[0].(*ast.CallExpr); ok {
					addAdvice(rootPkg, callExpr, definitions, node.Imports)
				}

				return
			}
		}
	}
}

func containsAdvices(file *ast.File) *ast.FuncDecl {
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

var aspectTypes = map[string]adviceType{
	"WithBefore":    before,
	"WithReturning": returning,
	"WithAround":    around,
}

func selectorToString(sel *ast.SelectorExpr) string {
	switch x := sel.X.(type) {
	case *ast.Ident:
		return fmt.Sprintf("%s.%s", x, sel.Sel.Name)
	default:
		logger.Error("unsupported type")

		return fmt.Sprintf("%s.%s", "?", sel.Sel.Name)
	}
}
func compositeToString(c *ast.CompositeLit) string {
	switch x := c.Type.(type) {
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s{}", selectorToString(x))
	default:
		logger.Error("unsupported type")
		return ""
	}
}

func unaryToString(c *ast.UnaryExpr) string {
	prefix := ""
	if c.Op == token.AND {
		prefix = "&"
	}

	switch x := c.X.(type) {
	case *ast.CompositeLit:
		return fmt.Sprintf("%s%s", prefix, compositeToString(x))
	default:
		logger.Error("unsupported type")
		return ""
	}
}

func addAdviceCallExpr(arg *ast.CallExpr, definition *Advice, importSpecs []*ast.ImportSpec) {
	args := make([]string, 0)

	for _, ar := range arg.Args {
		switch a := ar.(type) {
		case *ast.BasicLit:
			args = append(args, a.Value)
		case *ast.SelectorExpr:
			args = append(args, selectorToString(a))
		case *ast.CompositeLit:
			args = append(args, compositeToString(a))
		case *ast.UnaryExpr:
			args = append(args, unaryToString(a))
		default:
			fmt.Println(reflect.TypeOf(a))
		}
	}

	funcName := ""

	switch f := arg.Fun.(type) {
	case *ast.SelectorExpr:
		funcName = f.Sel.Name

		if x, ok := f.X.(*ast.Ident); ok {
			definition.pkg = pkgPathForType(x.Name, importSpecs)
		}
	default:
		fmt.Println(reflect.TypeOf(f))
	}

	definition.name = fmt.Sprintf("%s(%s)", funcName, strings.Join(args, ","))
}

func takeAdvice(expr ast.Expr, definition *Advice, importSpecs []*ast.ImportSpec) {
	switch arg := expr.(type) {
	case *ast.Ident:
		definition.name = arg.Name
	case *ast.SelectorExpr:
		definition.name = arg.Sel.Name
		if x, ok := arg.X.(*ast.Ident); ok {
			definition.pkg = pkgPathForType(x.Name, importSpecs)
		}
	case *ast.BasicLit:
		definition.regExp = internal.NormalizeExpression(arg.Value[1 : len(arg.Value)-1])
	case *ast.CallExpr:
		addAdviceCallExpr(arg, definition, importSpecs)
	default:
		fmt.Println(reflect.TypeOf(arg))
	}
}

func addAdvice(rootPkg string, expr *ast.CallExpr, definitions *Advices,
	importSpecs []*ast.ImportSpec) {
	if selExpr, ok := expr.Fun.(*ast.SelectorExpr); ok {
		if kind, ok := aspectTypes[selExpr.Sel.Name]; ok {
			definition := &Advice{
				kind: kind,
				pkg:  rootPkg,
			}
			takeAdvice(expr.Args[0], definition, importSpecs)

			if arg, ok := expr.Args[1].(*ast.BasicLit); ok {
				if len(arg.Value) < 2 {
					return
				}

				definition.regExp = internal.NormalizeExpression(arg.Value[1 : len(arg.Value)-1])
			}

			fmt.Printf("%#v", definition)
			definitions.Add(definition)
		}

		if callExpr, ok := selExpr.X.(*ast.CallExpr); ok {
			addAdvice(rootPkg, callExpr, definitions, importSpecs)
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
