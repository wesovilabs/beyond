package tests

import (
	"github.com/wesovilabs/beyond/parser"
	"os"
	"os/exec"
)

const (
	pkg          = "github.com/wesovilabs/beyond/testdata"
	examplesRepo = "http://github.com/wesovilabs/beyond-examples.git"
	goPath       = "../testdata"
)

func testPackages() map[string]*parser.Package {
	return parser.New(goPath, pkg).Parse("cmd")
}

func cloneBeyondExamplesRepo() {
	os.Mkdir("beyond-examples",os.ModeDir)

	cmd := exec.Command("git")
	cmd.Args = []string{"clone", examplesRepo,"beyond-examples"}
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
