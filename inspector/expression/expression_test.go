package expression

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func Test_processExprStr(t *testing.T) {
	cases := []struct {
		text       string
		expression *expression
	}{
		{
			text: "_package.build(string)int",
			expression: &expression{
				pkg:      "_package",
				instance: "",
				function: "build",
				in:       "string",
				out:      "int",
			},
		},
		{
			text: "pkg/any.any(any)any",
			expression: &expression{
				pkg:      "pkg/any",
				instance: "",
				function: "any",
				in:       "any",
				out:      "any",
			},
		},
		{
			text: "pkg/any.any.any(string)any",
			expression: &expression{
				pkg:      "pkg/any",
				instance: "any",
				function: "any",
				in:       "string",
				out:      "any",
			},
		},
		{
			text: "any.any(int,string)*",
			expression: &expression{
				pkg:      "any",
				instance: "",
				function: "any",
				in:       "int,string",
				out:      "*",
			},
		},
		{
			text: "parent/child.instance.func(int,map[string]int)([]*string,*int)",
			expression: &expression{
				pkg:      "parent/child",
				instance: "instance",
				function: "func",
				in:       "int,map[string]int",
				out:      "[]*string,*int",
			},
		},
		{
			text: "any.any(int,string)(string,int)",
			expression: &expression{
				pkg:      "any",
				instance: "",
				function: "any",
				in:       "int,string",
				out:      "string,int",
			},
		},
		{
			text: "an-y/a_ny.obj_as.test(int,person.Person)(string,int)",
			expression: &expression{
				pkg:      "an-y/a_ny",
				instance: "obj_as",
				function: "test",
				in:       "int,person.Person",
				out:      "string,int",
			},
		},
	}
	for _, c := range cases {
		exp := processExprStr(c.text)
		assert.NotNil(t, exp)
		assert.Equal(t, c.expression.pkg, exp.pkg)
		assert.Equal(t, c.expression.instance, exp.instance)
		assert.Equal(t, c.expression.function, exp.function)
		assert.Equal(t, c.expression.in, exp.in)
		assert.Equal(t, c.expression.out, exp.out)
	}

}

func Test_calculateRegExps(t *testing.T) {
	cases := []struct {
		expr   *expression
		regExp *regexp.Regexp
	}{
		{
			expr: &expression{
				pkg:      "*",
				instance: "",
				function: "func",
				in:       "string",
				out:      "any",
			},
			regExp: regexp.MustCompile(`\w+\.func\(string\)any`),
		},
		{
			expr: &expression{
				pkg:      "pkg/*",
				instance: "*",
				function: "func",
				in:       "string",
				out:      "any",
			},
			regExp: regexp.MustCompile(`pkg/\w+\.\w+\.func\(string\)any`),
		},
		{
			expr: &expression{
				pkg:      "pkg/*",
				instance: "obj",
				function: "*",
				in:       "",
				out:      "",
			},
			regExp: regexp.MustCompile(`pkg/\w+\.obj\.\w+\(\)`),
		},
		{
			expr: &expression{
				pkg:      "pkg/*",
				instance: "obj",
				function: "*",
				in:       "",
				out:      "string,*int,string",
			},
			regExp: regexp.MustCompile(`pkg/\w+\.obj\.\w+\(\)\(string\,\*int\,string\)`),
		},
		{
			expr: &expression{
				pkg:      "pkg/*",
				instance: "obj",
				function: "*",
				in:       "",
				out:      "string,*,string",
			},
			regExp: regexp.MustCompile(`pkg/\w+\.obj\.\w+\(\)\(string\,[\w\.]+\,string\)`),
		},
		{
			expr: &expression{
				pkg:      "pkg/*",
				instance: "obj",
				function: "*",
				in:       "",
				out:      "[]string",
			},
			regExp: regexp.MustCompile(`pkg/\w+\.obj\.\w+\(\)\[\]string`),
		},
		{
			expr: &expression{
				pkg:      "pkg/*",
				instance: "obj",
				function: "*",
				in:       "[]int,func(int,string)*int",
				out:      "",
			},
			regExp: regexp.MustCompile(`pkg/\w+\.obj\.\w+\(\[\]int\,func\(int\,string\)\*int\)`),
		},
	}
	for _, c := range cases {
		regExp(c.expr)
		assert.Equal(t, c.regExp, c.expr.regExp)
	}
}
