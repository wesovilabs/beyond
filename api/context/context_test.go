package context

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Context(t *testing.T) {
	ctx := context.Background()
	goaCtx := NewContext(ctx,
		WithPkg("parent/child"),
		WithName("function"),
		WithIn(&Args{
			NewArg("firstname", "John"),
		}),
		WithOut(&Args{
			NewArg("salary", 1200.23),
			NewArg("retired", false),
		}))
	assert := assert.New(t)
	assert.Equal("parent/child", goaCtx.Pkg())
	assert.Equal("function", goaCtx.Function())
	assert.Equal("John", goaCtx.In().Get("firstname"))
	assert.Equal(1200.23, goaCtx.Out().Get("salary"))
	assert.Equal(false, goaCtx.Out().Get("retired"))
	now := time.Now()
	goaCtx.Set("start.time", now)
	assert.Equal(now, goaCtx.Get("start.time"))
	assert.Equal(now, goaCtx.GetTime("start.time"))
	goaCtx.Set("married", true)
	assert.Equal(true, goaCtx.Get("married"))
	assert.Equal(true, goaCtx.GetBool("married"))

	goaCtx.Set("age", 34)
	assert.Equal(34, goaCtx.Get("age"))
	assert.Equal(34, goaCtx.GetInt("age"))

	goaCtx.Set("firstname", "Wenceslao")
	assert.Equal("Wenceslao", goaCtx.Get("firstname"))
	assert.Equal("Wenceslao", goaCtx.GetString("firstname"))
	assert.Equal("", goaCtx.GetString("unknown"))
	assert.Equal(0, goaCtx.GetInt("unknown"))
	assert.Equal(false, goaCtx.GetBool("unknown"))
	assert.Equal(time.Time{}, goaCtx.GetTime("unknown"))

	goaCtx = NewContext(context.Background())
	assert.Empty(goaCtx.Function())
	assert.Empty(goaCtx.Pkg())
	assert.Equal(&Args{},goaCtx.In())
	assert.Equal(&Args{},goaCtx.Out())
}
