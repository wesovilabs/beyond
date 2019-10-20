package context

import (
	"context"
)

type contextKey int32

const (
	funcName contextKey = iota
	pkgName
	funcIn
	funcOut
)

// NewAroundContext constructor for goa context
func NewAroundContext(ctx context.Context, p ...func(*builder)) (*AroundCtx, error) {
	b := defaultBuilder
	for _, fn := range p {
		fn(&b)
	}
	ctx = context.WithValue(ctx, funcName, b.funcName)
	ctx = context.WithValue(ctx, pkgName, b.pkgName)
	ctx = context.WithValue(ctx, funcIn, b.input)
	ctx = context.WithValue(ctx, funcOut, b.output)
	return &AroundCtx{ctx}, nil
}

// AroundCtx context for goa
type AroundCtx struct {
	ctx context.Context
}

// In returns the Input arguments
func (c *AroundCtx) In() Input {
	if input := c.ctx.Value(funcIn); input != nil {
		return input.(Input)
	}
	return nil
}

// Out returns the Output arguments
func (c *AroundCtx) Out() *Output {
	if output := c.ctx.Value(funcOut); output != nil {
		return output.(*Output)
	}
	return nil
}

// Name returns the function name
func (c *AroundCtx) Name() string {
	if name := c.ctx.Value(funcName); name != nil {
		return name.(string)
	}
	return ""
}

// Pkg returns the package name
func (c *AroundCtx) Pkg() string {
	if pkg := c.ctx.Value(pkgName); pkg != nil {
		return pkg.(string)
	}
	return ""
}

// WithPkg sets the package name
func WithPkg(pkgName string) Builder {
	return func(b *builder) {
		b.pkgName = pkgName
	}
}

// WithName sets the name
func WithName(name string) Builder {
	return func(b *builder) {
		b.funcName = name
	}
}

// WithInput sets the input
func WithInput(input Input) Builder {
	return func(b *builder) {
		b.input = input
	}
}

// WithOutput sets the output
func WithOutput(output Output) Builder {
	return func(b *builder) {
		b.output = output
	}
}
