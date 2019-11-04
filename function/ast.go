package function

import (
	"fmt"
	"github.com/wesovilabs/goa/parser"
	"go/ast"
	"strings"
)

// GetFunctions return the functions
func GetFunctions(packages map[string]*parser.Package) *Functions {
	functions := &Functions{}
	for _, pkg := range packages {
		for _, file := range pkg.Node().Files {
			searchFunctions(pkg.Path(), file, functions)
		}
	}
	return functions
}

func calculateImports(imports []*ast.ImportSpec) map[string]string {
	paths := make(map[string]string, 0)
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

func searchFunctions(pkg string, file *ast.File, functions *Functions) {
	imports := calculateImports(file.Imports)
	for _, obj := range file.Decls {
		switch decl := obj.(type) {
		case *ast.FuncDecl:
			objType := ""
			if decl.Name.Name == "j" {
				fmt.Println("stop")
			}
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
			path := buildPath(pkg, objType, decl, imports)
			functions.AddFunction(&Function{
				parent: file,
				decl:   decl,
				path:   path,
				pkg:    pkg,
			})
		default:
			fmt.Println(decl)
			continue
		}

	}
}
