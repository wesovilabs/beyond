package context

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Arg(t *testing.T){
	assert := assert.New(t)
	arg:=NewArg("name","John")
	assert.EqualValues("name",arg.Name())
	assert.EqualValues("John",arg.Value())
	arg.Update("Tom")
	assert.EqualValues("Tom",arg.Value())
}
