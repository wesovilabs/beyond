package internal

import (
	"fmt"
	"log"
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
	var hasOutputFlag bool

	for i := range args {
		arg := args[i]
		if arg == "-o" {
			hasOutputFlag = true
			args[i+1] = transformPath(args[i+1], settings.Pkg)
		}
	}

	if !hasOutputFlag {
		args = append(args, "-o", settings.Pkg)
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

//nolint
func (e *Executor) Do() *exec.Cmd {
	fmt.Printf("go %v\n", e.args)
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
