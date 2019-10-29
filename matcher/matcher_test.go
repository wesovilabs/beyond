package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/aspect"
	"github.com/wesovilabs/goa/function"
	"github.com/wesovilabs/goa/parser"
	"testing"
)

func TestFindMatches(t *testing.T) {
	assert := assert.New(t)
	packages := parser.
		New("testdata", "testdata", false).
		Parse("testdata", "")
	definitions := aspect.GetDefinitions("testdata", packages)
	functions := function.GetFunctions("testdata", packages)
	matches := FindMatches(functions, definitions)
	assert.Len(matches, 2)
	assert.Len(matches[0].Definitions, 2)
	assert.Len(matches[1].Definitions, 2)

}
