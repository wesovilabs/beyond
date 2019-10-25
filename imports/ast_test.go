package imports

import (
	"github.com/stretchr/testify/assert"
	"go/parser"
	"go/token"
	"testing"
)

func Test_GetImports(t *testing.T) {
	assert:=assert.New(t)
	file,err:=parser.ParseFile(&token.FileSet{}, "testdata/demo.go", nil, parser.ImportsOnly)
	assert.Nil(err)
	assert.NotNil(file)
	imports:=GetImports(file)
	assert.Len(imports,3)
	assert.Contains(imports,"encoding/json")
	assert.Contains(imports,"runtime")
	assert.Contains(imports,"time")
	assert.Equal("tt",imports["time"])
}
