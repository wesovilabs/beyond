package testdata

import (
	"github.com/wesovilabs/goa/api"
)

var beforeFn = func() api.Before {
	return nil
}

func Goa() *api.Goa {
	return api.New().
		WithBefore(beforeFn, "sayHello(...)...").
		WithBefore(beforeFn, "sayBye(...)...").
		WithBefore(beforeFn, "say*(...)...")
}

func sayHello() {

}

func sayBye() {

}
