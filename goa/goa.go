package goa

import (
	"fmt"
	"github.com/wesovilabs/goa/inspector"
	"github.com/wesovilabs/goa/inspector/aspect"
	goaAST "github.com/wesovilabs/goa/inspector/ast"
	"github.com/wesovilabs/goa/writer"
	"go/ast"
	"sync"
)

var (
	once     sync.Once
	instance *goa
)

// Goa returns an instace of goa structure
func Goa() *goa {
	once.Do(func() {
		instance = &goa{
			functions: []*inspector.Function{},
			aspects:   aspect.Aspects{},
		}
	})
	return instance
}

type goa struct {
	imports   map[string]string
	functions []*inspector.Function
	aspects   aspect.Aspects
}

func (g *goa) run() {
	for _, aspect := range g.aspects {
		for _, function := range g.functions {
			if aspect.Match(function.Path()) {
				// TODO all the aspects applied to a function should be passes together.
				// We must  avoid more than a if-else  per function in generated code
				goaAST.Transform(g.imports, function, aspect)
			}
		}
	}

}

func (g *goa) normalize() {
	normalizedFunctions := make([]*inspector.Function, 0)
	for _, function := range g.functions {
		skip := false
		for _, aspect := range g.aspects {
			if function.Name() == aspect.Name() && function.Pkg() == aspect.Pkg() {
				skip = true
				break
			}
		}
		if function.Name() == "main" {
			skip = true
		}
		if function.Name() == "Goa" {
			skip = true
		}
		if !skip {
			normalizedFunctions = append(normalizedFunctions, function)
		}
	}
	g.functions = normalizedFunctions
}

func (g *goa) Execute(node *ast.File) {
	inspector := inspector.NewInspector(node)
	g.aspects = inspector.SearchRegisteredAspects()
	g.functions = inspector.SearchFunctions()
	g.imports = inspector.SearchImports()
	g.normalize()
	g.run()
	writer.Node(node, fmt.Sprintf(".goa/main.go"))
}
