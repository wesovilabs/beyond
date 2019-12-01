package internal

import (
	"fmt"
	"github.com/wesovilabs/beyond/helper"
	"github.com/wesovilabs/beyond/logger"
	"github.com/wesovilabs/beyond/parser"
	"log"
	"os"
	"os/exec"
	"os/signal"
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

func ExecuteMain(goCmd *exec.Cmd, settings *Settings) int{
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	start := time.Now()
	setUp(settings.Path, settings.OutputDir, settings.ExcludeDirs)
	packages := parser.
		New(settings.Path, settings.Project).
		Parse(settings.Pkg)

	Run(settings.Project, packages, settings.OutputDir)

	end := time.Now()
	logger.Infof("[beyond] beyond transformation took %v milliseconds", end.Sub(start).Milliseconds())
	logger.Infof("[workdir] %s", settings.OutputDir)
	logger.Infof("[command] %s", goCmd.String())

	if settings.Verbose {
		println("---")
		println()
	}

	return runGoCommand(goCmd, settings, sigCh)
}

func runGoCommand(goCommand *exec.Cmd, settings *Settings, sigCh chan os.Signal) int{
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
	return exitStatus
}
