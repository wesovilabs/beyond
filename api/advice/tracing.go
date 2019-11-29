package advice

import (
	"fmt"
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond/api/context"
	"strings"
)

// TracingAdvice trace your function invocations
type TracingAdvice struct{}

// Before required by Before interface
func (c *TracingAdvice) Before(ctx *context.BeyondContext) {
	params := make([]string, ctx.Params().Count())
	ctx.Params().ForEach(func(index int, arg *context.Arg) {
		params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
	})
	fmt.Printf("[advice.tracing] %s.%s(%s)\n", ctx.Pkg(), ctx.Function(), strings.Join(params, ","))
}

// NewTracingAdvice return an instance of TracingAdvice
func NewTracingAdvice() api.Before {
	return &TracingAdvice{}
}
