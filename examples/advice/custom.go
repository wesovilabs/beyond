package advice

import (
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/api/context"
)

type EmptyAround struct {
}

func (c *EmptyAround) Before(ctx *context.GoaContext) {

}

func (c *EmptyAround) Returning(ctx *context.GoaContext) {

}

func NewEmptyAround() api.Around {
	return &EmptyAround{}
}

type ComplexAround struct {
	att string
}

func (c *ComplexAround) Before(ctx *context.GoaContext) {

}

func (c *ComplexAround) Returning(ctx *context.GoaContext) {

}

type Attribute struct{

}

func NewComplexAround(att string, com Attribute,_ interface{}) func() api.Around {
	return func() api.Around {
		return &ComplexAround{att}
	}
}

func NewComplexBefore(*Attribute) func() api.Before {
	return func() api.Before {
		return &ComplexAround{""}
	}
}
