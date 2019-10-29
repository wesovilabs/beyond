package aspects

import (
        "fmt"
        "github.com/wesovilabs/goa/api"
        "github.com/wesovilabs/goa/api/context"
)

type TracingAspect struct{}

func (t *TracingAspect) Before(ctx *context.Context) {
        text := fmt.Sprintf("[%s] => ", ctx.Function())
        for _, arg := range ctx.In().List() {
                text = fmt.Sprintf("%s | %s:%#v", text, arg.Name(), arg.Value())
        }
        fmt.Println(text)
}
func NewTracingAspect() api.Before {
        return &TracingAspect{}
}
