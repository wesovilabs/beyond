package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/beyond/imports"
	"testing"
)

var expectedImports = map[string]map[string]map[string]string{
	"cmd": {
		"../testdata/cmd/main.go": {
			"github.com/wesovilabs/beyond/api/context":      "context",
			"github.com/wesovilabs/beyond/testdata/advice":  "testAdvice",
			"github.com/wesovilabs/beyond/testdata/model":   "model",
			"github.com/wesovilabs/beyond/testdata/storage": "storage",
			"fmt":                              "fmt",
			"github.com/wesovilabs/beyond/api": "api",
			"github.com/wesovilabs/beyond/api/advice": "advice",
		},
	},
	"advice": {
		"../testdata/advice/custom.go": {
			"github.com/wesovilabs/beyond/api":         "api",
			"github.com/wesovilabs/beyond/api/context": "context",
		},
	},
	"model": {
		"../testdata/model/person.go": {
			"fmt": "fmt",
			"github.com/wesovilabs/beyond/testdata/advice": "advice",
		},
	},
	"storage": {
		"../testdata/storage/db.go": {
			"fmt": "fmt",
			"github.com/wesovilabs/beyond/testdata/storage/helper": "helper",
			"errors": "errors",
			"github.com/wesovilabs/beyond/testdata/model": "model",
		},
	},
	"storage/helper": {
		"../testdata/storage/helper/uid.go": {
			"math/rand": "rand",
			"time":      "time",
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
			fmt.Println("  " + fileName)
			imports := imports.GetImports(file)
			for importName, i := range imports {
				fmt.Println("    " + importName + ":" + i)
				assert.Equal(expectedImports[pkgName][fileName][importName], i)

			}
		}
	}
}
