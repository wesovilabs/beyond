package joinpoint

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/parser"
	"testing"
)

func TestGetFunctions(t *testing.T) {
	project := "github.com/wesovilabs/goa/joinpoint/testdata"
	assert := assert.New(t)
	packages := parser.
		New("testdata", project).
		Parse("")
	functions := GetFunctions(packages)
	assert.NotNil(functions)
	assert.NotNil(functions.List())
	assert.Equal(25, len(functions.List()))
}

func TestGetFunctionsComplex(t *testing.T) {
	project := "github.com/wesovilabs/goa/joinpoint/testdata"
	assert := assert.New(t)
	packages := parser.
		New("testdata", project).
		Parse("test")
	functions := GetFunctions(packages)
	cases := map[string]struct {
		path       string
		paramsLen  int
		resultsLen int
	}{
		"a": {
			path:       "test.a(string,string)",
			paramsLen:  1,
			resultsLen: 0,
		},
		"b": {
			path:       "test.b(string)",
			paramsLen:  1,
			resultsLen: 0,
		},
		"c": {
			path:       "test.c(string,int)",
			paramsLen:  2,
			resultsLen: 0,
		},
		"d": {
			path:       "test.d(string,string,int)",
			paramsLen:  2,
			resultsLen: 0,
		},

		"e": {
			path:       "test.e([]string)",
			paramsLen:  1,
			resultsLen: 0,
		},
		"f": {
			path:       "test.f(string,[]string)",
			paramsLen:  2,
			resultsLen: 0,
		},
		"g": {
			path:       "test.g([]*string)",
			paramsLen:  1,
			resultsLen: 0,
		},
		"h": {
			path:       "test.h([]func([]string))",
			paramsLen:  1,
			resultsLen: 0,
		},
		"i": {
			path:       "test.i([]*github.com/wesovilabs/goa/joinpoint/testdata/test/model.TestArgument)",
			paramsLen:  1,
			resultsLen: 0,
		},
		"j": {
			path:       "test.*Element.j(string)map[string]interface{}",
			paramsLen:  1,
			resultsLen: 1,
		},
		"k": {
			path:       "test.Element.k()",
			paramsLen:  0,
			resultsLen: 0,
		},
		"l": {
			path:       "test/model.TestArgument.l()",
			paramsLen:  0,
			resultsLen: 0,
		},
	}

	for _, function := range functions.List() {

		if c, ok := cases[function.Name()]; !ok {
			continue
		} else {
			assert.Equal(c.paramsLen, len(function.ParamsList()))
			assert.Equal(c.resultsLen, len(function.ResultsList()))
			assert.Equal(c.path, function.path)
		}
	}

}
