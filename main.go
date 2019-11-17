package main

import (
	"flag"
	"fmt"
	"github.com/wesovilabs/goa/helper"
	"github.com/wesovilabs/goa/internal"
	"github.com/wesovilabs/goa/logger"
	goaParser "github.com/wesovilabs/goa/parser"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const defaultTargetDir = ".goa"

var excludeDirs = map[string]string{
	defaultTargetDir: defaultTargetDir,
	".git":           ".git",
	".gitignore":     ".gitignore",
}

type settings struct {
	goPath     string
	path       string
	project    string
	outputDir  string
	showBanner bool
	verbose    bool
}

func loadSettings() *settings {
	var outputDir, goPath, project, path string

	var showBanner, verbose bool

	pwd, err := os.Getwd()

	if err != nil {
		logger.Error(err.Error())
	}

	flag.StringVar(&project, "project", "", "project name")
	flag.StringVar(&path, "path", pwd, "path")
	flag.StringVar(&goPath, "goPath", "", "go path")
	flag.StringVar(&outputDir, "output", filepath.Join(goPath, defaultTargetDir), "output directory")
	flag.BoolVar(&showBanner, "banner", false, "display goa banner")
	flag.BoolVar(&verbose, "goaVerbose", false, "print info level logs to stdout")
	flag.Parse()

	goPath = filepath.Join(pwd, goPath)
	outputDir = filepath.Join(goPath, outputDir)
	path = strings.TrimPrefix(path, goPath)

	return &settings{
		goPath:     goPath,
		project:    project,
		path:       path,
		outputDir:  outputDir,
		showBanner: showBanner,
		verbose:    verbose,
	}
}

func setpUpRootDir(sourceDir, rootDir string) {
	logger.Infof("copying resources to directory %s", rootDir)
	if err := helper.CopyDirectory(sourceDir, rootDir, excludeDirs); err != nil {
		panic(err.Error())
	}
	logger.Infof("directory %s contains a copy of your path", rootDir)
}

func run(rootDir string, arguments []string) {
	cmd := exec.Command("go", arguments...)
	cmd.Env = os.Environ()
	cmd.Dir = rootDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func main() {
	settings, err := internal.GoaSettingFromCommandLine()
	if err != nil {
		panic(err)
	}
	if settings.Verbose {
		logger.Enable()
		defer logger.Close()
		showBanner()
	}
	setpUpRootDir(settings.Path, settings.OutputDir)

	defer func() {
		logger.Infof("wipe out directory %s", settings.OutputDir)
		if err := os.RemoveAll(settings.OutputDir); err != nil {
			logger.Error(err.Error())
		}

	}()
	packages := goaParser.
		New(settings.Path, settings.Project).
		Parse("")
	internal.Run(settings.Project, packages, settings.OutputDir)
	goArgs := internal.RemoveGoaArguments(os.Args[1:])
	run(settings.OutputDir, goArgs)
}

func showBanner() {
	fmt.Println(internal.Banner)
}
