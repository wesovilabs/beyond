package aspect

import (
	"github.com/wesovilabs/goa/inspector/expression"
)

// Aspects array of aspect
type Aspects []*Aspect

// Aspect attributes for an aspect
type Aspect struct {
	pkg  string
	name string
	expr *expression.Expression
}

// WithPkg sets the package name
func (a *Aspect) WithPkg(pkg string) *Aspect {
	a.pkg = pkg
	return a
}

// WithName sets the function name
func (a *Aspect) WithName(name string) *Aspect {
	a.name = name
	return a
}

// WithExpr sets the expression
func (a *Aspect) WithExpr(expr *expression.Expression) *Aspect {
	a.expr = expr
	return a
}

// Pkg returns the package name
func (a *Aspect) Pkg() string {
	return a.pkg
}

// Name returns the function name
func (a *Aspect) Name() string {
	return a.name
}

// Expression returns the expression
func (a *Aspect) Expression() *expression.Expression {
	return a.expr
}

// Match checks if the fiven text matches with the aspect expression
func (a *Aspect) Match(text string) bool {
	return a.expr.Match(text)
}
