package advice

import (
	"fmt"
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/api/context"
	"strings"
)

type TracingAdvice struct{}

func (c *TracingAdvice) Before(ctx *context.GoaContext) {
	params := make([]string, ctx.Params().Count())
	ctx.Params().ForEach(func(index int, arg *context.Arg) {
		params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
	})
	fmt.Printf("[advice.tracing] %s.%s(%s)\n", ctx.Pkg(), ctx.Function(), strings.Join(params, ","))
}

func NewTracingAdvice() api.Before {
	return &TracingAdvice{}
}
