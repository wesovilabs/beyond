package tests

import (
	"github.com/wesovilabs/beyond/parser"
)

const (
	pkg          = "github.com/wesovilabs/beyond/testdata"
	examplesRepo = "http://github.com/wesovilabs/beyond-examples.git"
	goPath       = "../testdata"
)

func testPackages() map[string]*parser.Package {
	return parser.New(goPath, pkg).Parse("cmd")
}
