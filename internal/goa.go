package internal

import (
	"github.com/wesovilabs/goa/adapter"
	"github.com/wesovilabs/goa/advice"
	"github.com/wesovilabs/goa/helper"
	"github.com/wesovilabs/goa/joinpoint"
	"github.com/wesovilabs/goa/logger"
	"github.com/wesovilabs/goa/match"
	"github.com/wesovilabs/goa/parser"

	"os"
	"path/filepath"
)

type goa struct {
	joinPoints *joinpoint.JoinPoints
	advices    *advice.Advices
}

func (g *goa) removeNonInterceptedJoinPoints() {
	output := &joinpoint.JoinPoints{}

	for _, jp := range g.joinPoints.List() {
		valid := true

		if jp.Name() == "main" || jp.Name() == "Goa" {
			continue
		}

		for index := range g.advices.List() {
			a := g.advices.List()[index]
			if a.Name() == jp.Name() && a.Pkg() == jp.PkgPath() {
				valid = false
				continue
			}
		}

		if valid {
			output.AddJoinPoint(jp)
		}
	}

	g.joinPoints = output
}

// Run main function in charge of orchestrating code generation
func Run(rootPkg string, packages map[string]*parser.Package, outputDir string) {
	goa := &goa{}
	goa.advices = advice.GetAdvices(rootPkg, packages)
	goa.joinPoints = joinpoint.GetJoinPoints(rootPkg, packages)

	goa.removeNonInterceptedJoinPoints()

	for _, f := range goa.joinPoints.List() {
		logger.Infof(`[function] %s.%s => %s`, f.Pkg(), f.Name(), f.Path())
	}

	for _, a := range goa.advices.List() {
		logger.Infof(`[ advice ] %s.%s`, a.Pkg(), a.Name())
	}

	matches := match.GetMatches(goa.joinPoints, goa.advices)

	for _, match := range matches {
		logger.Infof("[ match  ] %s", match.JoinPoint.Name())

		for _, d := range match.Advices {
			logger.Infof("   - %s", d.Name())
		}

		adapter.Adapter(match.JoinPoint, match.Advices)
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

			if err := helper.Save(file, filepath.Join(outputPath, fileName)); err != nil {
				logger.Error(err.Error())
			}
		}
	}
}
