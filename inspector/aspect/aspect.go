package aspect

import (
	"github.com/wesovilabs/goa/logger"
	"regexp"
)

// Aspects array of aspect
type Aspects []*Aspect

// Aspect attributes for an aspect
type Aspect struct {
	pkg     string
	name    string
	pattern string
	regExp  *regexp.Regexp
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

// WithPattern sets the expression
func (a *Aspect) WithPattern(pattern string) (*Aspect, error) {
	a.pattern = pattern
	if regExp, err := regexp.Compile(pattern); err != nil {
		return nil, err
	} else {
		a.regExp = regExp
	}
	return a, nil
}

// Pkg returns the package name
func (a *Aspect) Pkg() string {
	return a.pkg
}

// Name returns the function name
func (a *Aspect) Name() string {
	return a.name
}

// Match checks if the fiven text matches with the aspect expression
func (a *Aspect) Match(text string) bool {
	logger.Infof("[check] %s with %s", a.regExp.String(), text)
	return a.regExp.MatchString(text)
}
