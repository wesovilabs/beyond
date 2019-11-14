package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/advice"
	"github.com/wesovilabs/goa/joinpoint"
	"github.com/wesovilabs/goa/parser"
	"testing"
)

const project = "github.com/wesovilabs/goa/matcher/testdata"

func TestFindMatches(t *testing.T) {
	assert := assert.New(t)
	packages := parser.
		New("testdata", project).
		Parse("")
	definitions := advice.GetAdvices("testdata", packages)
	functions := joinpoint.GetJoinPoints(project, packages)
	matches := FindMatches(functions, definitions)
	assert.Len(matches, 2)
	assert.Len(matches[0].Advices, 2)
	assert.Len(matches[1].Advices, 2)

}
