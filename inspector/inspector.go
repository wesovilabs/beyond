package inspector

import (
	"github.com/wesovilabs/goa/inspector/aspect"
	"go/ast"
	"strings"
)

type inspector struct {
	node *ast.File
}

// NewInspector creates an instance of struct inspector
func NewInspector(node *ast.File) *inspector {
	return &inspector{
		node: node,
	}
}

// SearchFunctions returns the list of functions
func (i *inspector) SearchFunctions() []*Function {
	functions := make([]*Function, 0)
	for _, object := range i.node.Scope.Objects {
		if isFun(object) {
			funcDecl := object.Decl.(*ast.FuncDecl)
			functions = append(functions, newFunction(i.node, funcDecl))
		}
	}
	return functions
}

// SearchImports return a dictionary key,value with the existing imports in the ast
func (i *inspector) SearchImports() map[string]string {
	imports := make(map[string]string)
	for _, im := range i.node.Imports {
		if im.Name != nil {
			imports[im.Path.Value] = im.Name.Name
			continue
		}
		parts := strings.Split(im.Path.Value[1:len(im.Path.Value)-1], "/")
		imports[im.Path.Value[1:len(im.Path.Value)-1]] = parts[len(parts)-1]
	}
	return imports
}

// SearchRegisteredAspects return the list of registere aspects
func (i *inspector) SearchRegisteredAspects() []*aspect.Aspect {
	aspects := make([]*aspect.Aspect, 0)
	for _, object := range i.node.Scope.Objects {
		if isFun(object) {
			decl := object.Decl.(*ast.FuncDecl)
			aspectInspector := &AspectInspector{decl}
			aspects = append(aspects, aspectInspector.TakeAspects(i.node.Name.Name)...)
		}
	}
	return aspects
}

func isFun(obj *ast.Object) bool {
	return obj.Kind == ast.Fun
}
