package api

type Goa struct{}

func Init() *Goa {
	return &Goa{}
}

func (g *Goa) WithAround(around Around) *Goa {
	// What did you expect to find here? Goa is just magic!
	return g
}
