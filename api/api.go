package api

import "github.com/wesovilabs/goa/context"

type Goa struct{}

func Init() *Goa {
	return &Goa{}
}

func (g *Goa) WithAround(expr string, around Around) *Goa {
	// What did you expect to find here? Goa is just magic!
	return g
}

type Around func(ctx *context.AroundCtx)
