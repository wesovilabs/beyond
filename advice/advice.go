package advice

import (
	"fmt"
	"regexp"
	"strings"
)

type adviceType int32

type adviceInvocationArg struct {
	pkg     string
	val     string
	pointer bool
}

type adviceInvocation struct {
	imports  []string
	pkg      string
	function string
	args     []*adviceInvocationArg
	isCall   bool
}

func (in *adviceInvocation) addImport(path string) {
	if in.imports == nil {
		in.imports = make([]string, 0)
	}

	for _, p := range in.imports {
		if p == path {
			return
		}
	}

	in.imports = append(in.imports, path)
}

// Advice struct
type Advice struct {
	call   *adviceInvocation
	kind   adviceType
	regExp *regexp.Regexp
}

func (a *Advice) Imports() []string {
	if a.call.imports == nil {
		return []string{a.call.pkg}
	}

	return append(a.call.imports, a.call.pkg)
}

// HasBefore return if before function is implemented
func (a *Advice) HasBefore() bool {
	return a.kind == around || a.kind == before
}

// HasReturning return if returning function is implemented
func (a *Advice) HasReturning() bool {
	return a.kind == around || a.kind == returning
}

// Match return if the given input matches with the advice
func (a *Advice) Match(text string) bool {
	if a.regExp == nil {
		return false
	}

	return a.regExp.MatchString(text)
}

// Pkg return the package
func (a *Advice) Pkg() string {
	return a.call.pkg
}

// Name return the function name
func (a *Advice) Name() string {
	return a.call.function
}

// GetAdviceCall returns the advice call
func (a *Advice) GetAdviceCall(currentPkg string, imports map[string]string) string {
	args := make([]string, len(a.call.args))

	for index := range a.call.args {
		arg := a.call.args[index]

		if arg.pkg != "" && currentPkg != arg.pkg {
			pkgName := imports[arg.pkg]
			args[index] = fmt.Sprintf("%s.%s", pkgName, arg.val)
		} else {
			args[index] = arg.val
		}

		if arg.pointer {
			args[index] = fmt.Sprintf("&%s", args[index])
		}
	}

	if a.call.isCall {
		return fmt.Sprintf("%s(%s)()", a.call.function, strings.Join(args, ","))
	}

	return fmt.Sprintf("%s(%s)", a.call.function, strings.Join(args, ","))
}

// Advices struct
type Advices struct {
	items []*Advice
}

// List return the list of advices
func (a *Advices) List() []*Advice {
	return a.items
}

// Add add a new advice
func (a *Advices) Add(def *Advice) {
	a.items = append(a.items, def)
}
