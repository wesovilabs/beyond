package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GoCommand(t *testing.T) {
	assert := assert.New(t)
	setting := &Settings{
		Verbose:   false,
		Work:      true,
		Pkg:       "cmd/app",
		OutputDir: "out",
		Path:      "current",
		Project:   "myproject",
	}
	cases := []struct {
		settings *Settings
		args     []string
		executor *Executor
	}{
		{
			args: []string{},
		},
		{
			args:     []string{"run", "main.go"},
			settings: setting,
			executor: &Executor{
				cmd:      "run",
				args:     []string{"run", "main.go"},
				settings: setting,
			},
		},
		{
			args:     []string{"build", "main.go"},
			settings: setting,
			executor: &Executor{
				cmd:      "build",
				args:     []string{"build", "main.go", "-o", "current/app"},
				settings: setting,
			},
		},
		{
			args:     []string{"generate", "main.go"},
			settings: setting,
			executor: &Executor{
				cmd:      "generate",
				args:     []string{"generate", "main.go"},
				settings: setting,
			},
		},
		{
			args:     []string{"test", "./..."},
			settings: setting,
			executor: nil,
		},
		{
			args:     []string{"build", "main.go", "-o", "build/app"},
			settings: setting,
			executor: &Executor{
				cmd:      "build",
				args:     []string{"build", "main.go", "-o", "current/build/app"},
				settings: setting,
			},
		},
	}
	for _, c := range cases {
		executor := GoCommand(c.settings, c.args)
		if c.executor == nil {
			assert.Nil(executor)
			continue
		} else {
			assert.NotNil(executor)
			assert.Equal(c.executor, executor)
		}
	}

}
