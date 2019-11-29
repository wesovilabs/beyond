package api

// Beyond struct used to register the aspects
type Beyond struct{}

// New initialize the Beyond type
func New() *Beyond {
	return &Beyond{}
}

// WithAround registers around aspects
func (g *Beyond) WithAround(func() Around, string) *Beyond {
	return g
}

// WithBefore registers before aspects
func (g *Beyond) WithBefore(func() Before, string) *Beyond {
	return g
}

// WithReturning registers returning aspects
func (g *Beyond) WithReturning(func() Returning, string) *Beyond {
	return g
}

// Exclude add path to be ignored by beyond advices
func (g *Beyond) Exclude(...string) *Beyond {
	return g
}
