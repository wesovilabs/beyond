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
	if v := c.ctx.Value(pkg); v != nil {
		return v.(string)
	}
	return ""
}
func (c *Context) Function() string {
	if v := c.ctx.Value(name); v != nil {
		return v.(string)
	}
	return ""
}

func (c *Context) In() *Args {
	if v := c.ctx.Value(in); v != nil {
		return v.(*Args)
	}
	return &Args{}
}

type Context struct {
	ctx context.Context
}

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
func (c *Context) WithPkg(v string) *Context {
	c.ctx = context.WithValue(c.ctx, pkg, v)
	return c
}

func (c *Context) WithName(v string) *Context {
	c.ctx = context.WithValue(c.ctx, name, v)
	return c
}

func (c *Context) WithIn(args []*Arg) *Context {
	c.ctx = context.WithValue(c.ctx, in, &Args{
		items: args,
	})
	return c
}

func (c *Context) WithOut(args []*Arg) *Context {
	c.ctx = context.WithValue(c.ctx, out, &Args{
		items: args,
	})
	return c
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

func (c *Context) GetIn(name string) *Arg {
	return c.In().Get(name)
}

func (c *Context) GetInAt(index int) *Arg {
	return c.In().At(index)
}

func (c *Context) GetOut(name string) *Arg {
	return c.Out().Get(name)
}

func (c *Context) GetOutAt(index int) *Arg {
	return c.Out().At(index)
}

func (c *Context) SetOut(name string, value interface{}) {
	c.Out().Set(name, value)
}

func (c *Context) SetIn(name string, value interface{}) {
	c.In().Set(name, value)
}

func (c *Context) SetInAt(index int, value interface{}) {
	c.In().SetAt(index, value)
}

func (c *Context) SetOutAt(index int, value interface{}) {
	c.Out().SetAt(index, value)
}
