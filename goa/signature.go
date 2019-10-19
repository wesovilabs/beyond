package goa

import (
	"github.com/wesovilabs/goa/context"
)

// Aspects non-real struct
type Aspects struct{}

// New create aspects
func New() *Aspects {
	return &Aspects{}
}

// AspectFunc signature to be implemented by aspects
type AspectFunc func(ctx *context.Ctx)

// WithAspect registers a new aspect
func (a *Aspects) WithAspect(string, AspectFunc) *Aspects {
	// What did you expect to find here? Goa is just magic!
	return a
}
