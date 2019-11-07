package context

import (
	"context"
	"time"
)

type contextKey string

const (
	funcType contextKey = "_goaFunctionType"
	name     contextKey = "_goaFunction"
	pkg      contextKey = "_goaPkg"
	in       contextKey = "_goaIn"
	out      contextKey = "_goaOut"
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

// Type returns the type
func (c *Context) Type() interface{} {
	if v := c.ctx.Value(funcType); v != nil {
		return v
	}

	return nil
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

// WithType set the function type
func (c *Context) WithType(v interface{}) *Context {
	c.ctx = context.WithValue(c.ctx, funcType, v)
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
	return c.In().get(name)
}

// GetInValue return the argument with the given name
func (c *Context) GetInValue(name string) interface{} {
	return c.In().get(name).value
}

// GetInAt return the argument in the given position
func (c *Context) GetInAt(index int) *Arg {
	return c.In().at(index)
}

// GetInValueAt return the argument with the given name
func (c *Context) GetInValueAt(index int) interface{} {
	return c.GetInAt(index).value
}

// GetOut return the argument with the given name
func (c *Context) GetOut(name string) *Arg {
	return c.Out().get(name)
}

// GetOutValue return the value for the argument with the given name
func (c *Context) GetOutValue(name string) interface{} {
	return c.Out().get(name).value
}

// GetOutAt return the argument in the given position
func (c *Context) GetOutAt(index int) *Arg {
	return c.Out().at(index)
}

// ParamsLen return the number of input arguments
func (c *Context) ParamsLen() int {
	return c.In().len()
}

// HasParams returns true if there are input arguments
func (c *Context) HasParams() bool {
	return !c.In().isEmpty()
}

// ResultsLen return the number of output arguments
func (c *Context) ResultsLen() int {
	return c.Out().len()
}

// HasResults returns true if there are output arguments
func (c *Context) HasResults() bool {
	return !c.Out().isEmpty()
}

// GetOutValueAt return the argument with the given name
func (c *Context) GetOutValueAt(index int) interface{} {
	return c.GetOutAt(index).value
}

// SetOut set the value for the given argument
func (c *Context) SetOut(name string, value interface{}) {
	args := c.Out().items
	for _, arg := range args {
		if arg.name == name {
			arg.update(value)
			c.WithOut(args)

			return
		}
	}

	args = append(args, NewArg(name, value))
	c.WithOut(args)
}

// SetIn set the value for the given argument
func (c *Context) SetIn(name string, value interface{}) {
	args := c.In().items
	for _, arg := range args {
		if arg.name == name {
			arg.update(value)
			c.WithIn(args)

			return
		}
	}

	args = append(args, NewArg(name, value))
	c.WithIn(args)
}

// SetInAt set the value for the argument in the given position
func (c *Context) SetInAt(index int, value interface{}) {
	args := c.In().items
	if len(args) > index && index >= 0 {
		args[index].update(value)
	}

	c.WithIn(args)
}

// SetOutAt set the value for the argument in the given position
func (c *Context) SetOutAt(index int, value interface{}) {
	args := c.Out().items
	if len(args) > index && index >= 0 {
		args[index].update(value)
	}

	c.WithOut(args)
}
