package api

import "github.com/wesovilabs/goa/api/context"

// Advice definition
type Advice interface {
}

// Before definition
type Before interface {
	Before(ctx *context.GoaContext)
}

// Returning definition
type Returning interface {
	Returning(ctx *context.GoaContext)
}

// Around definition
type Around interface {
	Before
	Returning
}
