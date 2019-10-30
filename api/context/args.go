package context

// Args struct
type Args struct {
	items []*Arg
}

// List returns the list of arguments
func (args *Args) List() []*Arg {
	return args.items
}

// UpdateAt updates the value for the argument in the provided position
func (args *Args) UpdateAt(index int, value interface{}) {
	if index >= 0 && index < args.Len() {
		arg := args.items[index]
		arg.Update(value)
		args.items[index] = arg
	}
}

// IsEmpty returns true if there's no elements, false in other case
func (args *Args) IsEmpty() bool {
	return len(args.items) == 0
}

// Len return the number of arguments
func (args *Args) Len() int {
	return len(args.items)
}

// Get returns an element with the given name
func (args *Args) Get(name string) *Arg {
	for _, arg := range args.items {
		if arg.name == name {
			return arg
		}
	}
	return nil
}

// At returns the argument in the given position
func (args *Args) At(index int) *Arg {
	if len(args.items) > index && index >= 0 {
		return args.items[index]
	}
	return nil
}

// SetAt sets value for the argument in the given position
func (args *Args) SetAt(index int, value interface{}) {
	if args.Len() > index && index >= 0 {
		args.items[index].Update(value)
	}
}

// Set sets the value for the argument if it's found
func (args *Args) Set(name string, value interface{}) {
	for _, arg := range args.items {
		if arg.name == name {
			arg.Update(value)
			return
		}
	}
	args.items = append(args.items, NewArg(name, value))
}
