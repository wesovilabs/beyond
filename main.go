package main

import (
	"flag"
	"fmt"
	"github.com/wesovilabs/goa/internal"
	"github.com/wesovilabs/goa/logger"
	goaParser "github.com/wesovilabs/goa/parser"
	"go/ast"
	"os"
	"path/filepath"
)

type settings struct {
	goPath     string
	path       string
	project    string
	outputDir  string
	showBanner bool
	verbose    bool
	vendor     bool
}

func parseInput() *settings {
	var outputDir, goPath, project, path string
	var showBanner, verbose, vendor bool
	pwd, _ := os.Getwd()
	flag.StringVar(&project, "project", "", "project name")
	flag.StringVar(&path, "path", "", "path")
	flag.StringVar(&outputDir, "output", filepath.Join(pwd, ".goa"), "output directory")
	flag.StringVar(&goPath, "goPath", pwd, "go path")
	flag.BoolVar(&showBanner, "banner", true, "display goa banner")
	flag.BoolVar(&verbose, "verbose", true, "print info level logs to stdout")
	flag.BoolVar(&vendor, "vendor", false, "add vendor files to be transoformed")
	flag.Parse()
	fmt.Println(goPath)
	return &settings{
		goPath:     goPath,
		project:    project,
		path:       path,
		outputDir:  outputDir,
		showBanner: showBanner,
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
	if err := os.MkdirAll(settings.outputDir, os.ModePerm); err != nil {
		panic("error while creating output directory")
	}

	// // This values must be taken from go.mod in `path`
	packages := findPackages(settings)
	internal.Run(packages, settings.outputDir)
	logger.Info("code was generated successfully!")

}

func showBanner() {
	fmt.Println(internal.Banner)
}

func findPackages(settings *settings) map[string]*ast.Package {
	return goaParser.
		New(settings.goPath, settings.project, settings.vendor).
		Parse(settings.path)
}
