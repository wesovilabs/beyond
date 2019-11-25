package context

import (
	"reflect"
)

// Args struct
type Args struct {
	items []*Arg
}

// ForEach executes the given function for each arg
func (args *Args) ForEach(fn func(int, *Arg)) {
	for index, arg := range args.items {
		fn(index, arg)
	}
}

// Find executes find the first argument that matche with the given expression
func (args *Args) Find(fn func(int, *Arg) bool) (int, *Arg) {
	for index, arg := range args.items {
		if fn(index, arg) {
			return index, arg
		}
	}

	return -1, nil
}

// Count returns the number of elements in the list
func (args *Args) Count() int {
	return len(args.items)
}

// At returns the argument in the given position
func (args *Args) At(index int) *Arg {
	if len(args.items) > index && index >= 0 {
		return args.items[index]
	}

	return nil
}

// Get returns the argument with given name
func (args *Args) Get(name string) *Arg {
	for _, arg := range args.items {
		if arg.name == name {
			return arg
		}
	}

	return nil
}

// Set set a value for the given argument
func (args *Args) Set(name string, value interface{}) {
	for _, arg := range args.items {
		if arg.name == name {
			arg.value = value
			arg.kind = reflect.TypeOf(value).Name()

			return
		}
	}

	args.items = append(args.items, &Arg{
		name:  name,
		value: value,
		kind:  reflect.TypeOf(value).Name(),
	})
}

// SetWithType set a value for the given argument
func (args *Args) SetWithType(name string, value interface{}, argType string) {
	for _, arg := range args.items {
		if arg.name == name {
			arg.value = value
			arg.kind = argType

			return
		}
	}

	args.items = append(args.items, &Arg{
		name:  name,
		value: value,
		kind:  argType,
	})
}

// SetAt set a value in the given position
func (args *Args) SetAt(index int, value interface{}) {
	if args.Count() > 0 && index >= 0 && index < args.Count() {
		args.items[index].value = value
		args.items[index].kind = reflect.TypeOf(value).String()
	}
}
