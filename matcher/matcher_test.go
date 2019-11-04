package matcher

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/aspect"
	"github.com/wesovilabs/goa/function"
	"github.com/wesovilabs/goa/parser"
	"testing"
)

const project = "github.com/wesovilabs/goa/matcher/testdata"

func TestFindMatches(t *testing.T) {
	assert := assert.New(t)
	packages := parser.
		New("testdata", project).
		Parse("")
	definitions := aspect.GetDefinitions("testdata", packages)
	functions := function.GetFunctions(packages)
	for _, f := range functions.List() {
		fmt.Println(f.Path())
	}
	matches := FindMatches(functions, definitions)
	assert.Len(matches, 2)
	assert.Len(matches[0].Definitions, 2)
	assert.Len(matches[1].Definitions, 2)

}
