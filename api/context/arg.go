package context

// Arg contains attributes to efine an argument
type Arg struct {
	name  string
	value interface{}
}

// NewArg constructs an instance of arg
func NewArg(name string, value interface{}) *Arg {
	return &Arg{
		name:  name,
		value: value,
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


// Update updates the argument value
func (a *Arg) Update(value interface{}) {
	a.value = value
}
