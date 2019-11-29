package context

type contextKey string

const (
	funcType contextKey = "_beyondFunctionType"
	name     contextKey = "_beyondFunction"
	pkg      contextKey = "_beyondPkg"
	in       contextKey = "_beyondIn"
	out      contextKey = "_beyondOut"
)

// BeyondContext Beyond context
type BeyondContext struct {
	ctx map[contextKey]interface{}
	completed bool
}

// Pkg returns the package
func (c *BeyondContext) Pkg() string {
	if v := c.ctx[pkg]; v != nil {
		return v.(string)
	}

	return ""
}

// Function returns the name of the function
func (c *BeyondContext) Function() string {
	if v := c.ctx[name]; v != nil {
		return v.(string)
	}

	return ""
}

// Type returns the type
func (c *BeyondContext) Type() interface{} {
	if v := c.ctx[funcType]; v != nil {
		return v
	}

	return nil
}

// Params returns the input arguments
func (c *BeyondContext) Params() *Args {
	if v := c.ctx[in]; v != nil {
		return v.(*Args)
	}

	return &Args{}
}

// Results returns the output arguments
func (c *BeyondContext) Results() *Args {
	if v := c.ctx[out]; v != nil {
		return v.(*Args)
	}

	return &Args{}
}

// NewContext constructor for beyond context
func NewContext() *BeyondContext {
	return &BeyondContext{make(map[contextKey]interface{}),false}
}

// WithPkg set the package
func (c *BeyondContext) WithPkg(v string) *BeyondContext {
	c.ctx[pkg] = v
	return c
}

// WithName set the function name
func (c *BeyondContext) WithName(v string) *BeyondContext {
	c.ctx[name] = v
	return c
}

// WithType set the function type
func (c *BeyondContext) WithType(v interface{}) *BeyondContext {
	c.ctx[funcType] = v
	return c
}

// SetParams set the input arguments
func (c *BeyondContext) SetParams(args *Args) *BeyondContext {
	c.ctx[in] = args
	return c
}

// SetResults set the output arguments
func (c *BeyondContext) SetResults(args *Args) *BeyondContext {
	c.ctx[out] = args
	return c
}

// Set set context value
func (c *BeyondContext) Set(key string, value interface{}) {
	c.ctx[contextKey(key)] = value
}

// Get return the argument value
func (c *BeyondContext) Get(key string) interface{} {
	return c.ctx[contextKey(key)]
}

// Exit won't call joinpoint
func (c *BeyondContext) Exit() {
	c.completed=true
}

// IsCompleted returns true if flow must stop
func (c *BeyondContext) IsCompleted() bool {
	return c.completed
}
