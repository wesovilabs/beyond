package time

import (
	"fmt"
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/api/context"
	"time"
)

type TimingMode int32

const (
	nanoseconds TimingMode = iota
	microseconds
)

type TimingAspect struct {
	mode TimingMode
}

func (t *TimingAspect) Before(ctx *context.GoaContext) {
	ctx.Set("start", time.Now())
}

func (t *TimingAspect) Returning(ctx *context.GoaContext) {
	start := ctx.Get("start").(time.Time)
	end := time.Now()
	duration := end.Sub(start)
	switch t.mode {
	case nanoseconds:
		fmt.Printf("%v nanoseconds\n",duration.Nanoseconds())
	case microseconds:
		fmt.Printf("%v microseconds\n",duration.Microseconds())
	}

}

func TimingNanoSeconds() api.Around {
	return &TimingAspect{
		mode: nanoseconds,
	}
}

func TimingMicroSeconds() api.Around {
	return &TimingAspect{
		mode: microseconds,
	}
}
