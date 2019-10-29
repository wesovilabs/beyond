package wrapper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/aspect"
	"github.com/wesovilabs/goa/function"
	"github.com/wesovilabs/goa/internal/writer"
	"go/parser"
	"go/token"
	"testing"
)

func Test(t *testing.T) {
	assert := assert.New(t)
	packages, err := parser.ParseDir(&token.FileSet{}, "testdata", nil, parser.ParseComments)
	assert.Nil(err)
	assert.NotNil(packages)
	defs := aspect.GetDefinitions("", packages)
	functions := function.GetFunctions(packages)
	file, err := parser.ParseFile(&token.FileSet{}, "testdata/sample.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for _, d := range file.Decls {
		fmt.Println(d)
	}

	var function *function.Function
	for _, f := range functions.List() {
		if f.Name() == "CreatePerson" {
			function = f
			break
		}
	}
	defsMap := make(map[string]*aspect.Definition)
	for index, def := range defs.List() {
		defsMap[fmt.Sprintf("aspect%v", index)] = def
	}

	funcDecl := wrapperFuncDecl(function, defsMap)
	fmt.Println(funcDecl)
	file.Decls = append(file.Decls, funcDecl)
	writer.SaveNode(file, "testdata/generated2/main.go")
}
