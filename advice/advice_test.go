package advice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/beyond/advice/internal"
	"regexp"
	"testing"
)

func TestMatch(t *testing.T) {
	fmt.Printf("[TEST] %s\n", t.Name())
	cases := []struct {
		regExp    *regexp.Regexp
		matches   []string
		noMatches []string
	}{
		{
			regExp: internal.NormalizePointcut("*.set*(*)..."),
			matches: []string{
				"a.setPerson(string)int",
				"a.setElement(int)",
				"a/b.setCat(string)*int",
			},
			noMatches: []string{
				"a.unsetPerson(string)int",
				"a.list(string)(int,*string)",
				"a/b.unsetCat(string)(int,*string)",
			},
		},
		{
			regExp: internal.NormalizePointcut("*.*(*)..."),
			matches: []string{
				"a.b(string)int",
				"a.b(string)(int,*string)",
				"a/b.b(string)(int,*string)",
			},
			noMatches: []string{
				"a/b.c.d(string)",
				"a.c.d(string)",
				"a/b.b()(int,*string)",
			},
		},
		{
			regExp: internal.NormalizePointcut("model/*.*(*)..."),
			matches: []string{
				"model/a.b(string)int",
				"model/a/b.b(string)(int,*string)",
			},
			noMatches: []string{
				"model.b(string)(int,*string)",
			},
		},
		{
			regExp: internal.NormalizePointcut("model*.*(*)..."),
			matches: []string{
				"model/a.b(string)int",
				"model/a/b.b(string)(int,*string)",
				"model.b(string)(int,*string)",
			},
			noMatches: []string{
				"a/model/a.b(string)int",
			},
		},
		{
			regExp: internal.NormalizePointcut("model*.set*(*)..."),
			matches: []string{
				"model/a.setB(string)int",
				"model/a/b.setElement(string)(int,*string)",
				"model.set(string)(int,*string)",
			},
			noMatches: []string{
				"a/model/a.b(string)int",
				"model/aset(string)int",
			},
		},
		{
			regExp: internal.NormalizePointcut("model*.*set*(*)..."),
			matches: []string{
				"model/a.setB(string)int",
				"model/a/b.setElement(string)(int,*string)",
				"model/a/b.unsetElement(string)(int,*string)",
				"model.set(string)(int,*string)",
				"model.setPerson(string)(int,*string)",
			},
			noMatches: []string{
				"a/model/a.b(string)int",
			},
		},
		{
			regExp: internal.NormalizePointcut("model.set(*)"),
			matches: []string{
				"model.set(string)",
				"model.set(int)",
				"model.set(*int)",
			},
			noMatches: []string{
				"model.set(string,int)int",
				"model.set(*int)int",
				"model.set()",
				"model.set(int)*int",
			},
		},
		{
			regExp: internal.NormalizePointcut("model.set(*)"),
			matches: []string{
				"model.set(string)",
				"model.set(int)",
				"model.set(*int)",
			},
			noMatches: []string{
				"model.set(string,int)int",
				"model.set(*int)int",
				"model.set()",
				"model.set(int)*int",
			},
		},
		{
			regExp: internal.NormalizePointcut("model.obj.set(*)"),
			matches: []string{
				"model.obj.set(string)",
			},
			noMatches: []string{
				"model.obj.set(string,int)",
				"model.object.set(string)",
				"model.myobj.set(string)",
			},
		},
		{
			regExp: internal.NormalizePointcut("model.*obj*.set(*)"),
			matches: []string{
				"model.obj.set(string)",
				"model.object.set(string)",
				"model.myobj.set(string)",
			},
			noMatches: []string{
				"model.obj.set(string)(string,string)",
				"model.object.set(string)int",
				"model.myobj.set(int,string)",
			},
		},
		{
			regExp: internal.NormalizePointcut("model.obj*.set(*)"),
			matches: []string{
				"model.obj.set(string)",
				"model.object.set(string)",
			},
			noMatches: []string{
				"model.unobj.set(string)",
				"model.Obj.set(string)",
			},
		},
		{
			regExp: internal.NormalizePointcut("*.*(...)..."),
			matches: []string{
				"a.b(string)int",
				"a.b(string)(int,*string)",
				"a/b.b(string)(int,*string)",
				"a/b.b()(int,*string)",
				"a/b.b()(int,*github.com/projec/repo.model.Person)",
			},
			noMatches: []string{
				"a/b.c.d(string)",
				"a.c.d(string)",
			},
		},
		{
			regExp: internal.NormalizePointcut("model.obj.set(*)"),
			matches: []string{
				"model.obj.set(string)",
				"model.obj.set(*int32)",
				"model.obj.set(func(string,int))",
				"model.obj.set(*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(map[string]interface{})",
			},
			noMatches: []string{
				"model.obj.set(string,int)",
				"model.obj.set(string,map[string]interface{})",
				"model.obj.set(string)*int",
			},
		},
		{
			regExp: internal.NormalizePointcut("model.obj.set(...)"),
			matches: []string{
				"model.obj.set()",
				"model.obj.set(string)",
				"model.obj.set(*int32)",
				"model.obj.set(func(string,int))",
				"model.obj.set(*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(map[string]interface{})",
				"model.obj.set(string,int)",
				"model.obj.set(string,map[string]interface{})",
				"model.obj.set(string)",
			},
			noMatches: []string{},
		},
		{
			regExp: internal.NormalizePointcut("model.obj.set(string,...)"),
			matches: []string{
				"model.obj.set(string,*int32)",
				"model.obj.set(string,func(string,int))",
				"model.obj.set(string,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(string,map[string]interface{})",
				"model.obj.set(string,int)",
				"model.obj.set(string,map[string]interface{})",
			},
			noMatches: []string{
				"model.obj.set(string)",
				"model.obj.set(*string,int)",
				"model.obj.set(int,string)",
			},
		},
		{
			regExp: internal.NormalizePointcut("model.obj.set(...,github.com/wesovilabs/beyond.model.Person)"),
			matches: []string{
				"model.obj.set(string,*int32,github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(string,func(string,int),github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(string,github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(*string,github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(string,map[string]interface{},github.com/wesovilabs/beyond.model.Person)",
			},
			noMatches: []string{
				"model.obj.set(string)",
				"model.obj.set(*string,int)",
				"model.obj.set(int,string)",
			},
		},
		{
			regExp: internal.NormalizePointcut("model.obj.set(int,...,*github.com/wesovilabs/beyond.model.Person)"),
			matches: []string{
				"model.obj.set(int,string,*int32,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(int,string,func(string,int),*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(int,string,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(int,*string,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(int,string,map[string]interface{},*github.com/wesovilabs/beyond.model.Person)",
			},
			noMatches: []string{
				"model.obj.set(int,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(*int,*string,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(int,string,map[string]interface{},github.com/wesovilabs/beyond.model.Person)",
			},
		},

		{
			regExp: internal.NormalizePointcut("model.obj.set()..."),
			matches: []string{
				"model.obj.set()",
				"model.obj.set()string",
				"model.obj.set()*int32",
				"model.obj.set()func(string,int)",
				"model.obj.set()*github.com/wesovilabs/beyond.model.Person",
				"model.obj.set()map[string]interface{}",
			},
			noMatches: []string{},
		},
		{
			regExp: internal.NormalizePointcut("model.obj.set()(string,...)"),
			matches: []string{
				"model.obj.set()(string,*int32)",
				"model.obj.set()(string,func(string,int))",
				"model.obj.set()(string,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set()(string,map[string]interface{})",
				"model.obj.set()(string,int)",
				"model.obj.set()(string,map[string]interface{})",
			},
			noMatches: []string{
				"model.obj.set()(string)",
				"model.obj.set()(*string,int)",
				"model.obj.set()(int,string)",
			},
		},
		{
			regExp: internal.NormalizePointcut("model.obj.set()(...,github.com/wesovilabs/beyond.model.Person)"),
			matches: []string{
				"model.obj.set()(string,*int32,github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set()(string,func(string,int),github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set()(string,github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set()(*string,github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set()(string,map[string]interface{},github.com/wesovilabs/beyond.model.Person)",
			},
			noMatches: []string{
				"model.obj.set()(string)",
				"model.obj.set()(*string,int)",
				"model.obj.set()(int,string)",
			},
		},
		{
			regExp: internal.NormalizePointcut("model.obj.set()(int,...,*github.com/wesovilabs/beyond.model.Person)"),
			matches: []string{
				"model.obj.set()(int,string,*int32,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set()(int,string,func(string,int),*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set()(int,string,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set()(int,*string,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set()(int,string,map[string]interface{},*github.com/wesovilabs/beyond.model.Person)",
			},
			noMatches: []string{
				"model.obj.set()(int,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set()(*int,*string,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set()(int,string,map[string]interface{},github.com/wesovilabs/beyond.model.Person)",
			},
		},
		{
			regExp: internal.NormalizePointcut("model.obj.set(func()string)(int,...,*github.com/wesovilabs/beyond.model.Person)"),
			matches: []string{
				"model.obj.set(func()string)(int,string,*int32,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(func()string)(int,string,func(string,int),*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(func()string)(int,string,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(func()string)(int,*string,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(func()string)(int,string,map[string]interface{},*github.com/wesovilabs/beyond.model.Person)",
			},
			noMatches: []string{
				"model.obj.set()(int,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set()(*int,*string,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set()(int,string,map[string]interface{},github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(func()string)(int,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(func()string)(*int,*string,*github.com/wesovilabs/beyond.model.Person)",
				"model.obj.set(func()string)(int,string,map[string]interface{},github.com/wesovilabs/beyond.model.Person)",
			},
		},
	}

	assert := assert.New(t)

	for index, c := range cases {
		fmt.Printf("\nScenario number %v:\n", index)

		def := &Advice{
			regExp: c.regExp,
		}
		fmt.Printf(" %s\n", c.regExp)
		fmt.Printf("[matches]\n")
		for _, m := range c.matches {
			fmt.Printf("  %s\n", m)
			if !assert.True(def.Match(m)) {
				t.FailNow()
			}
		}
		fmt.Printf("[no matches]\n")
		for _, m := range c.noMatches {
			fmt.Printf("  %s\n", m)
			if !assert.False(def.Match(m)) {
				t.FailNow()
			}
		}
	}
	fmt.Println()
}

func Test_addImport(t *testing.T) {
	assert := assert.New(t)
	inv := adviceInvocation{
		imports: []string{"import1", "import2"},
	}
	inv.addImport("import2")
	assert.Len(inv.imports, 2)
}

func Test_Advice_Imports(t *testing.T) {
	assert := assert.New(t)
	advice := &Advice{
		call: &adviceInvocation{
			pkg:     "mypkg",
			imports: nil,
		},
	}
	res := advice.Imports()
	assert.Len(res, 1)
	assert.Equal("mypkg", res[0])
}

func Test_Advice_Match(t *testing.T) {
	assert := assert.New(t)
	advice := &Advice{}
	res := advice.Match("test")
	assert.False(res)
}
