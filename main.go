package main

import (
	"fmt"
	"github.com/wesovilabs/goa/helper"
	"github.com/wesovilabs/goa/internal"
	"github.com/wesovilabs/goa/logger"
	goaParser "github.com/wesovilabs/goa/parser"

	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func setUp(sourceDir, rootDir string, excludeDirs map[string]bool) {
	logger.Infof("copying resources to directory %s", rootDir)

	if _, err := os.Stat(rootDir); err != nil {
		if err := os.MkdirAll(rootDir, 0755); err != nil {
			logger.Error(err.Error())
		}
	}

	if err := helper.CopyDirectory(sourceDir, rootDir, excludeDirs); err != nil {
		panic(err.Error())
	}

	logger.Infof("directory %s contains a copy of your path", rootDir)
}

func main() {
	start := time.Now()
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	settings, err := internal.GoaSettingFromCommandLine(os.Args[1:])

	if err != nil {
		panic(err)
	}

	if settings.Verbose {
		logger.Enable()
		defer logger.Close()
		showBanner()
	}

	setUp(settings.Path, settings.OutputDir, settings.ExcludeDirs)

	if !settings.Work {
		defer func() {
			logger.Infof("wipe out directory %s", settings.OutputDir)
			if err := os.RemoveAll(settings.OutputDir); err != nil {
				logger.Error(err.Error())
			}
		}()
	} else {
		fmt.Printf("[ WORKDIR ] %s\n", settings.OutputDir)
	}

	packages := goaParser.
		New(settings.Path, settings.Project).
		Parse(settings.Pkg)

	internal.Run(settings.Project, packages, settings.OutputDir)
	goArgs := internal.RemoveGoaArguments(os.Args[1:])
	end := time.Now()
	logger.Infof("[goa] goa transformation took %v milliseconds", end.Sub(start).Milliseconds())
	logger.Infof("[workdir] %s", settings.OutputDir)
	logger.Infof("[command] go %s", strings.Join(goArgs, " "))

	if settings.Verbose {
		println("---")
		println()
	}

	command(settings, goArgs, sigCh)
}

func command(settings *internal.Settings, goArgs []string, sigCh chan os.Signal) {
	goCommand := internal.GoCommand(settings, goArgs).Do()

	var execStatus syscall.WaitStatus

	exitStatus := 0

	if err := goCommand.Start(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	go func() {
		if err := goCommand.Wait(); err != nil {
			fmt.Println(err.Error())
		}
		execStatus = goCommand.ProcessState.Sys().(syscall.WaitStatus)
		exitStatus = execStatus.ExitStatus()
		if exitStatus >= 0 {
			close(sigCh)
		}
	}()
	<-sigCh

	if !settings.Work {
		logger.Infof("Removing directory %s", settings.OutputDir)
		os.RemoveAll(settings.OutputDir)
	}

	logger.Close()
	os.Exit(exitStatus)
}

func showBanner() {
	fmt.Println(internal.Banner)
}
