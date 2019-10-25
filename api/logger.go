package api

import (
	"fmt"
	goaContext "github.com/wesovilabs/goa/api/context"
)
import "context"

type Logger struct {
	ctx *context.Context
}

func (l *Logger) Before(ctx *goaContext.Context) {
	fmt.Println("EOE")
}

func (l *Logger) Returning(ctx *goaContext.Context) {

}
