package aspect

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/aspect/internal"
	"regexp"
	"testing"
)

func TestMatch(t *testing.T) {
	cases := []struct {
		regExp    *regexp.Regexp
		matches   []string
		noMatches []string
	}{
		{
			regExp: internal.NormalizeExpression("*.*(*)..."),
			matches: []string{
				"a.b(string)int",
				"a.b(map[string]interface{})(int,*string)",
				"a/b.b(map[string]interface{})(int,*string)",
			},
			noMatches: []string{
				"a/b.c.d(string)",
				"a.c.d(string)",
			},
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		def := &Definition{
			regExp: c.regExp,
		}
		for _, m := range c.matches {
			assert.True(def.Match(m))
		}
		for _, m := range c.noMatches {
			assert.False(def.Match(m))
		}
	}
}
