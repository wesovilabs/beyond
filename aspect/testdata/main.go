package main

import (
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/aspect/testdata/aspect"
	"github.com/wesovilabs/goa/aspect/testdata/other"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

func Goa() *api.Goa {
	return Goa().
		WithAround("test.demo(string)", aspect.LoggerAround).
		WithAround("test.demo(string)", aspect.LoggerAround)
}

func main(){
	person:=&other.Person{}
	fmt.Printf(person)
}
