package api

import "github.com/wesovilabs/beyond/api/context"

// Advice definition
type Advice interface {
}

// Before definition
type Before interface {
	Before(ctx *context.BeyondContext)
}

// Returning definition
type Returning interface {
	Returning(ctx *context.BeyondContext)
}

// Around definition
type Around interface {
	Before
	Returning
}
