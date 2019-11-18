package main

import (
	"fmt"
	"github.com/wesovilabs/goa/helper"
	"github.com/wesovilabs/goa/internal"
	"github.com/wesovilabs/goa/logger"
	goaParser "github.com/wesovilabs/goa/parser"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

const defaultTargetDir = ".goa"

var excludeDirs = map[string]string{
	defaultTargetDir: defaultTargetDir,
	".git":           ".git",
	".gitignore":     ".gitignore",
}

func setUp(sourceDir, rootDir string) {
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

	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

}

func main() {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	settings, err := internal.GoaSettingFromCommandLine()
	if err != nil {
		panic(err)
	}

	if settings.Verbose {
		logger.Enable()
		defer logger.Close()
		showBanner()
	}

	setUp(settings.Path, settings.OutputDir)

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

	select {
	case <-sigCh:
		os.RemoveAll(settings.OutputDir)
		logger.Close()
		os.Exit(0)
	}
}

func showBanner() {
	fmt.Println(internal.Banner)
}
