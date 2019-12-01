package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func Test_BeyondSettingFromCommandLine(t *testing.T) {
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
				ExcludeDirs: map[string]bool{
					".git":true,
				},
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
	}

	for i, c := range cases {
		setting := &Settings{}
		fmt.Println(i)
		setting.updateWithFlags([]string{}, c.project, c.path, c.outputDir, c.pkg, c.verbose, c.work)

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

func Test_RemoveBeyondArguments(t *testing.T) {
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
	out := RemoveBeyondArguments(args)
	assert.Equal(t, "run", out[0])
	assert.Equal(t, "main.go", out[1])
}

func Test_load(t *testing.T) {
	assert := assert.New(t)
	config := load("testdata/beyond.toml")
	assert.NotNil(config)
	assert.Equal("github.com/wesovilabs/beyond-examples/settings", config.Project)
	assert.Equal("generated", config.OutputDir)
	assert.True(config.Verbose)
	assert.True(config.Work)
	assert.Len(config.Excludes, 3)
}

func Test_takePackage(t *testing.T) {
	assert := assert.New(t)
	pkg := takePackage([]string{"build", "main.go"})
	assert.Equal(".", pkg)
	pkg = takePackage([]string{"build", "cmd/main.go"})
	assert.Equal("cmd", pkg)

}


func Test_BeyondSettingFromCommandLineFlags(t *testing.T){
	BeyondSettingFromCommandLine([]string{"buil"})
}
