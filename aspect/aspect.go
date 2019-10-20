package aspect

import (
	"github.com/wesovilabs/goa/logger"
	"regexp"
)

type Aspects struct {
	aroundList []*AroundAspect
}

func (a *Aspects) WithAround(aspect *AroundAspect) {
	a.aroundList = append(a.aroundList, aspect)
}

func (a *Aspects) AroundAspects() []*AroundAspect {
	return a.aroundList
}

type Aspect interface {
	Match(text string) bool
	Pkg() string
	Name() string
}

type AroundAspect struct {
	regExp   *regexp.Regexp
	pkgPath  string
	funcName string
}

func (a *AroundAspect) with(pkgPath, funcName string) *AroundAspect {
	a.funcName = funcName
	a.pkgPath = pkgPath
	return a
}

func (a *AroundAspect) Match(text string) bool {
	logger.Infof("matching \"%s\" with \"%s\"", a.regExp.String(), text)
	return a.regExp.MatchString(text)
}

func (a *AroundAspect) Pkg() string {
	return a.pkgPath
}

func (a *AroundAspect) Name() string {
	return a.funcName
}
