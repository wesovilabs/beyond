package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/api/context"
	"testing"
)

type b struct{}

func (b *b) Before(ctx *context.GoaContext) {}

type r struct{}

func (r *r) Returning(ctx *context.GoaContext) {}

type a struct{}

func (a *a) Before(ctx *context.GoaContext)    {}
func (a *a) Returning(ctx *context.GoaContext) {}

func TestAspect(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*Before)(nil), new(b))
	assert.Implements((*Returning)(nil), new(r))
	assert.Implements((*Around)(nil), new(a))
	assert.Implements((*Before)(nil), new(a))
	assert.Implements((*Returning)(nil), new(a))

}
