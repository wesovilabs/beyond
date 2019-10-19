package inspector

import (
	"github.com/wesovilabs/goa/inspector/aspect"
	"github.com/wesovilabs/goa/logger"
	"go/ast"
	"strings"
)

// Inspector struct for a inspecting a file
type Inspector struct {
	node *ast.File
}

// NewInspector creates an instance of struct Inspector
func NewInspector(node *ast.File) *Inspector {
	return &Inspector{
		node: node,
	}
}

// SearchFunctions returns the list of functions
func (i *Inspector) SearchFunctions() []*Function {
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
func (i *Inspector) SearchImports() map[string]string {
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
func (i *Inspector) SearchRegisteredAspects() []*aspect.Aspect {
	aspects := make([]*aspect.Aspect, 0)
	for _, object := range i.node.Scope.Objects {
		if isFun(object) && object.Name == "Goa" {
			logger.Infof("Goa function was found!")
			decl := object.Decl.(*ast.FuncDecl)
			aspectInspector := &AspectInspector{decl}
			logger.Info("Inspecting Goa function")
			aspects = append(aspects, aspectInspector.TakeAspects(i.node.Name.Name)...)
		}
	}
	return aspects
}

func isFun(obj *ast.Object) bool {
	return obj.Kind == ast.Fun
}
