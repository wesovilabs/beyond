package context

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_Args(t *testing.T) {
	args := Args{
		items: []*Arg{
			NewArg("name", "John"),
			NewArg("age", 20),
			NewArg("optional", nil),
		},
	}
	assert := assert.New(t)
	assert.Equal(args.Len(), 3)
	assert.Len(args.List(), 3)
	assert.False(args.IsEmpty())
	assert.Equal("name", args.items[0].Name())
	assert.Equal("John", args.items[0].Value())
	assert.Equal("age", args.items[1].Name())
	assert.Equal(20, args.items[1].Value())
	assert.Equal("optional", args.items[2].Name())
	assert.Equal(nil, args.items[2].Value())
	args.Set("optional", "none")
	assert.Equal("none", args.items[2].Value())
}

func TestArgsAt(t *testing.T) {
	assert := assert.New(t)
	arg1 := NewArg("name", "John")
	arg2 := NewArg("male", true)
	arg3 := NewArg("age", 20)
	args := &Args{
		items: []*Arg{arg1, arg2, arg3},
	}
	assert.Equal(arg1, args.At(0))
	assert.Equal(arg2, args.At(1))
	assert.Equal(arg3, args.At(2))
	assert.Nil(args.At(10))
}

func TestArgsGet(t *testing.T) {
	assert := assert.New(t)
	arg1 := NewArg("name", "John")
	arg2 := NewArg("male", true)
	arg3 := NewArg("age", 20)
	args := &Args{
		items: []*Arg{arg1, arg2, arg3},
	}
	assert.Nil(args.Get("unknown"))
	assert.Equal(arg2, args.Get("male"))
}

func TestArgsIsEmpty(t *testing.T) {
	assert := assert.New(t)
	arg1 := NewArg("name", "John")
	arg2 := NewArg("male", true)
	arg3 := NewArg("age", 20)
	args := &Args{
		items: []*Arg{arg1, arg2, arg3},
	}
	assert.False(args.IsEmpty())
	args.items = []*Arg{}
	assert.True(args.IsEmpty())

}

func TestArgsLen(t *testing.T) {
	assert := assert.New(t)
	arg1 := NewArg("name", "John")
	arg2 := NewArg("male", true)
	arg3 := NewArg("age", 20)
	args := &Args{
		items: []*Arg{arg1, arg2, arg3},
	}
	assert.Equal(3, args.Len())
	args.items = []*Arg{}
	assert.Equal(0, args.Len())
}

func TestArgsUpdateAt(t *testing.T) {
	assert := assert.New(t)
	arg1 := NewArg("name", "John")
	arg2 := NewArg("male", true)
	arg3 := NewArg("age", 20)
	args := &Args{
		items: []*Arg{arg1, arg2, arg3},
	}
	args.UpdateAt(2, "test")
	assert.Equal(args.items[2].value, "test")
}

func TestArgsSetAt(t *testing.T) {
	assert := assert.New(t)
	arg1 := NewArg("name", "John")
	arg2 := NewArg("male", true)
	arg3 := NewArg("age", 20)
	args := &Args{
		items: []*Arg{arg1, arg2, arg3},
	}
	assert.Equal(true, arg2.value)
	args.SetAt(1, false)
	assert.Equal(false, arg2.value)
	args.SetAt(1, "name")
	assert.Equal(reflect.TypeOf("name"), arg2.kind)
}

func TestArgsSet(t *testing.T) {
	assert := assert.New(t)
	arg1 := NewArg("name", "John")
	arg2 := NewArg("male", true)
	arg3 := NewArg("age", 20)
	args := &Args{
		items: []*Arg{arg1, arg2, arg3},
	}
	assert.Equal(true, arg2.value)
	args.Set("male", false)
	assert.Equal(false, arg2.value)
	args.Set("male", "name")
	assert.Equal(reflect.TypeOf("name"), arg2.kind)
	args.Set("other", 20)
	assert.Len(args.items, 4)
}
