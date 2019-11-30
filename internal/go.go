package internal

import (
	"os"
	"os/exec"
	"path/filepath"
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
	hasOutput := false

	for i := range args {
		arg := args[i]

		if arg == "-o" && len(args) > i+1 {
			args[i+1] = transformPath(args[i+1], settings.Path)
			hasOutput = true

			break
		}
	}

	if !hasOutput {
		newArgs := append(args[:1], "-o", filepath.Join(settings.Path, "app"))
		args = append(newArgs, args[1:]...)
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

// Do  execute go command
func (e *Executor) Do() *exec.Cmd {
	args := append([]string{e.cmd}, e.args...)
	//nolint
	cmd := exec.Command("go")
	cmd.Args = args
	cmd.Env = os.Environ()
	cmd.Dir = e.settings.OutputDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
