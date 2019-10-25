package internal

import (
	"github.com/wesovilabs/goa/aspect"
	"github.com/wesovilabs/goa/function"
	"github.com/wesovilabs/goa/imports"
	goaAST "github.com/wesovilabs/goa/internal/ast"
	"github.com/wesovilabs/goa/internal/writer"
	"github.com/wesovilabs/goa/logger"
	"go/ast"
	"os"
	"path/filepath"
)

type goa struct {
	functions   *function.Functions
	definitions *aspect.Definitions
}

func Run(packages map[string]*ast.Package, outputDir string) {
	goa := &goa{}
	goa.definitions = aspect.GetDefinitions(packages)
	goa.functions = function.GetFunctions(packages)
	for _, f := range goa.functions.List() {
		logger.Infof(`[function] %s.%s => %s`, f.Pkg(), f.Name(), f.Path())
	}
	for _, a := range goa.definitions.List() {
		logger.Infof(`[aspect  ] %s.%s`, a.Pkg(), a.Name())
	}
	goa.applyAroundAspects()
	goa.save(packages, outputDir)
}

func (g *goa) save(packages map[string]*ast.Package, outputDir string) {
	for pkgPath, pkg := range packages {
		logger.Infof("applying changes in %s", pkgPath)
		for filePath, file := range pkg.Files {
			logger.Infof("file %s  %s", filePath, pkgPath)
			fileName := filepath.Base(filePath)
			outputPath := filepath.Join(outputDir, pkgPath)
			logger.Infof("output path: %s", outputPath)
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
