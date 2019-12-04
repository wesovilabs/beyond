package advice

import (
	"fmt"
	"github.com/wesovilabs/beyond/advice/internal"
	"github.com/wesovilabs/beyond/logger"
	"github.com/wesovilabs/beyond/parser"
	"go/ast"
	"go/token"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const (
	around adviceType = iota
	before
	returning
	pkgSeparator = "/"
	apiPath      = "github.com/wesovilabs/beyond/api"
)

func unsupportedType(name string) string {
	logger.Errorf("unsupported type %s", name)
	return ""
}

// GetExcludePaths returns the paths to be excluded
func GetExcludePaths(packages map[string]*parser.Package) []*regexp.Regexp {
	paths := make([]*regexp.Regexp, 0)

	for _, pkgParser := range packages {
		if pkgParser.Node().Name == "main" {
			for _, file := range pkgParser.Node().Files {
				paths = append(paths, searchExcludePaths(file)...)
			}
		}
	}

	return paths
}

// GetAdvices return the list of advices (aspects)
func GetAdvices(packages map[string]*parser.Package) *Advices {
	defs := &Advices{
		items: make([]*Advice, 0),
	}

	for _, pkgParser := range packages {
		if pkgParser.Node().Name == "main" {
			for _, file := range pkgParser.Node().Files {
				searchAdvices(file, defs)
			}
		}
	}

	return defs
}

func searchAdvices(node *ast.File, advices *Advices) {
	if funcDecl := containsAdvices(node); funcDecl != nil {
		for _, stmt := range funcDecl.Body.List {
			if expr, ok := stmt.(*ast.ReturnStmt); ok {
				if callExpr, ok := expr.Results[0].(*ast.CallExpr); ok {
					addAdvice(callExpr, advices, node.Imports)
				}

				return
			}
		}
	}
}

func searchExcludePaths(node *ast.File) []*regexp.Regexp {
	paths := make([]*regexp.Regexp, 0)

	if funcDecl := containsAdvices(node); funcDecl != nil {
		for _, stmt := range funcDecl.Body.List {
			if expr, ok := stmt.(*ast.ReturnStmt); ok {
				if callExpr, ok := expr.Results[0].(*ast.CallExpr); ok {
					if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
						if selExpr.Sel.Name == "Ignore" {
							for i := range callExpr.Args {
								arg := callExpr.Args[i]
								if basic, ok := arg.(*ast.BasicLit); ok {
									val := basic.Value
									regExp := regexp.MustCompile(val[2 : len(val)-2])
									paths = append(paths, regExp)
								}
							}
						}
					}
				}

				return paths
			}
		}
	}

	return paths
}

func containsAdvices(file *ast.File) *ast.FuncDecl {
	for _, importSpec := range file.Imports {
		value := importSpec.Path.Value[1 : len(importSpec.Path.Value)-1]
		if apiPath == value {
			if importSpec.Name != nil {
				return findBeyondFunction(file, importSpec.Name.Name)
			}

			lastIndex := strings.LastIndex(value, pkgSeparator)

			return findBeyondFunction(file, value[lastIndex+1:])
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
		return unsupportedType(reflect.TypeOf(x).String())
	}
}
func compositeToString(c *ast.CompositeLit) string {
	switch x := c.Type.(type) {
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s{}", selectorToString(x))
	default:
		return unsupportedType(reflect.TypeOf(x).String())
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
		return unsupportedType(reflect.TypeOf(x).String())
	}
}
func adviceCallText(ar ast.Expr) string {
	var argText string
	switch a := ar.(type) {
	case *ast.BasicLit:
		argText = a.Value
	case *ast.SelectorExpr:
		argText = selectorToString(a)
	case *ast.CompositeLit:
		argText = compositeToString(a)
	case *ast.UnaryExpr:
		argText = unaryToString(a)
	case *ast.Ident:
		argText = a.Name
	case *ast.CallExpr:
		args := make([]string, len(a.Args))
		for i := range a.Args {
			args[i] = adviceCallText(a.Args[i])
		}

		argText = fmt.Sprintf("%s(%s)", adviceCallText(a.Fun), strings.Join(args, ","))
	case *ast.BinaryExpr:
		argText = fmt.Sprintf("%s %s %s",adviceCallText(a.X),a.Op.String(),adviceCallText(a.Y))
	default:
		argText = unsupportedType(reflect.TypeOf(a).String())
	}

	return argText
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func addAdviceCallExpr(arg *ast.CallExpr, importSpecs []*ast.ImportSpec, invocation *adviceInvocation) {
	invocationArgs := make([]*adviceInvocationArg, 0)

	for _, ar := range arg.Args {
		argText := ""
		switch ar.(type) {
		case *ast.CallExpr:
			argText = adviceCallText(ar)

			invocationArgs = append(invocationArgs, &adviceInvocationArg{
				val: argText,
			})
		default:
			argText = adviceCallText(ar)
			if isNumeric(argText) {
				invocationArgs = append(invocationArgs, &adviceInvocationArg{
					val: argText,
				})
			} else {
				fmt.Println(argText)
				items := strings.Split(argText, ".")
				if len(items) == 2 {
					isPointer := false

					if items[0][0] == '&' {
						items[0] = items[0][1:]
						isPointer = true
					}

					pkgPath := pkgPathForType(items[0], importSpecs)
					invocation.addImport(pkgPath)

					invocationArgs = append(invocationArgs, &adviceInvocationArg{
						pkg:     pkgPath,
						val:     items[1],
						pointer: isPointer,
					})
				} else {

					invocationArgs = append(invocationArgs, &adviceInvocationArg{
						val: argText,
					})
				}
			}
		}
	}

	switch f := arg.Fun.(type) {
	case *ast.SelectorExpr:
		invocation.function = f.Sel.Name
		if x, ok := f.X.(*ast.Ident); ok {
			invocation.pkg = pkgPathForType(x.Name, importSpecs)
		}
	default:
		unsupportedType(reflect.TypeOf(f).String())
	}

	invocation.args = invocationArgs
}

func takeAdvice(expr ast.Expr, advice *Advice, importSpecs []*ast.ImportSpec) {
	invocation := &adviceInvocation{}
	switch arg := expr.(type) {
	case *ast.Ident:
		invocation.function = arg.Name
	case *ast.SelectorExpr:
		invocation.function = arg.Sel.Name

		if x, ok := arg.X.(*ast.Ident); ok {
			pkgPath := pkgPathForType(x.Name, importSpecs)
			invocation.addImport(pkgPath)
			invocation.pkg = pkgPath
		}
	case *ast.CallExpr:
		addAdviceCallExpr(arg, importSpecs, invocation)
		invocation.isCall = true
	default:
		unsupportedType(reflect.TypeOf(arg).String())
	}

	advice.call = invocation
}

func addAdvice(expr *ast.CallExpr, advices *Advices,
	importSpecs []*ast.ImportSpec) {
	if selExpr, ok := expr.Fun.(*ast.SelectorExpr); ok {
		if kind, ok := aspectTypes[selExpr.Sel.Name]; ok {
			advice := &Advice{
				kind: kind,
			}
			takeAdvice(expr.Args[0], advice, importSpecs)

			if len(advice.call.function) > 0 {
				if unicode.IsUpper(rune(advice.call.function[0])) {
					if arg, ok := expr.Args[1].(*ast.BasicLit); ok {
						if len(arg.Value) < 2 {
							return
						}

						advice.regExp = internal.NormalizePointcut(arg.Value[1 : len(arg.Value)-1])
					}

					advices.Add(advice)
				}
			}
		}

		if callExpr, ok := selExpr.X.(*ast.CallExpr); ok {
			addAdvice(callExpr, advices, importSpecs)
		}
	}
}

func findBeyondFunction(file *ast.File, instanceName string) *ast.FuncDecl {
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

				if exprX.Name == instanceName && expr.Sel.Name == "Beyond" {
					return funcDecl
				}
			}
		}
	}

	return nil
}

func getName(spec *ast.ImportSpec) string {
	if spec.Name != nil && spec.Name.Name != "" {
		return spec.Name.Name
	}

	path := spec.Path.Value[1 : len(spec.Path.Value)-1]
	lastIndex := strings.LastIndex(path, "/")

	return path[lastIndex+1:]
}
func pkgPathForType(name string, importSpecs []*ast.ImportSpec) string {
	value := ""

	for _, importSpec := range importSpecs {
		path := importSpec.Path.Value[1 : len(importSpec.Path.Value)-1]
		importSpecName := getName(importSpec)

		if importSpecName == name {
			return path
		}

		if strings.HasSuffix(path, fmt.Sprintf("/%s", name)) {
			value = path
		}
	}

	return value
}
