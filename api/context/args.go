package context

// Args struct
type Args []*Arg

// List returns the list of arguments
func (args Args) List() []*Arg {
	return []*Arg(args)
}

func (args Args) UpdateAt(index int, value interface{}) {
	if index >= 0 && index < len(args.List()) {
		args.List()[index].Update(value)
	}
}

func (args Args) IsEmpty() bool {
	return args.Len() == 0
}

// Len return the number of arguments
func (args Args) Len() int {
	return len(args)
}

// Get returns an element with the given name
func (args Args) Get(name string) interface{} {
	for _, arg := range args.List() {
		if arg.name == name {
			return arg.Value()
		}
	}
	return nil
}

// Set sets the value for the argument if it's found
func (args Args) Set(name string, value interface{}) {
	for _, arg := range args.List() {
		if arg.name == name {
			arg.value = value
			return
		}
	}
}
