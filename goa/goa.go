package goa

import (
	"github.com/wesovilabs/goa/inspector"
	"github.com/wesovilabs/goa/inspector/aspect"
	goaAST "github.com/wesovilabs/goa/inspector/ast"
	"github.com/wesovilabs/goa/logger"
	"github.com/wesovilabs/goa/writer"
	"go/ast"
	"os"
	"path/filepath"
	"sync"
)

var (
	once     sync.Once
	instance *Goa
)

// Goa struct
type Goa struct {
	goa *goa
}

// Execute executes goa application
func (g *Goa) Execute(pkgPath,fileName string,node *ast.File) error {
	return g.goa.Execute(pkgPath,fileName,node)
}

// Init returns an instance of goa structure
func Init() *Goa {
	logger.Infof("Initializing goa")
	once.Do(func() {
		instance = &Goa{
			goa: &goa{
				functions: []*inspector.Function{},
				aspects:   aspect.Aspects{},
			},
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
				logger.Info("matched!")
				// TODO all the aspects applied to a function should be passes together.
				// We must  avoid more than a if-else  per function in generated code
				goaAST.Transform(g.imports, function, aspect)
			} else {
				logger.Info("no matched!")
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
		if function.Name() == "Init" {
			skip = true
		}
		if !skip {
			normalizedFunctions = append(normalizedFunctions, function)
		}
	}
	g.functions = normalizedFunctions
}

func (g *goa) Execute(pkgPath,fileName string,node *ast.File) error {
	inspector := inspector.NewInspector(node)
	g.aspects = inspector.SearchRegisteredAspects()
	logger.Infof("Registered aspects: %v", len(g.aspects))
	g.functions = inspector.SearchFunctions()
	g.imports = inspector.SearchImports()
	g.normalize()
	logger.Infof(`%s`, node.Name.String())
	for k, v := range g.imports {
		logger.Infof(`[import  ] %s "%s`, k, v)
	}
	for _, f := range g.functions {
		logger.Infof(`[function] %s.%s => %s`, f.Pkg(), f.Name(), f.Path())
	}
	for _, a := range g.aspects {
		logger.Infof(`[aspect  ] %s.%s`, a.Pkg(), a.Name())
	}
	g.run()
	if err := os.MkdirAll(filepath.Join("temporal","main-controller",pkgPath), os.ModePerm); err != nil {
		logger.Errorf("Error creating path %s",filepath.Join(".goa",pkgPath))
	}
	return writer.Node(node, filepath.Join("temporal","main-controller",pkgPath,fileName))
}
