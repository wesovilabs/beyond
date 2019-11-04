package aspect

import (
	"fmt"
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
				"a/b.b()(int,*string)",
			},
		},
		{
			regExp: internal.NormalizeExpression("*.*(...)..."),
			matches: []string{
				"a.b(string)int",
				"a.b(map[string]interface{})(int,*string)",
				"a/b.b(map[string]interface{})(int,*string)",
				"a/b.b()(int,*string)",
			},
			noMatches: []string{
				"a/b.c.d(string)",
				"a.c.d(string)",
			},
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		fmt.Println(c.regExp.String())
		def := &Definition{
			regExp: c.regExp,
		}
		for _, m := range c.matches {
			if !assert.True(def.Match(m)) {
				t.FailNow()
			}
		}
		for _, m := range c.noMatches {
			if !assert.False(def.Match(m)) {
				t.FailNow()
			}
		}
	}
}
