package main

import (
	"flag"
	"fmt"
	"github.com/wesovilabs/goa/goa"
	"github.com/wesovilabs/goa/logger"
	goaParser "github.com/wesovilabs/goa/parser"
	"go/ast"
	"os"
	"path/filepath"
)

type settings struct {
	input      string
	mainPkg    string
	project    string
	outputDir  string
	showBanner bool
	verbose    bool
}

func parseInput() *settings {
	var outputDir, inputFile string
	var showBanner, verbose bool
	pwd, _ := os.Getwd()
	flag.StringVar(&outputDir, "output", fmt.Sprintf("%s%s%s", pwd, filepath.Separator, ".goa"), "output directory")
	flag.StringVar(&inputFile, "input", fmt.Sprintf("%s%s%s", "main.go"), "main file")
	flag.BoolVar(&showBanner, "banner", true, "display goa banner")
	flag.BoolVar(&verbose, "verbose", true, "print info level logs to stdout")
	flag.Parse()
	return &settings{
		showBanner: showBanner,
		input:      inputFile,
		outputDir:  outputDir,
		verbose:    verbose,
	}
}

func main() {
	settings := parseInput()
	if settings.showBanner {
		showBanner()
	}
	if settings.verbose {
		logger.Enable()
		defer logger.Close()
	}
	if err := os.Mkdir(settings.outputDir, os.ModePerm); err != nil {
		//panic("error while creating output directory")
	}
	rootDir := "/Users/ivan/Workspace/BBVA/ECSKERNEL/main-controller"
	rootPath := "cmd/main-controller"
	project := "main-controller" // // This values must be taken from go.mod in `path`
	packages := findPackages(rootDir, rootPath, project)
	goaApp := goa.Init()
	for pkgName, pkg := range packages {
		logger.Infof("applying changes in %s", pkgName)
		for filePath, file := range pkg.Files {
			logger.Infof("file %s  %s", filePath, pkgName)
			fileName := filepath.Base(filePath)
			if err := goaApp.Execute(pkgName, fileName, file); err != nil {
				logger.Fatal("error while generating code: '%v'", err)
			}
		}
	}
	logger.Info("code was generated successfully!")
	fmt.Println("____________________")

}

func showBanner() {
	fmt.Println(goa.Banner)
}

func findPackages(rootDir, rootPath, project string) map[string]*ast.Package {
	return goaParser.
		New(rootDir, project, false).
		Parse(rootPath)
}
