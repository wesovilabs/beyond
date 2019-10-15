package context

import (
	"context"
)

type Key int32

const (
	funcName Key = iota
	pkgName
	funcIn
	funcOut
)

func New(ctx context.Context, p ...func(*builder)) (*Ctx, error) {
	builder := defaultBuilder
	for _, fn := range p {
		fn(&builder)
	}
	ctx = context.WithValue(ctx, funcName, builder.funcName)
	ctx = context.WithValue(ctx, pkgName, builder.pkgName)
	ctx = context.WithValue(ctx, funcIn, builder.input)
	ctx = context.WithValue(ctx, funcOut, builder.output)
	return &Ctx{ctx}, nil
}

type Ctx struct {
	ctx context.Context
}

func (c *Ctx) In() Input {
	if input := c.ctx.Value(funcIn); input != nil {
		return input.(Input)
	}
	return nil
}

func (c *Ctx) Out() *Output {
	if output := c.ctx.Value(funcOut); output != nil {
		return output.(*Output)
	}
	return nil
}

func (c *Ctx) Name() string {
	if name := c.ctx.Value(funcName); name != nil {
		return name.(string)
	}
	return ""
}

func (c *Ctx) Pkg() string {
	if pkg := c.ctx.Value(pkgName); pkg != nil {
		return pkg.(string)
	}
	return ""
}

type builder struct {
	pkgName  string
	funcName string
	input    Input
	output   Output
}

var defaultBuilder = builder{
	pkgName: "",
	input:   Input{},
	output:  Output{},
}

// Builder builder type
type Builder func(*builder)

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
