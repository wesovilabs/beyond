package joinpoint

import (
	"fmt"
	"github.com/wesovilabs/beyond/advice"
	"github.com/wesovilabs/beyond/parser"
	"go/ast"
	"regexp"
	"strings"
)

const mainPkg = "main"

// GetJoinPoints return the functions
func GetJoinPoints(rootPkg string, advices *advice.Advices,
	ignored []*regexp.Regexp, packages map[string]*parser.Package) *JoinPoints {
	functions := &JoinPoints{}

	for _, pkg := range packages {
		for _, file := range pkg.Node().Files {
			searchFunctions(rootPkg, pkg.Path(), pkg.Node().Name, file, advices, ignored, functions)
		}
	}

	return functions
}

func calculateImports(currentPkg string, imports []*ast.ImportSpec) map[string]string {
	paths := map[string]string{}

	for _, imp := range imports {
		path := imp.Path.Value

		if imp.Name != nil {
			name := imp.Name.String()
			paths[name] = path[1 : len(path)-1]

			continue
		}

		name := path[strings.LastIndex(path, "/")+1 : len(path)-1]
		if name == currentPkg {
			paths["_"+name] = path[1 : len(path)-1]
		} else {
			paths[name] = path[1 : len(path)-1]
		}
	}

	return paths
}

// This function must check if really returns an Aspect
func isAspectFunction(decl *ast.FuncDecl) bool {
	results := decl.Type.Results
	if results != nil && len(results.List) == 1 {
		if sel, ok := results.List[0].Type.(*ast.SelectorExpr); ok {
			if sel.Sel.Name == "Around" || sel.Sel.Name == "Returning" || sel.Sel.Name == "Before" {
				return true
			}
		}
	}

	if decl.Name.Name == "Before" || decl.Name.Name == "Returning" {
		return true
	}

	return false
}

func objectType(decl *ast.FuncDecl) string {
	objType := ""

	if decl.Recv != nil {
		switch p := decl.Recv.List[0].Type.(type) {
		case *ast.StarExpr:
			if id, ok := p.X.(*ast.Ident); ok {
				objType = fmt.Sprintf("*%s", id.String())
			}
		case *ast.Ident:
			objType = p.String()
		}
	}

	return objType
}

func searchFunctions(rootPkg string, parentPath, pkg string, file *ast.File, advices *advice.Advices,
	ignored []*regexp.Regexp, functions *JoinPoints) {
	imports := calculateImports(pkg, file.Imports)

	for _, obj := range file.Decls {
		if decl, ok := obj.(*ast.FuncDecl); ok {
			if isAspectFunction(decl) {
				continue
			}

			objType := objectType(decl)

			path := buildPath(rootPkg, parentPath, objType, decl, imports)
			jp := &JoinPoint{
				parent:  file,
				decl:    decl,
				path:    path,
				pkg:     pkg,
				pkgPath: fmt.Sprintf("%s/%s", rootPkg, parentPath),
			}

			if pkg == mainPkg {
				index := strings.Index(path, ".")
				jp.path = fmt.Sprintf("%s.%s", mainPkg, path[index+1:])
			}

			if jp.canBeIntercepted(ignored) {
				jp.findMatches(advices)

				if len(jp.Advices()) > 0 {
					functions.AddJoinPoint(jp)
				}
			}
		}
	}
}
