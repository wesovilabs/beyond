package internal

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func Test_GoaSettingFromCommandLine(t *testing.T) {
	assert := assert.New(t)
	pwd, _ := os.Getwd()
	cases := []struct {
		project, path, outputDir, pkg string
		verbose, work                 bool
		setting                       *Settings
	}{
		{
			project:   "myproject",
			path:      "",
			outputDir: "out",
			pkg:       "cmd/app",
			verbose:   false,
			work:      true,
			setting: &Settings{
				Verbose:   false,
				Work:      true,
				Pkg:       "cmd/app",
				OutputDir: filepath.Join(pwd, "out"),
				Path:      "",
				Project:   "myproject",
			},
		},
		{
			project:   "myproject",
			path:      "",
			outputDir: "",
			pkg:       "",
			setting: &Settings{
				Verbose:   false,
				Work:      false,
				Pkg:       "",
				OutputDir: filepath.Join(pwd, "out"),
				Path:      "",
				Project:   "myproject",
			},
		},
		{
			project:   "",
			path:      "",
			outputDir: "",
			pkg:       "",
		},
	}

	for _, c := range cases {
		setting, err := createSettings(c.project, c.path, c.outputDir, c.pkg, c.verbose, c.work)
		if c.setting == nil {
			assert.Nil(setting)
			assert.NotNil(err)
			continue
		}
		assert.Equal(c.setting.Project, setting.Project)
		assert.Equal(c.setting.Path, setting.Path)
		assert.Equal(c.setting.Pkg, setting.Pkg)
		if c.outputDir != "" {
			assert.Equal(c.setting.OutputDir, setting.OutputDir)
		} else {
			assert.NotEmpty(setting.OutputDir)
		}
		assert.Equal(c.setting.Work, setting.Work)
		assert.Equal(c.setting.Verbose, setting.Verbose)
		assert.Len(setting.ExcludeDirs, 2)

	}
}

func Test_RemoveGoaArguments(t *testing.T) {
	args := []string{
		"--work",
		"--verbose",
		"--project",
		"myproject",
		"--path",
		"mypath",
		"--package",
		"mypackage",
		"run",
		"main.go",
	}
	out := RemoveGoaArguments(args)
	assert.Equal(t, "run", out[0])
	assert.Equal(t, "main.go", out[1])
}
