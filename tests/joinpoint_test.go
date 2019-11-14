package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/joinpoint"
	"testing"
)

type jpTest struct {
	path       string
	fileName   string
	imports    map[string]string
	paramsLen  int
	resultsLen int
}

var joinpoints = map[string]map[string]map[string]*jpTest{
	"storage": {
		"": {
			"SetUpDatabase": {
				fileName: "storage",
				path:     "storage.SetUpDatabase()",
				imports: map[string]string{
					"errors": "",
					"github.com/wesovilabs/goa/examples/model":          "",
					"github.com/wesovilabs/goa/examples/storage/helper": "",
				},
				paramsLen:  0,
				resultsLen: 0,
			},

			"InsertPerson": {
				fileName: "storage",
				path:     "storage.InsertPerson(*github.com/wesovilabs/goa/examples/model.Person,*github.com/wesovilabs/goa/examples/storage/helper.Test)error",
				imports: map[string]string{
					"errors": "",
					"github.com/wesovilabs/goa/examples/model":          "",
					"github.com/wesovilabs/goa/examples/storage/helper": "",
				},
				paramsLen:  2,
				resultsLen: 1,
			},
			"FindPerson": {
				path: "storage.FindPerson(string)(*github.com/wesovilabs/goa/examples/model.Person,error)",
				imports: map[string]string{
					"errors": "",
					"github.com/wesovilabs/goa/examples/model":          "",
					"github.com/wesovilabs/goa/examples/storage/helper": "",
				},
				paramsLen:  1,
				resultsLen: 2,
			},
			"DeletePerson": {
				path: "storage.DeletePerson(string)([]*github.com/wesovilabs/goa/examples/model.Person,error)",
				imports: map[string]string{
					"errors": "",
					"github.com/wesovilabs/goa/examples/model":          "",
					"github.com/wesovilabs/goa/examples/storage/helper": "",
				},
				paramsLen:  1,
				resultsLen: 2,
			},
			"ListPeople": {
				path: "storage.ListPeople()([]*github.com/wesovilabs/goa/examples/model.Person,error)",
				imports: map[string]string{
					"errors": "",
					"github.com/wesovilabs/goa/examples/model":          "",
					"github.com/wesovilabs/goa/examples/storage/helper": "",
				},
				paramsLen:  0,
				resultsLen: 2,
			},
		},
	},
	"helper": {

		"t": {
			"test": {
				path: "storage/helper.*Test.test(map[string]interface{})",
				imports: map[string]string{
					"math/rand": "",
					"time":      "",
				},
				paramsLen:  1,
				resultsLen: 0,
			},
		},

		"": {

			"RandomUID": {
				path: "storage/helper.RandomUID(int)string",
				imports: map[string]string{
					"math/rand": "",
					"time":      "",
				},
				paramsLen:  1,
				resultsLen: 1,
			},
		},
	},
	"main": {
		"": {
			"main": {
				path: "main.main()",
				imports: map[string]string{
					"fmt":                                        "",
					"github.com/wesovilabs/goa/api":              "",
					"github.com/wesovilabs/goa/api/advice":       "",
					"github.com/wesovilabs/goa/api/context":      "",
					"github.com/wesovilabs/goa/examples/advice":  "testAdvice",
					"github.com/wesovilabs/goa/examples/model":   "",
					"github.com/wesovilabs/goa/examples/storage": "",
				},
				paramsLen:  0,
				resultsLen: 0,
			},
			"Goa": {
				path: "main.Goa()*github.com/wesovilabs/goa/api.Goa",
				imports: map[string]string{
					"fmt":                                        "",
					"github.com/wesovilabs/goa/api":              "",
					"github.com/wesovilabs/goa/api/advice":       "",
					"github.com/wesovilabs/goa/api/context":      "",
					"github.com/wesovilabs/goa/examples/advice":  "testAdvice",
					"github.com/wesovilabs/goa/examples/model":   "",
					"github.com/wesovilabs/goa/examples/storage": "",
				},
				paramsLen:  0,
				resultsLen: 1,
			},
		},
	},
	"advice": {
		"": {
			"NewComplexAround": {
				path: "advice.NewComplexAround(string,github.com/wesovilabs/goa/examples/advice.Attribute,interface{})func()github.com/wesovilabs/goa/api.Around",
				imports: map[string]string{
					"github.com/wesovilabs/goa/api":         "",
					"github.com/wesovilabs/goa/api/context": "",
				},
				paramsLen:  3,
				resultsLen: 1,
			},
			"NewComplexBefore": {
				path: "advice.NewComplexBefore(*github.com/wesovilabs/goa/examples/advice.Attribute)func()github.com/wesovilabs/goa/api.Before",
				imports: map[string]string{
					"github.com/wesovilabs/goa/api":         "",
					"github.com/wesovilabs/goa/api/context": "",
				},
				paramsLen:  1,
				resultsLen: 1,
			},
		},
	},
	"model": {
		"p": {
			"FullName": {
				path: "model.*Person.FullName()string",
				imports: map[string]string{
					"fmt": "",
					"github.com/wesovilabs/goa/examples/advice": "",
				},
				paramsLen:  0,
				resultsLen: 1,
			},
			"Apply": {
				path: "model.*Person.Apply([]github.com/wesovilabs/goa/examples/advice.Attribute)func(string,int)",
				imports: map[string]string{
					"fmt": "",
					"github.com/wesovilabs/goa/examples/advice": "",
				},
				paramsLen:  1,
				resultsLen: 1,
			},
		},
	},
}

func Test_JoinPoint(t *testing.T) {
	assert := assert.New(t)
	packages := testPackages()
	assert.NotNil(packages)
	jps := joinpoint.GetJoinPoints(pkg, packages)
	for _, jp := range jps.List() {
		jpType := ""
		if jp.GetRecv() != nil {
			jpType = jp.GetRecv().List[0].Names[0].String()
		}
		expected := joinpoints[jp.Pkg()][jpType][jp.Name()]
		if expected == nil {
			fmt.Println(jp.Pkg())
			fmt.Println(jpType)
			fmt.Println(jp.Name())
			panic("")
		}
		assert.NotNil(expected)
		assert.Equal(expected.path, jp.Path())
		assert.Len(jp.ImportSpecs(), len(expected.imports))
		for _, jpImport := range jp.ImportSpecs() {
			jpImportValue := jpImport.Path.Value[1 : len(jpImport.Path.Value)-1]
			value := expected.imports[jpImportValue]
			if jpImport.Name == nil {
				assert.Empty(value)
			} else {
				assert.Equal(value, jpImport.Name.String())
			}
		}
		for path := range expected.imports {
			found := false
			for _, jpImport := range jp.ImportSpecs() {
				if jpImport.Path.Value == `"`+path+`"` {
					found = true
					break
				}
			}
			assert.True(found)
		}
		assert.Len(jp.ParamsList(), expected.paramsLen)
		assert.Len(jp.ResultsList(), expected.resultsLen)

	}
}
