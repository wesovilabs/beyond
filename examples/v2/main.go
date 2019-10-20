package main

import (
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/examples/v2/aspect"
)

func Goa() *api.Goa {
	loggerAspect := &aspect.LoggerAround{
		Prefix: "[goa]",
	}
	return Goa().
		WithAround(loggerAspect)
}
