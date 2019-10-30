package api

import "github.com/wesovilabs/goa/api/context"

// Aspect definition
type Aspect interface {
}

// Before definition
type Before interface {
	Before(ctx *context.Context)
}

// Returning definition
type Returning interface {
	Returning(ctx *context.Context)
}

// Around definition
type Around interface {
	Before
	Returning
}
