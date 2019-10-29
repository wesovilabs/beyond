package aspect

import (
	"regexp"
)

type definitionKind int32

type Definition struct {
	pkg    string
	name   string
	kind   definitionKind
	regExp *regexp.Regexp
}

func (d *Definition) HasBefore() bool {
	return d.kind == around || d.kind == before
}

func (d *Definition) HasReturning() bool {
	return d.kind == around || d.kind == returning
}

func (d *Definition) Match(text string) bool {
	return d.regExp.MatchString(text)
}

func (a *Definition) Pkg() string {
	return a.pkg
}

func (d *Definition) Name() string {
	return d.name
}

type Definitions struct {
	items []*Definition
}

func (d *Definitions) List() []*Definition {
	return d.items
}

func (d *Definitions) Add(def *Definition) {
	d.items = append(d.items, def)
}
