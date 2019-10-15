package context

type Input []*Arg

func (i Input) List() []*Arg {
	return []*Arg(i)
}

func (i Input) isEmpty() bool {
	return i.Len() == 0
}

func (i Input) Len() int {
	return len(i)
}

func (i Input) Get(name string) interface{} {
	for _, arg := range i.List() {
		if arg.name == name {
			return arg.Value()
		}
	}
	return nil
}
