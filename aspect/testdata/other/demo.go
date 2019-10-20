package other

import (
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/aspect/testdata/aspect"
)

type Person struct {

}

func Goa() *api.Goa {
	return Goa().
		WithAround("test.other(string)", aspect.LoggerAround).
		WithAround("s.other(string)", aspect.LoggerAround)
}
