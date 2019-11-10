package context

import (
	"github.com/stretchr/testify/assert"
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
	assert.Equal(args.Count(), 3)
	assert.False(args.isEmpty())
	assert.Equal("name", args.items[0].Name())
	assert.Equal("John", args.items[0].Value())
	assert.Equal("age", args.items[1].Name())
	assert.Equal(20, args.items[1].Value())
	assert.Equal("optional", args.items[2].Name())
	assert.Equal(nil, args.items[2].Value())

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
	assert.False(args.isEmpty())
	args.items = []*Arg{}
	assert.True(args.isEmpty())

}

func TestArgsLen(t *testing.T) {
	assert := assert.New(t)
	arg1 := NewArg("name", "John")
	arg2 := NewArg("male", true)
	arg3 := NewArg("age", 20)
	args := &Args{
		items: []*Arg{arg1, arg2, arg3},
	}
	assert.Equal(3, args.Count())
	args.items = []*Arg{}
	assert.Equal(0, args.Count())
}
