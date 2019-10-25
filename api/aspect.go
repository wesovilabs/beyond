package api

import "github.com/wesovilabs/goa/api/context"

type Aspect interface {

}

type Before interface {
	Before(ctx *context.Context)
}

type Returning interface {
	Returning(ctx *context.Context)
}

type Around interface {
	Before
	Returning
}
