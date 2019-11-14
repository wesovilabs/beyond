package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/imports"
	"testing"
)

var expectedImports = map[string]map[string]map[string]string{
	"cmd": {
		"../examples/cmd/main.go": {
			"github.com/wesovilabs/goa/api/context": "context",
			"github.com/wesovilabs/goa/examples/advice": "testAdvice",
			"github.com/wesovilabs/goa/examples/model": "model",
			"github.com/wesovilabs/goa/examples/storage": "storage",
			"fmt": "fmt",
			"github.com/wesovilabs/goa/api": "api",
			"github.com/wesovilabs/goa/api/advice": "advice",
		},
	},
	"advice":{
		"../examples/advice/custom.go":{
			"github.com/wesovilabs/goa/api": "api",
			"github.com/wesovilabs/goa/api/context": "context",
		},
	},
	"model":{
		"../examples/model/person.go":{
			"fmt": "fmt",
		},
	},
	"storage":{
		"../examples/storage/db.go":{
			"fmt": "fmt",
			"github.com/wesovilabs/goa/examples/storage/helper": "helper",
			"errors":"errors",
			"github.com/wesovilabs/goa/examples/model":"model",
		},
	},
	"storage/helper":{
		"../examples/storage/helper/uid.go":{
			"math/rand": "rand",
			"time":"time",
		},
	},
}

func Test_Imports(t *testing.T) {
	assert := assert.New(t)
	packages := testPackages()
	assert.NotNil(packages)
	for pkgName, pkg := range packages {
		fmt.Println(pkgName)
		for fileName, file := range pkg.Node().Files {
			fmt.Println("  "+fileName)
			imports := imports.GetImports(file)
			for importName, i := range imports {
				fmt.Println("    "+importName+":"+i)
				assert.Equal(expectedImports[pkgName][fileName][importName],i)

			}
		}
	}
}
