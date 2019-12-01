package context

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func Test_Context(t *testing.T) {
	beyondCtx := NewContext()
	beyondCtx.WithPkg("parent/child").
		WithName("function").
		SetParams(&Args{items: []*Arg{
			NewArg("firstname", "John"),
		}}).
		SetResults(&Args{[]*Arg{
			NewArg("salary", 1200.23),
			NewArg("retired", false),
		}})
	assert := assert.New(t)
	assert.Equal("parent/child", beyondCtx.Pkg())
	assert.Equal("function", beyondCtx.Function())
	assert.Equal("John", beyondCtx.Params().Get("firstname").value)
	assert.Equal(1200.23, beyondCtx.Results().Get("salary").value)
	assert.Equal(false, beyondCtx.Results().Get("retired").value)
	now := time.Now()
	beyondCtx.Set("start.time", now)
	assert.Equal(now, beyondCtx.Get("start.time"))
	beyondCtx.Set("married", true)
	assert.Equal(true, beyondCtx.Get("married"))

	beyondCtx.Set("age", 34)
	assert.Equal(34, beyondCtx.Get("age"))

	beyondCtx.Set("firstname", "Wenceslao")
	assert.Equal("Wenceslao", beyondCtx.Get("firstname"))

	beyondCtx = NewContext()
	assert.Empty(beyondCtx.Function())
	assert.Empty(beyondCtx.Pkg())
	assert.Empty(beyondCtx.Params().items)
	assert.Empty(beyondCtx.Results().items)
}

type Data struct {
}

func TestContext_ParamsGet(t *testing.T) {
	beyondCtx := NewContext()
	d := Data{}
	beyondCtx.WithPkg("parent/child").
		WithName("function").
		WithType(d).
		SetParams(&Args{[]*Arg{
			NewArg("firstname", "John"),
		}}).
		SetResults(&Args{[]*Arg{
			NewArg("salary", 1200.23),
			NewArg("retired", false),
		}})
	assert := assert.New(t)
	assert.Equal("John", beyondCtx.Params().Get("firstname").value)
	assert.Nil(nil, beyondCtx.Params().Get("unknown"))

	assert.Equal("John", beyondCtx.Params().At(0).value)
	assert.Nil(nil, beyondCtx.Params().At(20))

	assert.Equal(1200.23, beyondCtx.Results().At(0).value)
	assert.Nil(nil, beyondCtx.Results().At(10))

	assert.Equal(d, beyondCtx.Type())

	beyondCtx.Params().Set("name", "tom")
	assert.Equal("tom", beyondCtx.Params().Get("name").value)
	beyondCtx.Params().Set("name", "Tim")
	assert.Equal("Tim", beyondCtx.Params().Get("name").value)

	beyondCtx.Results().Set("name", "tom")
	assert.Equal("tom", beyondCtx.Results().Get("name").value)
	beyondCtx.Results().Set("name", "Tim")
	assert.Equal("Tim", beyondCtx.Results().Get("name").value)

	beyondCtx.Params().SetAt(0, "tom")
	assert.Equal("tom", beyondCtx.Params().Get("firstname").value)
	beyondCtx.Params().SetAt(20, "Tim")

	beyondCtx.Results().SetAt(0, "tom")
	assert.Equal("tom", beyondCtx.Results().Get("salary").value)
	assert.Equal(reflect.TypeOf("tom").String(), beyondCtx.Results().Get("salary").kind)

}

func Test_Exit(t *testing.T) {
	assert := assert.New(t)
	beyondCtx := NewContext()
	assert.False(beyondCtx.IsCompleted())
	beyondCtx.Exit()
	assert.True(beyondCtx.IsCompleted())
	assert.Nil(beyondCtx.Type())
}
