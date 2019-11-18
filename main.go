package main

import (
	"fmt"
	"github.com/wesovilabs/goa/helper"
	"github.com/wesovilabs/goa/internal"
	"github.com/wesovilabs/goa/logger"
	goaParser "github.com/wesovilabs/goa/parser"
	"os"
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

func main() {
	sigCh := make(chan os.Signal, 1)

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

	if !settings.Work {
		defer func() {
			logger.Infof("wipe out directory %s", settings.OutputDir)
			if err := os.RemoveAll(settings.OutputDir); err != nil {
				logger.Error(err.Error())
			}
		}()
	} else {
		fmt.Printf("Temporary directory is %s\n", settings.OutputDir)
	}
	packages := goaParser.
		New(settings.Path, settings.Project).
		Parse(settings.Pkg)
	internal.Run(settings.Project, packages, settings.OutputDir)
	goArgs := internal.RemoveGoaArguments(os.Args[1:])

	if goCommand := internal.GoCommand(settings, goArgs); goCommand != nil {
		cmd := goCommand.Do()
		if cmd.Wait() != nil {
			<-sigCh
			logger.Infof("Removing directory %s", settings.OutputDir)

			if !settings.Work {
				os.RemoveAll(settings.OutputDir)
			}

			logger.Close()
			os.Exit(0)
		}

		logger.Info("execution completed successfully!")
	}
}

func showBanner() {
	fmt.Println(internal.Banner)
}
