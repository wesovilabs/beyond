package context

// Input struct
type Input []*Arg

// List returns the list of arguments
func (i Input) List() []*Arg {
	return []*Arg(i)
}

func (i Input) isEmpty() bool {
	return i.Len() == 0
}

// Len return the number of arguments
func (i Input) Len() int {
	return len(i)
}

// Get returns an element with the given name
func (i Input) Get(name string) interface{} {
	for _, arg := range i.List() {
		if arg.name == name {
			return arg.Value()
		}
	}
	return nil
}
