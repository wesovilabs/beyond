package internal

import (
	"github.com/wesovilabs/beyond/adapter"
	"github.com/wesovilabs/beyond/advice"
	"github.com/wesovilabs/beyond/helper"
	"github.com/wesovilabs/beyond/joinpoint"
	"github.com/wesovilabs/beyond/logger"
	"github.com/wesovilabs/beyond/parser"

	"os"
	"path/filepath"
)

// Run main function in charge of orchestrating code generation
func Run(rootPkg string, packages map[string]*parser.Package, outputDir string) {
	excludePaths := advice.GetExcludePaths(packages)
	advices := advice.GetAdvices(packages)
	joinPoints := joinpoint.GetJoinPoints(rootPkg, advices, excludePaths, packages)

	for _, jp := range joinPoints.List() {
		if len(jp.Advices()) > 0 {
			logger.Infof(`[joinpoint] %s.%s => %s`, jp.Pkg(), jp.Name(), jp.Path())

			for _, d := range jp.Advices() {
				logger.Infof("   - %s", d.Name())
			}

			adapter.Adapter(jp, jp.Advices())
		}
	}

	save(packages, outputDir)
}

func save(packages map[string]*parser.Package, outputDir string) {
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
