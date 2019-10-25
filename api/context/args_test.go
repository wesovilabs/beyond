package context

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Args(t *testing.T){
	args:=Args{
		NewArg("name","John"),
		NewArg("age",20),
		NewArg("optional",nil),
	}
	assert:=assert.New(t)
	assert.Equal(args.Len(),3)
	assert.Len(args.List(),3)
	assert.False(args.IsEmpty())
	assert.Equal("name",args[0].Name())
	assert.Equal("John",args[0].Value())
	assert.Equal("age",args[1].Name())
	assert.Equal(20,args[1].Value())
	assert.Equal("optional",args[2].Name())
	assert.Equal(nil,args[2].Value())
	args.Set("optional","none")
	assert.Equal("none",args[2].Value())
	args.UpdateAt(2,"all")
	args.UpdateAt(20,"all")
	assert.Equal("all",args[2].Value())
	assert.Equal(20,args.Get("age"))
	assert.Equal(nil,args.Get("notFound"))
	assert.True(Args{}.IsEmpty())
}
