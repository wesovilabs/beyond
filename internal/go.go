package internal

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var goCmds = map[string]func(*Settings, []string) *Executor{
	"build":    newGoBuild,
	"run":      newGoRun,
	"generate": newGoGenerate,
}

// GoCommand reutrns the go command to be executed
func GoCommand(settings *Settings, args []string) *Executor {
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

func newGoBuild(settings *Settings, args []string) *Executor {
	var hasOutputFlag bool

	for i := range args {
		arg := args[i]

		if arg == "-o" && len(args) > i+1 {
			args[i+1] = transformPath(args[i+1], settings.Path)
			break
		}
	}

	if !hasOutputFlag {
		args = append(args, "-o", filepath.Join(settings.Path, "app"))
	}

	return &Executor{"build", args, settings}
}

func newGoRun(settings *Settings, args []string) *Executor {
	return &Executor{"run", args, settings}
}
func newGoGenerate(settings *Settings, args []string) *Executor {
	return &Executor{"generate", args, settings}
}

// Executor struct for wrapping go commands
type Executor struct {
	cmd      string
	args     []string
	settings *Settings
}

func (e *Executor) Do() *exec.Cmd {
	args := append([]string{e.cmd}, e.args[1:]...)
	//nolint
	cmd := exec.Command("go", args...)
	cmd.Env = os.Environ()
	cmd.Dir = e.settings.OutputDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
