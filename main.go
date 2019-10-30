package main

import (
	"flag"
	"fmt"
	"github.com/wesovilabs/goa/internal"
	"github.com/wesovilabs/goa/logger"
	goaParser "github.com/wesovilabs/goa/parser"
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
	pwd, err := os.Getwd()
	if err != nil {
		logger.Error(err.Error())
	}
	flag.StringVar(&project, "project", "", "project name")
	flag.StringVar(&path, "path", "", "path")
	flag.StringVar(&goPath, "goPath", pwd, "go path")
	flag.StringVar(&outputDir, "output", filepath.Join(goPath, ".goa"), "output directory")
	flag.BoolVar(&showBanner, "banner", false, "display goa banner")
	flag.BoolVar(&verbose, "verbose", false, "print info level logs to stdout")
	flag.BoolVar(&vendor, "vendor", false, "add vendor files to be transoformed")
	flag.Parse()
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
	fmt.Printf("%#v", settings)
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
	internal.Run(settings.project, packages, settings.outputDir)
	logger.Info("code was generated successfully!")

}

func showBanner() {
	fmt.Println(internal.Banner)
}

func findPackages(settings *settings) map[string]*goaParser.Package {
	return goaParser.
		New(settings.goPath, settings.project, settings.vendor).
		Parse(settings.project, settings.path)
}
