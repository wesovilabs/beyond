package tests

import "github.com/wesovilabs/goa/parser"

const (
	pkg    = "github.com/wesovilabs/goa/testdata"
	goPath = "../testdata"
)

func testPackages() map[string]*parser.Package {
	return parser.New(goPath, pkg).Parse("cmd")
}
