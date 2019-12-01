package main

import (
	"flag"
	"fmt"
	"github.com/wesovilabs/beyond/internal"
	"github.com/wesovilabs/beyond/logger"
	"os"
	"os/exec"
)

func usage() {
	fmt.Println("usage: [env_vars] beyond [beyond_flags] go_command [go_flags]")
	fmt.Println("[beyond_flags]")
	flag.PrintDefaults()
	fmt.Println("\n[go_command]")
	fmt.Println("  build: Build compiles the packages named by the import paths")
	fmt.Println("  run: Run compiles and runs the named main Go package.")
	fmt.Println("  generate: Generate runs commands described by directives within existing files.")
}

func showBanner() {
	fmt.Println(internal.Banner)
}

func goCommand(settings *internal.Settings, goArgs []string) *exec.Cmd {
	executor := internal.GoCommand(settings, goArgs)
	if executor == nil {
		return nil
	}

	return executor.Do()
}

func main() {

	settings := internal.BeyondSettingFromCommandLine(os.Args[1:])
	goArgs := internal.RemoveBeyondArguments(os.Args[1:])

	goCmd := goCommand(settings, goArgs)
	if goCmd == nil {
		showBanner()
		usage()
		return
	}

	if settings.Verbose {
		logger.Enable()
		defer logger.Close()
		showBanner()
	}

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
	internal.ExecuteMain(goCmd, settings)
}
