package aspect

import (
	"fmt"
	"github.com/wesovilabs/goa/aspect/internal"
	"github.com/wesovilabs/goa/logger"
	"go/ast"
	"go/token"
	"regexp"
	"strconv"
	"strings"
)

type definitionKind int32

const (
	around definitionKind = iota
	before
	returning
	pkgSeparator = "/"
	apiPath      = "github.com/wesovilabs/goa/api/v2"
)

type definitionFunc struct {
	pkg    string
	name   string
	isType bool
	args   []interface{}
}

type Definition struct {
	pkgPath  string
	//function *definitionFunc
	kind     definitionKind
	regExp   *regexp.Regexp
}

func (def *Definition) update(rootPkg string, expr *ast.CallExpr, importSpecs []*ast.ImportSpec) {
	pattern := expr.Args[0].(*ast.BasicLit).Value
	regExp := internal.NormalizeExpression(pattern[1 : len(pattern)-1])
	if regExp == nil {
		logger.Errorf("invalid regExp: %s ", pattern)
		return
	}
	def.regExp = regExp
	def.function = &definitionFunc{}
	switch arg := expr.Args[1].(type) {
	case *ast.SelectorExpr:
		ident := arg.X.(*ast.Ident)
		def.pkgPath = pkgPathForType(ident.Name, importSpecs)
		def.function.name = arg.Sel.Name
	case *ast.Ident:
		if arg.Obj != nil && arg.Obj.Decl != nil {
			switch decl := arg.Obj.Decl.(type) {
			case *ast.AssignStmt:
				switch fun := decl.Rhs[0].(type) {
				case *ast.CallExpr:
					def.function.fromCallExpr(rootPkg, fun, importSpecs)
					return
				case *ast.UnaryExpr:
					def.function.fromUnaryExpr(rootPkg, fun)
				default:
					fmt.Println("Not found")
				}
			case *ast.ValueSpec:
				switch a:=decl.Type.(type) {

				case *ast.StarExpr:
					fmt.Println("star")
					switch c:=a.X.(type) {
					case *ast.Ident:
						def.function.name=c.Obj.Name
						def.function.pkg=rootPkg
					default:fmt.Println(c)
					}
				default:
					fmt.Println("not found")
				}
			default:
				fmt.Println("not found")
			}

		}
	case *ast.CallExpr:
		def.function.fromCallExpr(rootPkg, arg, importSpecs)
	case *ast.UnaryExpr:
		def.function.fromUnaryExpr(rootPkg, arg)
	default:
		fmt.Println("not found")
	}

}
func (fun *definitionFunc) fromUnaryExpr(rootPkg string, expr *ast.UnaryExpr) {
	fun.isType = true
	fun.pkg = rootPkg
	switch p := expr.X.(type) {
	case *ast.CompositeLit:
		switch cm := p.Type.(type) {
		case *ast.Ident:
			fun.name = cm.Obj.Name
			fun.args = make([]interface{}, len(p.Elts))
			for index, e := range p.Elts {
				switch ea := e.(type) {
				case *ast.BasicLit:
					if ea.Kind == token.INT {
						fun.args[index], _ = strconv.Atoi(ea.Value)
					} else {
						fun.args[index] = ea.Value[1 : len(ea.Value)-1]
					}
				default:
					fmt.Println("Not found")
				}

			}
		default:
			fmt.Println("unexpected")
		}
	default:
		fmt.Println("unexpected")
	}
}
func (fun *definitionFunc) fromCallExpr(rootPkg string, expr *ast.CallExpr, importSpecs []*ast.ImportSpec) {
	fun.args = make([]interface{}, len(expr.Args))
	for index, arg := range expr.Args {
		switch a := arg.(type) {
		case *ast.BasicLit:
			fun.args[index] = a.Value[1 : len(a.Value)-1]
		case *ast.Ident:
			fun.args[index] = a.Name
		default:
			fmt.Println(";(")
		}
	}
	switch f := expr.Fun.(type) {
	case *ast.SelectorExpr:
		switch t := f.X.(type) {
		case *ast.Ident:
			fun.pkg = pkgPathForType(t.Name, importSpecs)
		default:
			fmt.Println("Not found")
		}
		fun.name = f.Sel.Name
	case *ast.Ident:
		fun.name = f.Name
		fun.pkg = rootPkg
	default:
		fmt.Println("Not doun")
	}
}

func (d *Definition) Match(text string) bool {
	logger.Infof("matching \"%s\" with \"%s\"", d.regExp.String(), text)
	return d.regExp.MatchString(text)
}

func (a *Definition) Pkg() string {
	return a.pkgPath
}

func (d *Definition) Name() string {
	return d.function.name

}

type Definitions struct {
	items []*Definition
}

func (d *Definitions) List() []*Definition {
	return d.items
}

func (d *Definitions) Add(def *Definition) {
	d.items = append(d.items, def)
}

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
			definition.update(rootPkg, expr, importSpecs)
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
