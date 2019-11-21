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
	Path        string
	Project     string
	OutputDir   string
	Pkg         string
	ExcludeDirs map[string]bool
	Verbose     bool
	Work        bool
}

// GoaSettingFromCommandLine returns the GoaSettings from the command line args
func GoaSettingFromCommandLine() (*Settings, error) {
	var path, project, outputDir, pkg string

	var verbose, work bool

	pwd, _ := os.Getwd()

	flag.StringVar(&project, "project", "", "project name")
	flag.StringVar(&path, "path", pwd, "path")
	flag.StringVar(&outputDir, "output", "", "output directory")
	flag.StringVar(&pkg, "package", "", "relative path to the main package")
	flag.BoolVar(&verbose, "verbose", false, "print info level logs to stdout")
	flag.BoolVar(&work, "work", false, "print the name of the temporary work directory and do not delete it when exiting")
	flag.Parse()
	return createSettings(project, path, outputDir, pkg, verbose, work)

}

func createSettings(project, path, outputDir, pkg string, verbose, work bool) (*Settings, error) {
	if project == "" {
		module, err := helper.GetModuleName(path)
		if err != nil {
			return nil, err
		}

		project = module
	}

	var outErr error

	if outputDir == "" {
		if targetDir, err := ioutil.TempDir("", "goa"); err == nil {
			outputDir = targetDir
		} else {
			outputDir = filepath.Join(path, defaultTargetDir)
		}
	} else {
		if outputDir, outErr = filepath.Abs(outputDir); outErr != nil {
			outputDir = filepath.Join(path, defaultTargetDir)
		}
	}

	excludeDirs := map[string]bool{}
	addDefaultExcludes(".git", excludeDirs)
	addDefaultExcludes(outputDir, excludeDirs)

	return &Settings{
		Path:        path,
		Project:     project,
		OutputDir:   outputDir,
		Verbose:     verbose,
		ExcludeDirs: excludeDirs,
		Pkg:         pkg,
		Work:        work,
	}, nil
}

// RemoveGoaArguments removes goa arguments from the list of arguments
func RemoveGoaArguments(input []string) []string {
	out := make([]string, 0)
	argsIndex := make(map[int]bool)

	for i, arg := range input {
		switch arg {
		case "--project", "--output", "--path", "--package":
			argsIndex[i] = true
			argsIndex[i+1] = true
		case "--verbose", "--work":
			argsIndex[i] = true
		}
	}

	for i := range input {
		if !argsIndex[i] {
			out = append(out, input[i])
		}
	}

	return out
}

func addDefaultExcludes(localPath string, excludes map[string]bool) {
	if absPath, err := filepath.Abs(localPath); err == nil {
		excludes[absPath] = true
	}
}
