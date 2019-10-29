package parser

import (
	"github.com/stretchr/testify/assert"
	"go/parser"
	"go/token"
	"testing"
)

func Test_excludeTestPackages(t *testing.T) {
	assert := assert.New(t)
	assert.False(excludeTestPackages("name_test"))
	assert.True(excludeTestPackages("name"))
}

func Test_applyPkgFilters(t *testing.T) {
	assert := assert.New(t)
	packages, _ := parser.ParseDir(&token.FileSet{}, "testdata", nil, parser.ParseComments)
	pkg := applyPkgFilters(packages, excludeTestPackages)
	assert.NotNil(pkg)
	assert.Equal(pkg.Name, "main")
}
