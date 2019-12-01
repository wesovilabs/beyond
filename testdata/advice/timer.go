package advice

import (
	"fmt"
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond/api/context"
	"strings"
	"time"
)

const timeStartKey = "time.start"

// TimerMode supported timer modes
type TimerMode int32

const (
	// Nanoseconds timerMode
	Nanoseconds TimerMode = iota
	// Microseconds timerMode
	Microseconds
)

// TimerAdvice advice definition
type TimerAdvice struct {
	mode TimerMode
}

// Before required by Around interface
func (a *TimerAdvice) Before(ctx *context.BeyondContext) {
	ctx.Set(timeStartKey, time.Now())
}

// Returning required by Around interface
func (a *TimerAdvice) Returning(ctx *context.BeyondContext) {
	start := ctx.Get(timeStartKey).(time.Time)
	timeDuration := "?"

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
	fmt.Printf("[beyond.timer] %s.%s(%s) took %s", ctx.Pkg(), ctx.Function(), strings.Join(params, ","), timeDuration)
}

// NewTimerAdvice returns an instance of Timer advice
func NewTimerAdvice(mode TimerMode) func() api.Around {
	return func() api.Around {
		return &TimerAdvice{mode}
	}
}
