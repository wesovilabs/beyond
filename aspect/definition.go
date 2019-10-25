package aspect

import (
	"github.com/wesovilabs/goa/logger"
	"regexp"
)

type definitionKind int32

type Definition struct {
	pkgPath string
	name    string
	kind    definitionKind
	regExp  *regexp.Regexp
}

func (d *Definition) Match(text string) bool {
	logger.Infof("matching \"%s\" with \"%s\"", d.regExp.String(), text)
	return d.regExp.MatchString(text)
}

func (a *Definition) Pkg() string {
	return a.pkgPath
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
