package function

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/parser"
	"testing"
)

func TestGetFunctions(t *testing.T) {
	assert := assert.New(t)
	packages := parser.
		New("testdata", "testdata", false).
		Parse("testdata", "")
	functions := GetFunctions(packages)
	assert.NotNil(functions)
	assert.NotNil(functions.List())
	assert.Equal(23, len(functions.List()))
}
