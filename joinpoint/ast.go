package joinpoint

import (
	"fmt"
	"github.com/wesovilabs/goa/parser"
	"go/ast"
	"strings"
)

// GetJoinPoints return the functions
func GetJoinPoints(rootPkg string, packages map[string]*parser.Package) *JoinPoints {
	functions := &JoinPoints{}

	for pkgPath, pkg := range packages {
		for _, file := range pkg.Node().Files {
			fmt.Println("---- " + pkgPath + "  " + pkg.Path())
			searchFunctions(rootPkg, pkg.Path(), pkg.Node().Name, file, functions)
		}
	}

	return functions
}

func calculateImports(imports []*ast.ImportSpec) map[string]string {
	paths := map[string]string{}

	for _, imp := range imports {
		path := imp.Path.Value

		if imp.Name != nil {
			name := imp.Name.String()
			paths[name] = path[1 : len(path)-1]

			continue
		}

		name := path[strings.LastIndex(path, "/")+1 : len(path)-1]
		paths[name] = path[1 : len(path)-1]
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

func searchFunctions(rootPkg string, parentPath, pkg string, file *ast.File, functions *JoinPoints) {
	imports := calculateImports(file.Imports)

	for _, obj := range file.Decls {
		if decl, ok := obj.(*ast.FuncDecl); ok {
			if isAspectFunction(decl) {
				continue
			}

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

			path := buildPath(rootPkg, parentPath, objType, decl, imports)

			if pkg == "main" {
				index := strings.Index(path, ".")
				path = fmt.Sprintf("%s.%s", "main", path[index+1:])
			}

			functions.AddJoinPoint(&JoinPoint{
				parent: file,
				decl:   decl,
				path:   path,
				pkg:    pkg,
			})
		}
	}
}
