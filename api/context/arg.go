package context

import (
	"fmt"
	"reflect"
)

// Arg contains attributes to efine an argument
type Arg struct {
	name  string
	kind  string
	value interface{}
}

// NewArg constructs an instance of arg
func NewArg(name string, value interface{}) *Arg {
	kind := ""

	if value != nil {
		fmt.Println(reflect.TypeOf(value).String())
		kind = reflect.TypeOf(value).String()
	}

	return &Arg{
		name:  name,
		value: value,
		kind:  kind,
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
func (a *Arg) Kind() string {
	return a.kind
}

// Is check if argument has the provided type
func (a *Arg) Is(t reflect.Type) bool {
	return a.kind == t.String()
}

// IsError check if argument is an error
func (a *Arg) IsError() bool {
	return a.kind == "error" || a.kind == "*errors.fundamental"
}
