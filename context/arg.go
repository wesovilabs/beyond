package context

import "reflect"

// Arg contains attributes to efine an argument
type Arg struct {
	name  string
	kind  reflect.Type
	value interface{}
}

func NewArg(name string, value interface{}) *Arg {
	return &Arg{
		name:  name,
		value: value,
		kind:  reflect.TypeOf(value),
	}
}

func (a *Arg) Name() string {
	return a.name
}

func (a *Arg) Value() interface{} {
	return a.value
}

func (a *Arg) Kind() reflect.Type {
	return a.kind
}

func (a *Arg) Update(value interface{}) {
	a.value = value
}
