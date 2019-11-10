package context

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Arg(t *testing.T) {
	assert := assert.New(t)
	arg := NewArg("name", "John")
	assert.EqualValues("name", arg.Name())
	assert.EqualValues("John", arg.Value())
}

func Test_AsInt(t *testing.T) {
	assert := assert.New(t)
	var value = 32
	arg := &Arg{
		value: value,
	}
	assert.EqualValues(value, arg.AsInt())
	arg = &Arg{}
	assert.EqualValues(0, arg.AsInt())
}

func Test_AsInt8(t *testing.T) {
	assert := assert.New(t)
	var value int8 = 32
	arg := &Arg{
		value: value,
	}
	assert.EqualValues(value, arg.AsInt8())
	arg = &Arg{}
	assert.EqualValues(0, arg.AsInt8())
}

func Test_AsInt16(t *testing.T) {
	assert := assert.New(t)
	var value int16 = 32
	arg := &Arg{
		value: value,
	}
	assert.EqualValues(value, arg.AsInt16())
	arg = &Arg{}
	assert.EqualValues(0, arg.AsInt16())
}

func Test_AsInt32(t *testing.T) {
	assert := assert.New(t)
	var value int32 = 32
	arg := &Arg{
		value: value,
	}
	assert.EqualValues(value, arg.AsInt32())
	arg = &Arg{}
	assert.EqualValues(0, arg.AsInt32())
}

func Test_AsInt64(t *testing.T) {
	assert := assert.New(t)
	var value int64 = 32
	arg := &Arg{
		value: value,
	}
	assert.EqualValues(value, arg.AsInt64())
	arg = &Arg{}
	assert.EqualValues(0, arg.AsInt64())
}

func Test_AsIntPtr(t *testing.T) {
	assert := assert.New(t)
	var value = 32
	arg := &Arg{
		value: &value,
	}
	assert.EqualValues(&value, arg.AsIntPtr())
	arg = &Arg{}
	assert.Nil(arg.AsIntPtr())
}

func Test_AsIntPtr8(t *testing.T) {
	assert := assert.New(t)
	var value int8 = 32
	arg := &Arg{
		value: &value,
	}
	assert.EqualValues(&value, arg.AsInt8Ptr())
	arg = &Arg{}
	assert.Nil(arg.AsInt8Ptr())
}

func Test_AsIntPtr16(t *testing.T) {
	assert := assert.New(t)
	var value int16 = 32
	arg := &Arg{
		value: &value,
	}
	assert.EqualValues(&value, arg.AsInt16Ptr())
	arg = &Arg{}
	assert.Nil(arg.AsInt16Ptr())
}

func Test_AsIntPtr32(t *testing.T) {
	assert := assert.New(t)
	var value int32 = 32
	arg := &Arg{
		value: &value,
	}
	assert.EqualValues(&value, arg.AsInt32Ptr())
	arg = &Arg{}
	assert.Nil(arg.AsInt32Ptr())
}

func Test_AsIntPtr64(t *testing.T) {
	assert := assert.New(t)
	var value int64 = 32
	arg := &Arg{
		value: &value,
	}
	assert.EqualValues(&value, arg.AsInt64Ptr())
	arg = &Arg{}
	assert.Nil(arg.AsInt64Ptr())
}

func Test_AsString(t *testing.T) {
	assert := assert.New(t)
	var value = "hey"
	arg := &Arg{
		value: value,
	}
	assert.EqualValues(value, arg.AsString())
	arg = &Arg{}
	assert.EqualValues("", arg.AsString())
}

func Test_AsStringPtr(t *testing.T) {
	assert := assert.New(t)
	var value = "hey"
	arg := &Arg{
		value: &value,
	}
	assert.EqualValues(&value, arg.AsStringPtr())
	arg = &Arg{}
	assert.Nil(arg.AsStringPtr())
}

func Test_AsBool(t *testing.T) {
	assert := assert.New(t)
	var value = true
	arg := &Arg{
		value: value,
	}
	assert.EqualValues(value, arg.AsBool())
	arg = &Arg{}
	assert.EqualValues(false, arg.AsBool())
}

func Test_AsBoolPtr(t *testing.T) {
	assert := assert.New(t)
	var value = true
	arg := &Arg{
		value: &value,
	}
	assert.EqualValues(&value, arg.AsBoolPtr())
	arg = &Arg{}
	assert.Nil(arg.AsBoolPtr())
}

func TestArg_AsFloat32(t *testing.T) {
	assert := assert.New(t)
	var value float32 = 32
	arg := &Arg{
		value: value,
	}
	assert.EqualValues(value, arg.AsFloat32())
	arg = &Arg{}
	assert.EqualValues(0, arg.AsFloat32())
}

func TestArg_AsFloat64(t *testing.T) {
	assert := assert.New(t)
	var value float64 = 64
	arg := &Arg{
		value: value,
	}
	assert.EqualValues(value, arg.AsFloat64())
	arg = &Arg{}
	assert.EqualValues(0, arg.AsFloat64())
}

func TestArg_AsFloat32Ptr(t *testing.T) {
	assert := assert.New(t)
	var value float32 = 32
	arg := &Arg{
		value: &value,
	}
	assert.EqualValues(&value, arg.AsFloat32Ptr())
	arg = &Arg{}
	assert.Nil(arg.AsFloat32Ptr())
}

func TestArg_AsFloat64Ptr(t *testing.T) {
	assert := assert.New(t)
	var value float64 = 64
	arg := &Arg{
		value: &value,
	}
	assert.EqualValues(&value, arg.AsFloat64Ptr())
	arg = &Arg{}
	assert.Nil(arg.AsFloat64Ptr())
}
