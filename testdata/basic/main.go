package main

import (
	"fmt"
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/api/context"
)

type Tracing struct {
}

func (t *Tracing) Before(ctx *context.Context) {

}

func (t *Tracing) Returning(ctx *context.Context) {

}

func NewTracing() api.Around {
	return &Tracing{}
}

var beforeFn = func()api.Before{
	return nil
}

func Goa() *api.Goa {
	return api.Init().
		WithBefore("*.*(*)", beforeFn).
		WithAround("*.*(*)...", NewTracing)
}

func main() {
	sayHello("john")
}

func sayHello(name string) {
	fmt.Printf("Hi %s\n", name)
}
