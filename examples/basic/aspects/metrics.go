package aspects

import (
	"fmt"
	"github.com/wesovilabs/goa/context"
)

var counter = map[string]int{}

func CalculateMetrics(ctx *context.Ctx) {
	if _, ok := counter[ctx.Name()]; ok {
		counter[ctx.Name()]++
		return
	}
	counter[ctx.Name()] = 1

}

func PrintCounter() {
	for k, v := range counter {
		fmt.Printf("%s: %v\n", k, v)
	}
}
