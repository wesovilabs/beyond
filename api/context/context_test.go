package context

import (
	"context"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func Test_Context(t *testing.T) {
	ctx := context.Background()
	goaCtx := NewContext(ctx)
	goaCtx.WithPkg("parent/child").
		WithName("function").
		SetParams(&Args{items: []*Arg{
			NewArg("firstname", "John"),
		}}).
		SetResults(&Args{[]*Arg{
			NewArg("salary", 1200.23),
			NewArg("retired", false),
		}})
	assert := assert.New(t)
	assert.Equal("parent/child", goaCtx.Pkg())
	assert.Equal("function", goaCtx.Function())
	assert.Equal("John", goaCtx.Params().Get("firstname").value)
	assert.Equal(1200.23, goaCtx.Results().Get("salary").value)
	assert.Equal(false, goaCtx.Results().Get("retired").value)
	now := time.Now()
	goaCtx.Set("start.time", now)
	assert.Equal(now, goaCtx.Get("start.time"))
	goaCtx.Set("married", true)
	assert.Equal(true, goaCtx.Get("married"))

	goaCtx.Set("age", 34)
	assert.Equal(34, goaCtx.Get("age"))

	goaCtx.Set("firstname", "Wenceslao")
	assert.Equal("Wenceslao", goaCtx.Get("firstname"))

	goaCtx = NewContext(context.Background())
	assert.Empty(goaCtx.Function())
	assert.Empty(goaCtx.Pkg())
	assert.Empty(goaCtx.Params().items)
	assert.Empty(goaCtx.Results().items)
}

func TestContext_ParamsGet(t *testing.T) {
	ctx := context.Background()
	goaCtx := NewContext(ctx)
	goaCtx.WithPkg("parent/child").
		WithName("function").
		SetParams(&Args{[]*Arg{
			NewArg("firstname", "John"),
		}}).
		SetResults(&Args{[]*Arg{
			NewArg("salary", 1200.23),
			NewArg("retired", false),
		}})
	assert := assert.New(t)
	assert.Equal("John", goaCtx.Params().Get("firstname").value)
	assert.Nil(nil, goaCtx.Params().Get("unknown"))

	assert.Equal("John", goaCtx.Params().At(0).value)
	assert.Nil(nil, goaCtx.Params().At(20))

	assert.Equal(1200.23, goaCtx.Results().At(0).value)
	assert.Nil(nil, goaCtx.Results().At(10))

	goaCtx.Params().Set("name", "tom")
	assert.Equal("tom", goaCtx.Params().Get("name").value)
	goaCtx.Params().Set("name", "Tim")
	assert.Equal("Tim", goaCtx.Params().Get("name").value)

	goaCtx.Results().Set("name", "tom")
	assert.Equal("tom", goaCtx.Results().Get("name").value)
	goaCtx.Results().Set("name", "Tim")
	assert.Equal("Tim", goaCtx.Results().Get("name").value)

	goaCtx.Params().SetAt(0, "tom")
	assert.Equal("tom", goaCtx.Params().Get("firstname").value)
	goaCtx.Params().SetAt(20, "Tim")

	goaCtx.Results().SetAt(0, "tom")
	assert.Equal("tom", goaCtx.Results().Get("salary").value)
	assert.Equal(reflect.TypeOf("tom"), goaCtx.Results().Get("salary").kind)

}
