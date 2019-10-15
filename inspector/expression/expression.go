package expression

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	pkgExp  = `(?P<Pkg>[|a-zA-Z0-9_*]+)`
	nameExp = `(?P<func>[|a-zA-Z0-9_*]+)`
	inExp   = `(?P<in>[|()a-zA-Z0-9_*.,{} ]+)`
	outExp  = `(?P<out>[|a-zA-Z0-9_*., ]+)`
	any     = "(.*)"
)

var fegExpStr = fmt.Sprintf(`^%s\.%s\(%s\)[ ]?\(?%s\)?$`, pkgExp, nameExp, inExp, outExp)

var regExp = regexp.MustCompile(fegExpStr)

type args []*arg

type arg struct {
	name      string
	kind      string
	isPointer bool
}

// Expression contains the required attributes to define an expression
type Expression struct {
	pattern string
	regExp  *regexp.Regexp
	pkg     string
	name    string
	in      args
	out     args
}

func (e *Expression) Match(value string) bool {
	return e.regExp.MatchString(value)
}

func newArg(text string) (*arg, error) {
	if parts := strings.Split(text, " "); len(parts) == 2 {
		kind := strings.TrimPrefix(parts[1], "*")
		return &arg{
			name:      parts[0],
			kind:      kind,
			isPointer: kind != parts[1],
		}, nil
	} else if len(parts) == 1 {
		kind := strings.TrimPrefix(parts[0], "*")
		return &arg{
			kind:      kind,
			isPointer: kind != parts[0],
		}, nil
	}
	return nil, errors.New("invalid arg")
}

func newArgs(text string) (args, error) {
	args := args{}
	for _, argStr := range strings.Split(text, ",") {
		arg, err := newArg(argStr)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}
	return args, nil
}

// NewExpression create an instance of Expression from a given text
func NewExpression(text string) (*Expression, error) {
	vars := regExp.FindStringSubmatch(text)
	if len(vars) < 4 {
		return nil, errors.New("invalid Expression")
	}
	in, err := newArgs(vars[3])
	if err != nil {
		return nil, err
	}
	out, err := newArgs(vars[4])
	if err != nil {
		return nil, err
	}

	pattern := buildPattern(vars[1], vars[2], in, out)
	pattern = strings.ReplaceAll(text, "*", any)
	pattern = strings.ReplaceAll(pattern, ",", "\\,")
	regExp, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	return &Expression{
		pattern: pattern,
		regExp:  regExp,
		pkg:     vars[1],
		name:    vars[2],
		in:      in,
		out:     out,
	}, nil
}

func buildPattern(pkg, name string, in, out args) string {
	regExpStr := ""
	if pkg == "*" {
		regExpStr += any
	} else {
		regExpStr += pkg
	}
	if name == "*" {
		regExpStr += any
	} else {
		regExpStr += name
	}
	regExpStr += "("
	for index, i := range in {
		if i.kind == "*" {
			regExpStr += any
		} else {
			regExpStr += i.kind
		}
		if index < len(in)-1 {
			regExpStr += ","
		}
	}
	regExpStr += ")"
	regExpStr += "("

	for index, i := range out {
		if i.kind == "*" || i.kind == "" {
			regExpStr += any
		}
		if index < len(out)-1 {
			regExpStr += ","
		}
	}
	regExpStr += ")"
	regExpStr = strings.ReplaceAll(regExpStr, ",", "\\,")
	return regExpStr
}
