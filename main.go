package main

import (
	"flag"
	"fmt"
	"github.com/wesovilabs/goa/goa"
	"github.com/wesovilabs/goa/logger"
	"go/parser"
	"go/token"
	"os"
)

type settings struct {
	input      string
	outputDir  string
	showBanner bool
	verbose    bool
}

func parseInput() *settings {
	var outputDir, inputFile string
	var showBanner, verbose bool
	flag.StringVar(&outputDir, "output", ".goa", "output directory")
	flag.StringVar(&inputFile, "input", "main.go", "main file")
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
	fileSet := token.NewFileSet()
	logger.Log().Infof("parsing file %s", settings.input)
	file, err := parser.ParseFile(fileSet, settings.input, nil, parser.ParseComments)
	if err != nil {
		logger.Log().Fatalf("error while parsing file: '%v'", err)
	}
	if err := goa.Init().Execute(file); err != nil {
		logger.Log().Fatalf("error while generating code: '%v'", err)
	}
	logger.Log().Info("code was generated successfully!")

}

func showBanner() {
	fmt.Println(goa.Banner)
}
