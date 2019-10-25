package api

type Goa struct{}

func Init() *Goa {
	return &Goa{}
}

func (g *Goa) WithAround(string, func() Around) *Goa {
	return g
}

type BeforeFn func() Before

func (g *Goa) WithBefore(string, BeforeFn) *Goa {
	return g
}

func (g *Goa) WithReturning(string, func() Returning) *Goa {
	return g
}
