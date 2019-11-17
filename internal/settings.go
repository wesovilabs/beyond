package internal

import (
	"flag"
	"github.com/wesovilabs/goa/helper"
	"os"
	"path/filepath"
)

const defaultTargetDir = ".goa"

// Settings Goa settings
type Settings struct {
	Path      string
	Project   string
	OutputDir string
	Verbose   bool
}

// GoaSettingFromCommandLine returns the GoaSettings from the command line args
func GoaSettingFromCommandLine() (*Settings, error) {
	var path, project, outputDir string

	var verbose bool

	pwd, _ := os.Getwd()

	flag.StringVar(&project, "project", "", "project name")
	flag.StringVar(&path, "path", pwd, "path")
	flag.StringVar(&outputDir, "output", filepath.Join(path, defaultTargetDir), "output directory")
	flag.BoolVar(&verbose, "verbose", false, "print info level logs to stdout")
	flag.Parse()

	if project == "" {
		module, err := helper.GetModuleName(path)
		if err != nil {
			return nil, err
		}

		project = module
	}

	return &Settings{
		Path:      path,
		Project:   project,
		OutputDir: outputDir,
		Verbose:   verbose,
	}, nil
}

// RemoveGoaArguments removes goa arguments from the list of arguments
func RemoveGoaArguments(input []string) []string {
	arguments := input

	for i, arg := range input {
		switch arg {
		case "--project", "--output", "--path":
			arguments = append(arguments[0:i], arguments[i+2:]...)
		case "--verbose":
			arguments = append(arguments[0:i], arguments[i+1:]...)
		}
	}

	return arguments
}
