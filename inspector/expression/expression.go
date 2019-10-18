package expression

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const pkgChars = `[a-zA-Z0-9_\-*/]+`
const modChars = `[a-zA-Z0-9_\-*]+\.`
const funcChars = `[a-zA-Z0-9_\-*]+`
const inChars = `[a-zA-Z0-9\[\]_\-*\,\.]*`
const outChars = `[a-zA-Z0-9_\[\]\-*\,\.]*|\([a-zA-Z0-9_\[\]\-*\,\.]*\)`

func exprStr() string {
	out := `^(` + `?P<pkg>` + pkgChars + `)\.`
	out += `(` + `?P<mod>` + modChars + `)?`
	out += `(` + `?P<func>` + funcChars + `)`
	out += `\((` + `?P<in>` + inChars + `)\)`
	out += `[ ]*`
	out += `(` + `?P<out>` + outChars + `)?$`
	return out
}

var expRegExp = regexp.MustCompile(exprStr())

// Expression contains the required attributes to define an expression
type Expression struct {
	expr *expression
}

type expression struct {
	pkg      string
	instance string
	function string
	in       string
	out      string
	regExp   *regexp.Regexp
}

func (e *expression) match(value string) bool {
	return e.regExp.MatchString(value)
}

// Match checks if value match with expression
func (e *Expression) Match(value string) bool {
	return e.expr.match(value)
}

const (
	lParen = "\\("
	rParen = "\\)"
)

func betweenParen(value string) string {
	return lParen + value + rParen
}

func anyWord(args ...string) string {
	if len(args) == 0 {
		return `\w+`
	}
	out := `\w`
	for _, arg := range args {
		out += arg
	}
	return fmt.Sprintf("[%s]+", out)
}

func expandPkg(value string) string {
	blocks := strings.Split(value, "/")
	out := ""
	for index, b := range blocks {
		if b == "*" {
			out += anyWord()
		} else {
			out += b
		}
		if index < len(blocks)-1 {
			out += `/`
		}
	}
	return out
}

func replaceSpecialCharacters(value string) string {
	output := strings.ReplaceAll(value, ".", `\.`)
	output = strings.ReplaceAll(output, "[", `\[`)
	output = strings.ReplaceAll(output, "]", `\]`)
	output = strings.ReplaceAll(output, "*", `\*`)
	output = strings.ReplaceAll(output, "(", `\(`)
	output = strings.ReplaceAll(output, ")", `\)`)
	return output
}

func expandArgs(value string) (string, int) {
	blocks := strings.Split(value, ",")
	out := ""
	c := 0
	for index, b := range blocks {
		if b == "*" {
			out += anyWord(`\.`)
		} else {
			out += replaceSpecialCharacters(b)
		}
		if index < len(blocks)-1 {
			out += `\,`
		}
		c++
	}
	return out, c
}

func expandInstance(value string) string {
	if value == "*" {
		return anyWord()
	}
	return value
}

func expandFunction(value string) string {
	if value == "*" {
		return anyWord()
	}
	return value
}

func regExp(e *expression) {
	value := expandPkg(e.pkg)
	value += `\.`
	if e.instance != "" {
		value += expandInstance(e.instance)
		value += `\.`
	}
	value += expandFunction(e.function)
	inValue, _ := expandArgs(e.in)
	value += betweenParen(inValue)
	outValue, outArgsLen := expandArgs(e.out)
	if outArgsLen > 1 {
		outValue = betweenParen(outValue)
	}
	value += outValue
	e.regExp = regexp.MustCompile(value)
}

func processExprStr(text string) *expression {
	items := expRegExp.FindStringSubmatch(text)
	if len(items) != 6 {
		return nil
	}
	instance := strings.TrimRight(items[2], ".")
	out := func(val string) string {
		outLen := len(val)
		if outLen >= 2 && strings.HasPrefix(val, "(") && strings.HasSuffix(val, ")") {
			return val[1 : outLen-1]
		}
		return val
	}(items[5])

	e := &expression{
		pkg:      items[1],
		instance: instance,
		function: items[3],
		in:       items[4],
		out:      out,
		regExp:   nil,
	}
	regExp(e)
	return e
}

// NewExpression create an instance of Expression from a given text
func NewExpression(text string) (*Expression, error) {
	expr := processExprStr(text)
	if expr == nil {
		return nil, errors.New("invalid expression")
	}
	return &Expression{
		expr: expr,
	}, nil
}
