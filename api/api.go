package api

// Goa struct used to register the aspects
type Goa struct{}

// New initialize the Goa type
func New() *Goa {
	return &Goa{}
}

// WithAround registers around aspects
func (g *Goa) WithAround(func() Around, string) *Goa {
	return g
}

// WithBefore registers before aspects
func (g *Goa) WithBefore(func() Before, string) *Goa {
	return g
}

// WithReturning registers returning aspects
func (g *Goa) WithReturning(func() Returning, string) *Goa {
	return g
}
