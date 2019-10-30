package context

import (
	"context"
	"time"
)

type contextKey string

const (
	name contextKey = "_goaFunction"
	pkg  contextKey = "_goaPkg"
	in   contextKey = "_goaIn"
	out  contextKey = "_goaOut"
)

// Pkg returns the package
func (c *Context) Pkg() string {
	if v := c.ctx.Value(pkg); v != nil {
		return v.(string)
	}
	return ""
}

// Function returns the name of the function
func (c *Context) Function() string {
	if v := c.ctx.Value(name); v != nil {
		return v.(string)
	}
	return ""
}

// In returns the input arguments
func (c *Context) In() *Args {
	if v := c.ctx.Value(in); v != nil {
		return v.(*Args)
	}
	return &Args{}
}

// Context Goa context
type Context struct {
	ctx context.Context
}

// Out returns the output arguments
func (c *Context) Out() *Args {
	if v := c.ctx.Value(out); v != nil {
		return v.(*Args)
	}
	return &Args{}
}

// NewContext constructor for goa context
func NewContext(ctx context.Context) *Context {

	return &Context{ctx}
}

// WithPkg set the package
func (c *Context) WithPkg(v string) *Context {
	c.ctx = context.WithValue(c.ctx, pkg, v)
	return c
}

// WithName set the function name
func (c *Context) WithName(v string) *Context {
	c.ctx = context.WithValue(c.ctx, name, v)
	return c
}

// WithIn set the input arguments
func (c *Context) WithIn(args []*Arg) *Context {
	c.ctx = context.WithValue(c.ctx, in, &Args{
		items: args,
	})
	return c
}

// WithOut set the output arguments
func (c *Context) WithOut(args []*Arg) *Context {
	c.ctx = context.WithValue(c.ctx, out, &Args{
		items: args,
	})
	return c
}

// Set set context value
func (c *Context) Set(key string, value interface{}) {
	c.ctx = context.WithValue(c.ctx, contextKey(key), value)
}

// GetString return the argument value
func (c *Context) GetString(key string) string {
	if value := c.Get(key); value != nil {
		return value.(string)
	}
	return ""
}

// GetInt return the argument value
func (c *Context) GetInt(key string) int {
	if value := c.Get(key); value != nil {
		return value.(int)
	}
	return 0
}

// GetBool return the argument value
func (c *Context) GetBool(key string) bool {
	if value := c.Get(key); value != nil {
		return value.(bool)
	}
	return false
}

// GetTime return the argument value
func (c *Context) GetTime(key string) time.Time {
	if value := c.Get(key); value != nil {
		return value.(time.Time)
	}
	return time.Time{}
}

// Get return the argument value
func (c *Context) Get(key string) interface{} {
	return c.ctx.Value(contextKey(key))
}

// GetIn return the argument with the given name
func (c *Context) GetIn(name string) *Arg {
	return c.In().Get(name)
}

// GetInAt return the argument in the given position
func (c *Context) GetInAt(index int) *Arg {
	return c.In().At(index)
}

// GetOut return the argument with the given name
func (c *Context) GetOut(name string) *Arg {
	return c.Out().Get(name)
}

// GetOutAt return the argument in the given position
func (c *Context) GetOutAt(index int) *Arg {
	return c.Out().At(index)
}

// SetOut set the value for the given argument
func (c *Context) SetOut(name string, value interface{}) {
	c.Out().Set(name, value)
}

// SetIn set the value for the given argument
func (c *Context) SetIn(name string, value interface{}) {
	c.In().Set(name, value)
}

// SetInAt set the value for the argument in the given position
func (c *Context) SetInAt(index int, value interface{}) {
	c.In().SetAt(index, value)
}

// SetOutAt set the value for the argument in the given position
func (c *Context) SetOutAt(index int, value interface{}) {
	c.Out().SetAt(index, value)
}
