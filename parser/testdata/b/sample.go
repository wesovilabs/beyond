package b

import (
	"github.com/wesovilabs/goa/parser/testdata/b/b1"
	"github.com/wesovilabs/goa/parser/testdata/c"
)

type Address struct {
	country *c.Country
	demo *b1.Demo
}
