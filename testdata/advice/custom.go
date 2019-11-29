package advice

import (
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond/api/context"
)

type EmptyAround struct {
}

func (c *EmptyAround) Before(ctx *context.BeyondContext) {

}

func (c *EmptyAround) Returning(ctx *context.BeyondContext) {

}

func NewEmptyAround() api.Around {
	return &EmptyAround{}
}

type ComplexAround struct {
	att string
}

func (c *ComplexAround) Before(ctx *context.BeyondContext) {

}

func (c *ComplexAround) Returning(ctx *context.BeyondContext) {

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
