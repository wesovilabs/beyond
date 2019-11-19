package context

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
	ctx map[contextKey]interface{}
}

// Pkg returns the package
func (c *GoaContext) Pkg() string {
	if v := c.ctx[pkg]; v != nil {
		return v.(string)
	}

	return ""
}

// Function returns the name of the function
func (c *GoaContext) Function() string {
	if v := c.ctx[name]; v != nil {
		return v.(string)
	}

	return ""
}

// Type returns the type
func (c *GoaContext) Type() interface{} {
	if v := c.ctx[funcType]; v != nil {
		return v
	}

	return nil
}

// Params returns the input arguments
func (c *GoaContext) Params() *Args {
	if v := c.ctx[in]; v != nil {
		return v.(*Args)
	}

	return &Args{}
}

// Results returns the output arguments
func (c *GoaContext) Results() *Args {
	if v := c.ctx[out]; v != nil {
		return v.(*Args)
	}

	return &Args{}
}

// NewContext constructor for goa context
func NewContext() *GoaContext {
	return &GoaContext{make(map[contextKey]interface{})}
}

// WithPkg set the package
func (c *GoaContext) WithPkg(v string) *GoaContext {
	c.ctx[pkg] = v
	return c
}

// WithName set the function name
func (c *GoaContext) WithName(v string) *GoaContext {
	c.ctx[name] = v
	return c
}

// WithType set the function type
func (c *GoaContext) WithType(v interface{}) *GoaContext {
	c.ctx[funcType] = v
	return c
}

// SetParams set the input arguments
func (c *GoaContext) SetParams(args *Args) *GoaContext {
	c.ctx[in] = args
	return c
}

// SetResults set the output arguments
func (c *GoaContext) SetResults(args *Args) *GoaContext {
	c.ctx[out] = args
	return c
}

// Set set context value
func (c *GoaContext) Set(key string, value interface{}) {
	c.ctx[contextKey(key)] = value
}

// Get return the argument value
func (c *GoaContext) Get(key string) interface{} {
	return c.ctx[contextKey(key)]
}
