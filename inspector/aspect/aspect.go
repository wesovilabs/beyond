package aspect

import (
	"github.com/wesovilabs/goa/inspector/expression"
)

type Aspects []*Aspect

type Aspect struct {
	pkg  string
	name string
	expr *expression.Expression
}

func (a *Aspect) WithPkg(pkg string) *Aspect {
	a.pkg = pkg
	return a
}

func (a *Aspect) WithName(name string) *Aspect {
	a.name = name
	return a
}

func (a *Aspect) WithExpr(expr *expression.Expression) *Aspect {
	a.expr = expr
	return a
}

func (a *Aspect) Pkg() string {
	return a.pkg
}
func (a *Aspect) Name() string {
	return a.name
}
func (a *Aspect) Expression() *expression.Expression {
	return a.expr
}

func (a *Aspect) Match(text string) bool {
	return a.expr.Match(text)
}
