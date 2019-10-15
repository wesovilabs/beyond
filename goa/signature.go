package goa

import (
	"github.com/wesovilabs/goa/context"
)

type Aspects struct{}

func New() *Aspects {
	return &Aspects{}
}

type AspectFunc func(ctx *context.Ctx)

func (a *Aspects) WithAspect(string, AspectFunc) *Aspects {
	// What did you expect to find here? This is just magic
	return a
}
