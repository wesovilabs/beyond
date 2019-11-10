package context

import (
	"context"
)

type contextKey string

const (
	funcType contextKey = "_goaFunctionType"
	name     contextKey = "_goaFunction"
	pkg      contextKey = "_goaPkg"
	in       contextKey = "_goaIn"
	out      contextKey = "_goaOut"
)

// GoaContext Goa context
type GoaContext struct {
	ctx context.Context
}

// Pkg returns the package
func (c *GoaContext) Pkg() string {
	if v := c.ctx.Value(pkg); v != nil {
		return v.(string)
	}

	return ""
}

// Function returns the name of the function
func (c *GoaContext) Function() string {
	if v := c.ctx.Value(name); v != nil {
		return v.(string)
	}

	return ""
}

// Type returns the type
func (c *GoaContext) Type() interface{} {
	if v := c.ctx.Value(funcType); v != nil {
		return v
	}

	return nil
}

// Params returns the input arguments
func (c *GoaContext) Params() *Args {
	if v := c.ctx.Value(in); v != nil {
		return v.(*Args)
	}

	return &Args{}
}

// Results returns the output arguments
func (c *GoaContext) Results() *Args {
	if v := c.ctx.Value(out); v != nil {
		return v.(*Args)
	}

	return &Args{}
}

// NewContext constructor for goa context
func NewContext(ctx context.Context) *GoaContext {
	return &GoaContext{ctx}
}

// WithPkg set the package
func (c *GoaContext) WithPkg(v string) *GoaContext {
	c.ctx = context.WithValue(c.ctx, pkg, v)
	return c
}

// WithName set the function name
func (c *GoaContext) WithName(v string) *GoaContext {
	c.ctx = context.WithValue(c.ctx, name, v)
	return c
}

// WithType set the function type
func (c *GoaContext) WithType(v interface{}) *GoaContext {
	c.ctx = context.WithValue(c.ctx, funcType, v)
	return c
}

// WithIn set the input arguments
func (c *GoaContext) WithIn(args []*Arg) *GoaContext {
	c.ctx = context.WithValue(c.ctx, in, &Args{
		items: args,
	})

	return c
}

// WithOut set the output arguments
func (c *GoaContext) WithOut(args []*Arg) *GoaContext {
	c.ctx = context.WithValue(c.ctx, out, &Args{
		items: args,
	})

	return c
}

// Set set context value
func (c *GoaContext) Set(key string, value interface{}) {
	c.ctx = context.WithValue(c.ctx, contextKey(key), value)
}

// Get return the argument value
func (c *GoaContext) Get(key string) interface{} {
	return c.ctx.Value(contextKey(key))
}

/**
// GetString return the argument value
func (c *GoaContext) GetString(key string) string {
	if value := c.Get(key); value != nil {
		return value.(string)
	}

	return ""
}

// GetInt return the argument value
func (c *GoaContext) GetInt(key string) int {
	if value := c.Get(key); value != nil {
		return value.(int)
	}

	return 0
}

// GetBool return the argument value
func (c *GoaContext) GetBool(key string) bool {
	if value := c.Get(key); value != nil {
		return value.(bool)
	}

	return false
}

// GetTime return the argument value
func (c *GoaContext) GetTime(key string) time.Time {
	if value := c.Get(key); value != nil {
		return value.(time.Time)
	}

	return time.Time{}
}



// GetIn return the argument with the given name
func (c *GoaContext) GetIn(name string) *Arg {
	return c.Params().get(name)
}

// GetInValue return the argument with the given name
func (c *GoaContext) GetInValue(name string) interface{} {
	return c.Params().get(name).value
}

// GetInAt return the argument in the given position
func (c *GoaContext) GetInAt(index int) *Arg {
	return c.Params().at(index)
}

// GetInValueAt return the argument with the given name
func (c *GoaContext) GetInValueAt(index int) interface{} {
	return c.GetInAt(index).value
}

// GetOut return the argument with the given name
func (c *GoaContext) GetOut(name string) *Arg {
	return c.Results().get(name)
}

// GetOutValue return the value for the argument with the given name
func (c *GoaContext) GetOutValue(name string) interface{} {
	return c.Results().get(name).value
}

// GetOutAt return the argument in the given position
func (c *GoaContext) GetOutAt(index int) *Arg {
	return c.Results().at(index)
}

// ParamsLen return the number of input arguments
func (c *GoaContext) ParamsLen() int {
	return c.Params().len()
}

// HasParams returns true if there are input arguments
func (c *GoaContext) HasParams() bool {
	return !c.Params().isEmpty()
}

// ResultsLen return the number of output arguments
func (c *GoaContext) ResultsLen() int {
	return c.Results().len()
}

// HasResults returns true if there are output arguments
func (c *GoaContext) HasResults() bool {
	return !c.Results().isEmpty()
}

// GetOutValueAt return the argument with the given name
func (c *GoaContext) GetOutValueAt(index int) interface{} {
	return c.GetOutAt(index).value
}

// SetOut set the value for the given argument
func (c *GoaContext) SetOut(name string, value interface{}) {
	args := c.Results().items
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
func (c *GoaContext) SetIn(name string, value interface{}) {
	args := c.Params().items
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
func (c *GoaContext) SetInAt(index int, value interface{}) {
	args := c.Params().items
	if len(args) > index && index >= 0 {
		args[index].update(value)
	}

	c.WithIn(args)
}

// UpdateResultAt set the value for the argument in the given position
func (c *GoaContext) UpdateResultAt(index int, value interface{}) {
	args := c.Results().items
	if len(args) > index && index >= 0 {
		args[index].update(value)
	}

	c.WithOut(args)
}
**/
