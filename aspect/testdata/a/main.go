package a

import (
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/api/context"
	"github.com/wesovilabs/goa/aspect/testdata/a/a1"
	"github.com/wesovilabs/goa/aspect/testdata/a/a1/a11"
	b2 "github.com/wesovilabs/goa/aspect/testdata/a/a2"
)

type Tracing struct {
	value  string
	value2 int
}

func (t *Tracing) Before(ctx *context.BeforeCtx) {

}

func (t *Tracing) Returning(ctx *context.Context) {

}

func NewTracingBefore() api.Before {
	return &Tracing{}
}

func NewTracingReturning() api.Returning {
	return &Tracing{}
}

func NewTracingAround() *Tracing {
	return &Tracing{}
}

func Around() api.Around {
	return nil
}

func Goa() *api.Goa {
	return api.Init().
		WithAround("*.*(*)...", a1.AroundA1).
		WithAround("*.*(*)...", Around).
		WithBefore("*.*(*)...", NewTracingBefore).
		WithReturning("*.*(*)...", NewTracingReturning).
		WithAround("*.*(*)...", a11.AroundA11).
		WithAround("*.*(*)...", b2.AroundA2)
}
