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
		WithIn([]*Arg{
			NewArg("firstname", "John"),
		}).
		WithOut([]*Arg{
			NewArg("salary", 1200.23),
			NewArg("retired", false),
		})
	assert := assert.New(t)
	assert.Equal("parent/child", goaCtx.Pkg())
	assert.Equal("function", goaCtx.Function())
	assert.Equal("John", goaCtx.In().Get("firstname").value)
	assert.Equal(1200.23, goaCtx.Out().Get("salary").value)
	assert.Equal(false, goaCtx.Out().Get("retired").value)
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
	assert.Empty(goaCtx.In().items)
	assert.Empty(goaCtx.Out().items)
}

func TestContext_GetIn(t *testing.T) {
	ctx := context.Background()
	goaCtx := NewContext(ctx)
	goaCtx.WithPkg("parent/child").
		WithName("function").
		WithIn([]*Arg{
			NewArg("firstname", "John"),
		}).
		WithOut([]*Arg{
			NewArg("salary", 1200.23),
			NewArg("retired", false),
		})
	assert := assert.New(t)
	assert.Equal("John", goaCtx.GetIn("firstname").value)
	assert.Nil(nil, goaCtx.GetIn("unknown"))

	assert.Equal("John", goaCtx.GetInAt(0).value)
	assert.Nil(nil, goaCtx.GetInAt(20))

	assert.Equal(1200.23, goaCtx.GetOutAt(0).value)
	assert.Nil(nil, goaCtx.GetOutAt(10))

	goaCtx.SetIn("name", "tom")
	assert.Equal("tom", goaCtx.GetIn("name").value)
	goaCtx.SetIn("name", "Tim")
	assert.Equal("Tim", goaCtx.GetIn("name").value)

	goaCtx.SetOut("name", "tom")
	assert.Equal("tom", goaCtx.GetOut("name").value)
	goaCtx.SetOut("name", "Tim")
	assert.Equal("Tim", goaCtx.GetOut("name").value)

	goaCtx.SetInAt(0, "tom")
	assert.Equal("tom", goaCtx.GetIn("firstname").value)
	goaCtx.SetInAt(20, "Tim")

	goaCtx.SetOutAt(0, "tom")
	assert.Equal("tom", goaCtx.GetOut("salary").value)
	assert.Equal(reflect.TypeOf("tom"), goaCtx.GetOut("salary").kind)

}
