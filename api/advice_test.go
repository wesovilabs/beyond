package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/beyond/api/context"
	"testing"
)

type b struct{}

func (b *b) Before(ctx *context.BeyondContext) {}

type r struct{}

func (r *r) Returning(ctx *context.BeyondContext) {}

type a struct{}

func (a *a) Before(ctx *context.BeyondContext)    {}
func (a *a) Returning(ctx *context.BeyondContext) {}

func TestAspect(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*Before)(nil), new(b))
	assert.Implements((*Returning)(nil), new(r))
	assert.Implements((*Around)(nil), new(a))
	assert.Implements((*Before)(nil), new(a))
	assert.Implements((*Returning)(nil), new(a))

}
