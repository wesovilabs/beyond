package funcs

import (
	"github.com/wesovilabs/goa/context"
	"github.com/wesovilabs/goa/goa"
	"strings"
)

func Goa() *goa.Aspects {
	return goa.New().
		WithAspect("*.*(func(string)string)*", ToUpper).
		WithAspect("*.*(string)*", ToUpper)

}

func main() {

}

func ToUpper(ctx *context.Ctx) {
	ctx.In().List()[0].Update(strings.ToUpper(ctx.In().List()[0].Value().(string)))
}
