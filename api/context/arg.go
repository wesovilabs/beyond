package context

import (
	"reflect"
)

// Arg contains attributes to efine an argument
type Arg struct {
	name  string
	kind  reflect.Type
	value interface{}
}

// NewArg constructs an instance of arg
func NewArg(name string, value interface{}) *Arg {
	return &Arg{
		name:  name,
		value: value,
		kind:  reflect.TypeOf(value),
	}
}

// Name returns the argument name
func (a *Arg) Name() string {
	return a.name
}

// Value returns the argument value
func (a *Arg) Value() interface{} {
	return a.value
}

// Kind returns the argument type
func (a *Arg) Kind() reflect.Type {
	return a.kind
}

// Update updates the argument value
func (a *Arg) update(value interface{}) {
	a.value = value
	a.kind = reflect.TypeOf(value)
}

// AsString return argument value as string
func (a *Arg) AsString() string {
	if a != nil && a.value != nil {
		return a.value.(string)
	}

	return ""
}

// AsStringPtr return argument value as pointer to string
func (a *Arg) AsStringPtr() *string {
	if a.value != nil {
		return a.value.(*string)
	}

	return nil
}

// AsInt return argument value as int
func (a *Arg) AsInt() int {
	if a.value != nil {
		return a.value.(int)
	}

	return 0
}

// AsIntPtr return argument value as pointer of int
func (a *Arg) AsIntPtr() *int {
	if a.value != nil {
		return a.value.(*int)
	}

	return nil
}

func (a *Arg) Is(t reflect.Type) bool {
	return a.kind == t
}

func (a *Arg) IsError() bool {
	if _, ok := a.value.(error); ok {
		return true
	}

	return false
}

// AsInt16 return argument value as int16
func (a *Arg) AsInt16() int16 {
	if a.value != nil {
		return a.value.(int16)
	}

	return 0
}

// AsInt16Ptr return argument value as pointer of int16
func (a *Arg) AsInt16Ptr() *int16 {
	if a.value != nil {
		return a.value.(*int16)
	}

	return nil
}

// AsInt8 return argument value as int8
func (a *Arg) AsInt8() int8 {
	if a.value != nil {
		return a.value.(int8)
	}

	return 0
}

// AsInt8Ptr return argument value as pointer of int8
func (a *Arg) AsInt8Ptr() *int8 {
	if a.value != nil {
		return a.value.(*int8)
	}

	return nil
}

// AsInt32 return argument value as int32
func (a *Arg) AsInt32() int32 {
	if a.value != nil {
		return a.value.(int32)
	}

	return 0
}

// AsInt32Ptr return argument value as pointer of int32
func (a *Arg) AsInt32Ptr() *int32 {
	if a.value != nil {
		return a.value.(*int32)
	}

	return nil
}

// AsInt64 return argument value as int64
func (a *Arg) AsInt64() int64 {
	if a.value != nil {
		return a.value.(int64)
	}

	return 0
}

// AsInt64Ptr return argument value as pointer of int64
func (a *Arg) AsInt64Ptr() *int64 {
	if a.value != nil {
		return a.value.(*int64)
	}

	return nil
}

// AsFloat32 return argument value as float32
func (a *Arg) AsFloat32() float32 {
	if a.value != nil {
		return a.value.(float32)
	}

	return 0
}

// AsFloat32Ptr return argument value as pointer of float32
func (a *Arg) AsFloat32Ptr() *float32 {
	if a.value != nil {
		return a.value.(*float32)
	}

	return nil
}

// AsFloat64 return argument value as float64
func (a *Arg) AsFloat64() float64 {
	if a.value != nil {
		return a.value.(float64)
	}

	return 0
}

// AsFloat64Ptr return argument value as pointer of float64
func (a *Arg) AsFloat64Ptr() *float64 {
	if a.value != nil {
		return a.value.(*float64)
	}

	return nil
}

// AsBool return argument value as bool
func (a *Arg) AsBool() bool {
	if a.value != nil {
		return a.value.(bool)
	}

	return false
}

// AsBoolPtr return argument value as pointer of bool
func (a *Arg) AsBoolPtr() *bool {
	if a.value != nil {
		return a.value.(*bool)
	}

	return nil
}
