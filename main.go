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

func main() {
	start := time.Now()
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	settings := internal.GoaSettingFromCommandLine(os.Args[1:])
	goArgs := internal.RemoveGoaArguments(os.Args[1:])

	goCmd := goCommand(settings, goArgs)
	if goCmd == nil {
		showBanner()

		fmt.Println("usage: [env_vars] goa [goa_flags] go_command [go_flags]\n\n")
		fmt.Println("[goa_flags]")
		flag.PrintDefaults()
		fmt.Println("\n[go_command]")
		fmt.Println("  build: Build compiles the packages named by the import paths")
		fmt.Println("  run: Run compiles and runs the named main Go package.")
		fmt.Println("  generate: Generate runs commands described by directives within existing files.")
		return
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

	end := time.Now()
	logger.Infof("[goa] goa transformation took %v milliseconds", end.Sub(start).Milliseconds())
	logger.Infof("[workdir] %s", settings.OutputDir)
	logger.Infof("[command] %s", goCmd.String())

	if settings.Verbose {
		println("---")
		println()
	}

	runGoCommand(goCmd, settings, sigCh)
}

func goCommand(settings *internal.Settings, goArgs []string) *exec.Cmd {
	executor := internal.GoCommand(settings, goArgs)
	if executor == nil {
		return nil
	}

	return executor.Do()
}

func runGoCommand(goCommand *exec.Cmd, settings *internal.Settings, sigCh chan os.Signal) {
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
