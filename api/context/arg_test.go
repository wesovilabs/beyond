package context

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_Arg(t *testing.T) {
	assert := assert.New(t)
	arg := NewArg("name", "John")
	assert.EqualValues("name", arg.Name())
	assert.EqualValues("John", arg.Value())
	assert.False(arg.IsError())
	assert.True(arg.Is(reflect.TypeOf("")))
	arg = NewArg("err", errors.New("error"))
	assert.EqualValues("err", arg.Name())
	assert.EqualValues("error", arg.Value().(error).Error())
	assert.True(arg.IsError())
	assert.Equal("*errors.fundamental", arg.Kind())
}
