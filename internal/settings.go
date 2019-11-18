package internal

import (
	"flag"
	"github.com/wesovilabs/goa/helper"
	"io/ioutil"
	"os"
	"path/filepath"
)

const defaultTargetDir = ".goa"

// Settings Goa settings
type Settings struct {
	Path      string
	Project   string
	OutputDir string
	RootDir   string
	Verbose   bool
	Work      bool
}

// GoaSettingFromCommandLine returns the GoaSettings from the command line args
func GoaSettingFromCommandLine() (*Settings, error) {
	var path, project, outputDir string

	var verbose, work bool

	pwd, _ := os.Getwd()

	flag.StringVar(&project, "project", "", "project name")
	flag.StringVar(&path, "path", pwd, "path")
	flag.StringVar(&outputDir, "output", "", "output directory")
	flag.BoolVar(&verbose, "verbose", false, "print info level logs to stdout")
	flag.BoolVar(&work, "work", false, "print the name of the temporary work directory and do not delete it when exiting")
	flag.Parse()

	if project == "" {
		module, err := helper.GetModuleName(path)
		if err != nil {
			return nil, err
		}

		project = module
	}

	if outputDir == "" {
		if targetDir, err := ioutil.TempDir("", "goa"); err == nil {
			outputDir = targetDir
		} else {
			outputDir = filepath.Join(path, defaultTargetDir)
		}
	}

	return &Settings{
		Path:      path,
		Project:   project,
		OutputDir: outputDir,
		Verbose:   verbose,
		RootDir:   pwd,
		Work:      work,
	}, nil
}

// RemoveGoaArguments removes goa arguments from the list of arguments
func RemoveGoaArguments(input []string) []string {
	arguments := input

	for i, arg := range input {
		switch arg {
		case "--project", "--output", "--path", "--package":
			arguments = append(arguments[0:i], arguments[i+2:]...)
		case "--verbose", "--work":
			arguments = append(arguments[0:i], arguments[i+1:]...)
		}
	}

	return arguments
}
