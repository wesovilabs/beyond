package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/beyond/advice"
	"github.com/wesovilabs/beyond/joinpoint"
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
					"github.com/wesovilabs/beyond/testdata/model":          "",
					"github.com/wesovilabs/beyond/testdata/storage/helper": "",
				},
				paramsLen:  0,
				resultsLen: 0,
			},

			"InsertPerson": {
				fileName: "storage",
				path:     "storage.InsertPerson(*github.com/wesovilabs/beyond/testdata/model.Person,*github.com/wesovilabs/beyond/testdata/storage/helper.Test)error",
				imports: map[string]string{
					"errors": "",
					"github.com/wesovilabs/beyond/testdata/model":          "",
					"github.com/wesovilabs/beyond/testdata/storage/helper": "",
				},
				paramsLen:  2,
				resultsLen: 1,
			},
			"FindPerson": {
				path: "storage.FindPerson(string)(*github.com/wesovilabs/beyond/testdata/model.Person,error)",
				imports: map[string]string{
					"errors": "",
					"github.com/wesovilabs/beyond/testdata/model":          "",
					"github.com/wesovilabs/beyond/testdata/storage/helper": "",
				},
				paramsLen:  1,
				resultsLen: 2,
			},
			"DeletePerson": {
				path: "storage.DeletePerson(string)([]*github.com/wesovilabs/beyond/testdata/model.Person,error)",
				imports: map[string]string{
					"errors": "",
					"github.com/wesovilabs/beyond/testdata/model":          "",
					"github.com/wesovilabs/beyond/testdata/storage/helper": "",
				},
				paramsLen:  1,
				resultsLen: 2,
			},
			"ListPeople": {
				path: "storage.ListPeople()([]*github.com/wesovilabs/beyond/testdata/model.Person,error)",
				imports: map[string]string{
					"errors": "",
					"github.com/wesovilabs/beyond/testdata/model":          "",
					"github.com/wesovilabs/beyond/testdata/storage/helper": "",
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
					"fmt":                              "",
					"github.com/wesovilabs/beyond/api": "",
					"github.com/wesovilabs/beyond/api/advice":       "",
					"github.com/wesovilabs/beyond/api/context":      "",
					"github.com/wesovilabs/beyond/testdata/advice":  "testAdvice",
					"github.com/wesovilabs/beyond/testdata/model":   "",
					"github.com/wesovilabs/beyond/testdata/storage": "",
				},
				paramsLen:  0,
				resultsLen: 0,
			},
			"Beyond": {
				path: "main.Beyond()*github.com/wesovilabs/beyond/api.Beyond",
				imports: map[string]string{
					"fmt":                              "",
					"github.com/wesovilabs/beyond/api": "",
					"github.com/wesovilabs/beyond/api/advice":       "",
					"github.com/wesovilabs/beyond/api/context":      "",
					"github.com/wesovilabs/beyond/testdata/advice":  "testAdvice",
					"github.com/wesovilabs/beyond/testdata/model":   "",
					"github.com/wesovilabs/beyond/testdata/storage": "",
				},
				paramsLen:  0,
				resultsLen: 1,
			},
		},
	},
	"advice": {
		"": {
			"NewComplexAround": {
				path: "advice.NewComplexAround(string,github.com/wesovilabs/beyond/testdata/advice.Attribute,interface{})func()github.com/wesovilabs/beyond/api.Around",
				imports: map[string]string{
					"github.com/wesovilabs/beyond/api":         "",
					"github.com/wesovilabs/beyond/api/context": "",
				},
				paramsLen:  3,
				resultsLen: 1,
			},
			"NewComplexBefore": {
				path: "advice.NewComplexBefore(*github.com/wesovilabs/beyond/testdata/advice.Attribute)func()github.com/wesovilabs/beyond/api.Before",
				imports: map[string]string{
					"github.com/wesovilabs/beyond/api":         "",
					"github.com/wesovilabs/beyond/api/context": "",
				},
				paramsLen:  1,
				resultsLen: 1,
			},
			"NewTimerAdvice": {
				path: "advice.NewTimerAdvice(github.com/wesovilabs/beyond/testdata/advice.TimerMode)func()github.com/wesovilabs/beyond/api.Around",
				imports: map[string]string{
					"github.com/wesovilabs/beyond/api":         "",
					"github.com/wesovilabs/beyond/api/context": "",
					"strings": "",
					"time":    "",
					"fmt":     "",
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
					"github.com/wesovilabs/beyond/testdata/advice": "testAdvice",
				},
				paramsLen:  0,
				resultsLen: 1,
			},
			"Apply": {
				path: "model.*Person.Apply([]github.com/wesovilabs/beyond/testdata/advice.Attribute)func(string,int)",
				imports: map[string]string{
					"fmt": "",
					"github.com/wesovilabs/beyond/testdata/advice": "testAdvice",
				},
				paramsLen:  1,
				resultsLen: 1,
			},
			"Other": {
				path: "model.*Person.Other()string",
				imports: map[string]string{
					"fmt": "",
					"github.com/wesovilabs/beyond/testdata/advice": "testAdvice",
				},
				paramsLen:  0,
				resultsLen: 1,
			},
		},
	},
}

func Test_JoinPoint(t *testing.T) {
	assert := assert.New(t)
	packages := testPackages()
	assert.NotNil(packages)
	advices := advice.GetAdvices(packages)
	excluds := advice.GetExcludePaths(packages)
	jps := joinpoint.GetJoinPoints(pkg, advices, excluds, packages)
	for _, jp := range jps.List() {
		jpType := ""
		if jp.GetRecv() != nil {
			jpType = jp.GetRecv().List[0].Names[0].String()
		}
		fmt.Println(jp.Pkg())
		fmt.Println(jpType)
		fmt.Println(jp.Name())
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
