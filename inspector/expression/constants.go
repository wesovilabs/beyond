package expression

import (
	"fmt"
	"regexp"
	"strings"
)

const types = `[a-zA-Z0-9_\.]+`
const maps = `map\[` + types + `\]` + types
const arrays = `\[\]` + types
const lists = `\.\.\.` + types
const structs = `struct{}`
const interfaces = `interface{}`
const pointers = `\*` + types

const all = types + "|" + maps + "|" + arrays + "|" + lists + "|" + structs + "|" + interfaces + "|" + pointers

const funcs = `func\(` + all + "|" + `func\(` + all + `\)\(?` + all + `\)?` + `\)\(?` + all + `\)?`

const packages = `[a-zA-Z0-9_*/]+`
const instances = `[a-zA-Z0-9_*]+\.`
const names = `[a-zA-Z0-9_*]+\.`

var funcRegExp = buildRegExp()

func buildRegExp() *regexp.Regexp {
	exprStr := `^(` + `?P<package>` + packages + `)\.`
	exprStr += `(` + `?P<instance>` + instances + `)?`
	exprStr += `(` + `?P<function>` + names + `)`
	exprStr += `\((` + `?P<argIn>` + all + `)\)`
	exprStr += `[ ]*`
	exprStr += `(` + `?P<argOut>` + all + `)?$`
	return regexp.MustCompile(exprStr)
}

func evaluate(value string) *expression {
	fmt.Println(funcRegExp.String())
	items := funcRegExp.FindStringSubmatch(value)
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
