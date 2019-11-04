package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	assert := assert.New(t)
	packages := New("testdata",
		"github.com/wesovilabs/goa/parser/testdata").
		Parse("")
	assert.NotNil(packages)
	assert.Equal(5, len(packages))
	assert.Contains(packages, "")
	assert.Contains(packages, "a")
	assert.Contains(packages, "b")
	assert.Contains(packages, "b/b1")
	assert.Contains(packages, "c")
	assert.Equal("a", packages["a"].Node().Name)
	assert.Equal("b", packages["b"].Node().Name)
	assert.Equal("b1", packages["b/b1"].Node().Name)
	assert.Equal("c", packages["c"].Node().Name)
	assert.Equal(0, len(packages["c"].Node().Imports))
	assert.Equal(1, len(packages["b"].Node().Files))
	assert.Equal(2, len(packages["a"].Node().Files))

}
