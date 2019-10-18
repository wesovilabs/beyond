package context

import "reflect"

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

// Kind returns the argument kind
func (a *Arg) Kind() reflect.Type {
	return a.kind
}

// Update updates the argument value
func (a *Arg) Update(value interface{}) {
	a.value = value
}
