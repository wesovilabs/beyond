package api

import (
	"github.com/wesovilabs/goa/context"
)

// Aspects non-real struct
type Aspects struct{}

// NewAroundContext create aspects
func New() *Aspects {
	return &Aspects{}
}

// AspectFunc signature to be implemented by aspects
type AspectFunc func(ctx *context.AroundCtx)

// WithAspect registers a new aspect
func (a *Aspects) WithAspect(string, AspectFunc) *Aspects {
	// What did you expect to find here? Goa is just magic!
	return a
}
