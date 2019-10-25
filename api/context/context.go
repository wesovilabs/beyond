package context

import (
	"context"
	"time"
)

type contextKey int32

const (
	name contextKey = iota
	pkg
	in
	out
)

func (c *Context) Pkg() string {
	return c.ctx.Value(pkg).(string)
}
func (c *Context) Function() string {
	return c.ctx.Value(name).(string)
}

func (c *Context) In() *Args {
	return c.ctx.Value(in).(*Args)
}

type Context struct {
	ctx context.Context
}

func (c *Context) Out() *Args {
	return c.ctx.Value(out).(*Args)
}

// NewContext constructor for goa context
func NewContext(ctx context.Context, p ...func(*builder)) *Context {
	b := defaultBuilder
	for _, fn := range p {
		fn(&b)
	}
	ctx = context.WithValue(ctx, name, b.funcName)
	ctx = context.WithValue(ctx, pkg, b.pkgName)
	ctx = context.WithValue(ctx, in, b.input)
	ctx = context.WithValue(ctx, out, b.output)
	return &Context{ctx}
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
func WithIn(args *Args) Builder {
	return func(b *builder) {
		b.input = args
	}
}

// WithOutput sets the output
func WithOut(args *Args) Builder {
	return func(b *builder) {
		b.output = args
	}
}

func (c *Context) Set(key string, value interface{}) {
	c.ctx = context.WithValue(c.ctx, key, value)
}

func (c *Context) GetString(key string) string {
	if value := c.Get(key); value != nil {
		return value.(string)
	}
	return ""
}

func (c *Context) GetInt(key string) int {
	if value := c.Get(key); value != nil {
		return value.(int)
	}
	return 0
}

func (c *Context) GetBool(key string) bool {
	if value := c.Get(key); value != nil {
		return value.(bool)
	}
	return false
}

func (c *Context) GetTime(key string) time.Time {
	if value := c.Get(key); value != nil {
		return value.(time.Time)
	}
	return time.Time{}
}

func (c *Context) Get(key string) interface{} {
	return c.ctx.Value(key)
}

type builder struct {
	pkgName  string
	funcName string
	input    *Args
	output   *Args
}

var defaultBuilder = builder{
	input:    &Args{},
	output:   &Args{},
}

// Builder builder type
type Builder func(*builder)
