package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/beyond/internal"
	"github.com/wesovilabs/beyond/logger"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func Test_Beyond(t *testing.T) {
	assert := assert.New(t)
	packages := testPackages()
	assert.NotNil(packages)
	path, err := ioutil.TempDir(".", "prefix")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(path)
	internal.Run(pkg, packages, path)
}

func Test_BeyondApp(t *testing.T) {
	logger.Enable()
	settings := &internal.Settings{
		Work:    true,
		Verbose: true,
		Output:  filepath.Join("generated"),
		Path:    filepath.Join("../testdata"),
		Project: "github.com/wesovilabs/beyond/testdata",
		ExcludeDirs: map[string]bool{
			"generated": true,
			".git":      true,
		},
		Pkg: "cmd",
	}
	goCmd := internal.GoCommand(settings, []string{"run", "cmd/main.go"}).Do()

	exitStatus := internal.ExecuteMain(goCmd, settings)
	assert.Equal(t, 0, exitStatus)
	assert.True(t, true)
	if err := os.RemoveAll("generated"); err != nil {
		panic(err.Error())
	}

}
