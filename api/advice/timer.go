package advice

import (
	"fmt"
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/api/context"
	"strings"
	"time"
)

const timeStartKey = "time.start"

type TimerMode int32

const (
	Nanoseconds TimerMode = iota
	Microseconds
)

type TimerAdvice struct {
	mode TimerMode
}

func (a *TimerAdvice) Before(ctx *context.GoaContext) {
	ctx.Set(timeStartKey, time.Now())
}

func (a *TimerAdvice) Returning(ctx *context.GoaContext) {
	start := ctx.Get(timeStartKey).(time.Time)
	timeDuration:="?"
	switch a.mode {
	case Nanoseconds:
		timeDuration = fmt.Sprintf("%v nanoseconds\n", time.Since(start).Nanoseconds())
	case Microseconds:
		timeDuration = fmt.Sprintf("%v microseconds\n", time.Since(start).Microseconds())
	}
	params := make([]string, ctx.Params().Count())
	ctx.Params().ForEach(func(index int, arg *context.Arg) {
		params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
	})
	fmt.Printf("[goa.timer] %s.%s(%s) took %s", ctx.Pkg(), ctx.Function(), strings.Join(params, ","),timeDuration)
}

func NewTimerAdvice(mode TimerMode) func() api.Around {
	return func() api.Around{
		return &TimerAdvice{mode}
	}
}
