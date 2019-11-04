package context

// Args struct
type Args struct {
	items []*Arg
}

// IsEmpty returns true if there's no elements, false in other case
func (args *Args) isEmpty() bool {
	return len(args.items) == 0
}

// Len return the number of arguments
func (args *Args) len() int {
	return len(args.items)
}

func (args *Args) get(name string) *Arg {
	for _, arg := range args.items {
		if arg.name == name {
			return arg
		}
	}

	return nil
}

// At returns the argument in the given position
func (args *Args) at(index int) *Arg {
	if len(args.items) > index && index >= 0 {
		return args.items[index]
	}

	return nil
}
