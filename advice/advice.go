package advice

import (
	"regexp"
)

type adviceType int32

// Advice struct
type Advice struct {
	pkg    string
	name   string
	kind   adviceType
	regExp *regexp.Regexp
}

// HasBefore return if before function is implemented
func (a *Advice) HasBefore() bool {
	return a.kind == around || a.kind == before
}

// HasReturning return if returning function is implemented
func (a *Advice) HasReturning() bool {
	return a.kind == around || a.kind == returning
}

// Match return if the given input matches with the definition
func (a *Advice) Match(text string) bool {
	if a.regExp == nil {
		return false
	}

	return a.regExp.MatchString(text)
}

// Pkg return the package
func (a *Advice) Pkg() string {
	return a.pkg
}

// Name return the function name
func (a *Advice) Name() string {
	return a.name
}

// Advices struct
type Advices struct {
	items []*Advice
}

// List return the list of definitions
func (a *Advices) List() []*Advice {
	return a.items
}

// Add add a new definition
func (a *Advices) Add(def *Advice) {
	a.items = append(a.items, def)
}
