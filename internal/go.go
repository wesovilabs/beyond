package internal

import (
	"github.com/wesovilabs/goa/logger"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var goCmds = map[string]func(*Settings, []string) *executor{
	"build":    newGoBuild,
	"run":      newGoRun,
	"generate": newGoGenerate,
}

func GoCommand(settings *Settings, args []string) *executor {
	for i := range args {
		arg := args[i]
		if fn, ok := goCmds[arg]; ok {
			return fn(settings, args)
		}
	}
	return nil
}

func transformPath(old string, baseDir string) string {
	if filepath.IsAbs(old) {
		return old
	}
	return filepath.Join(baseDir, old)
}

func newGoBuild(settings *Settings, args []string) *executor {
	var hasOutputFlag bool
	for i := range args {
		arg := args[i]
		if arg == "-o" {
			hasOutputFlag = true
			args[i+1] = transformPath(args[i+1], settings.RootDir)
		}
	}
	if !hasOutputFlag {
		args = append(args, "-o", settings.RootDir)
	}
	return &executor{"build", args, settings}
}

func newGoRun(settings *Settings, args []string) *executor {

	return &executor{"run", args, settings}
}
func newGoGenerate(settings *Settings, args []string) *executor {

	return &executor{"generate", args, settings}
}

type executor struct {
	cmd      string
	args     []string
	settings *Settings
}

func (e *executor) Do() *exec.Cmd {
	logger.Infof("Running go %v", e.args)
	cmd := exec.Command("go", e.args...)
	cmd.Env = os.Environ()
	cmd.Dir = e.settings.OutputDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return cmd
}
