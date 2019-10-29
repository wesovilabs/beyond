package internal

import (
	"github.com/wesovilabs/goa/aspect"
	"github.com/wesovilabs/goa/function"
	"github.com/wesovilabs/goa/imports"
	goaAST "github.com/wesovilabs/goa/internal/ast"
	"github.com/wesovilabs/goa/internal/writer"
	"github.com/wesovilabs/goa/logger"
	"github.com/wesovilabs/goa/matcher"
	"github.com/wesovilabs/goa/parser"
	"github.com/wesovilabs/goa/wrapper"
	"os"
	"path/filepath"
)

type goa struct {
	functions   *function.Functions
	definitions *aspect.Definitions
}

func (g *goa) cleanInvalidFunctions() {
	output := &function.Functions{}

	for _, f := range g.functions.List() {
		valid := true
		for _, d := range g.definitions.List() {
			if (d.Name() == f.Name() && d.Pkg() == f.Pkg()) || f.Name() == "main" || f.Name() == "Goa" {
				valid = false
				continue
			}
		}
		if valid {
			output.WithFunction(f)
		}
	}
	g.functions = output
}

func Run(rootPkg string, packages map[string]*parser.Package, outputDir string) {
	goa := &goa{}
	goa.definitions = aspect.GetDefinitions(rootPkg, packages)
	goa.functions = function.GetFunctions(rootPkg, packages)
	goa.cleanInvalidFunctions()
	for _, f := range goa.functions.List() {
		logger.Infof(`[function] %s.%s => %s`, f.Pkg(), f.Name(), f.Path())
	}
	for _, a := range goa.definitions.List() {
		logger.Infof(`[ aspect ] %s.%s`, a.Pkg(), a.Name())
	}
	matches := matcher.FindMatches(goa.functions, goa.definitions)
	for _, match := range matches {
		logger.Infof("[ match  ] %s", match.Function.Name())
		for _, d := range match.Definitions {
			logger.Infof("   - %s", d.Name())
		}
		wrapper.Wrap(match.Function, match.Definitions)
	}
	goa.save(packages, outputDir)
}

func (g *goa) save(packages map[string]*parser.Package, outputDir string) {
	for pkgPath, pkg := range packages {
		for filePath, file := range pkg.Node().Files {
			fileName := filepath.Base(filePath)
			outputPath := filepath.Join(outputDir, pkgPath)
			if err := os.MkdirAll(outputPath, os.ModePerm); err != nil {
				logger.Errorf("error creating output directory %s", err.Error())
			}
			writer.SaveNode(file, filepath.Join(outputPath, fileName))
		}
	}
}

func (g *goa) applyAroundAspects() {
	for _, definition := range g.definitions.List() {
		for _, function := range g.functions.List() {
			if definition.Match(function.Path()) {
				logger.Info("matched!")
				executor := &goaAST.AroundExecutor{
					Function:       function,
					Definition:     definition,
					CurrentImports: imports.GetImports(function.Parent()),
				}
				executor.Execute()
			} else {
				logger.Info("no matched!")
			}
		}
	}
}
