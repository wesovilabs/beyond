package aspect

import (
	"regexp"
)

type definitionKind int32

// Definition struct
type Definition struct {
	pkg    string
	name   string
	kind   definitionKind
	regExp *regexp.Regexp
}

// HasBefore return if before function is implemented
func (d *Definition) HasBefore() bool {
	return d.kind == around || d.kind == before
}

// HasReturning return if returning function is implemented
func (d *Definition) HasReturning() bool {
	return d.kind == around || d.kind == returning
}

// Match return if the given input matches with the definition
func (d *Definition) Match(text string) bool {
	if d.regExp == nil {
		return false
	}

	return d.regExp.MatchString(text)
}

// Pkg return the package
func (d *Definition) Pkg() string {
	return d.pkg
}

// Name return the function name
func (d *Definition) Name() string {
	return d.name
}

// Definitions struct
type Definitions struct {
	items []*Definition
}

// List return the list of definitions
func (d *Definitions) List() []*Definition {
	return d.items
}

// Add add a new definition
func (d *Definitions) Add(def *Definition) {
	d.items = append(d.items, def)
}
