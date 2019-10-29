package testdata

import (
	"github.com/wesovilabs/goa/api"
)

var beforeFn = func() api.Before {
	return nil
}

func Goa() *api.Goa {
	return api.Init().
		WithBefore("*.sayHello(...)...", beforeFn).
		WithBefore("*.sayBye(...)...", beforeFn).
		WithBefore("*.say*(...)...", beforeFn)
}

func sayHello(){

}

func sayBye(){

}
