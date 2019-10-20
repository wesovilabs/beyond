package funcs

import (
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/context"
	"strings"
)

func Goa() *api.Aspects {
	return api.New().
		WithAspect("*.*(func(string)string)*", ToUpper).
		WithAspect("*.*(string)*", ToUpper)

}

func main() {

}

func ToUpper(ctx *context.Ctx) {
	ctx.In().List()[0].Update(strings.ToUpper(ctx.In().List()[0].Value().(string)))
}
