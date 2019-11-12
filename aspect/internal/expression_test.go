package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NormalizeExpression(t *testing.T) {

	cases := []struct {
		text    string
		pattern string
	}{

		{
			text:    "a.b()",
			pattern: `^a\.b\(\)$`,
		},
		{
			text:    "a.b.c()",
			pattern: `^a\.b\.c\(\)$`,
		},
		{
			text:    "a/b.c()",
			pattern: `^a\/b\.c\(\)$`,
		},
		{
			text:    "a/b.c.d()",
			pattern: `^a\/b\.c\.d\(\)$`,
		},
		{
			text:    "a/b.c.d(string)",
			pattern: `^a\/b\.c\.d\(string\)$`,
		},
		{
			text:    "a/b.c.d(string,*int)",
			pattern: `^a\/b\.c\.d\(string\,\*int\)$`,
		},
		{
			text:    "a/b.c.d()string",
			pattern: `^a\/b\.c\.d\(\)string$`,
		},
		{
			text:    "a/b.c.d()(string,[]int)",
			pattern: `^a\/b\.c\.d\(\)\(string\,\[\]int\)$`,
		},
		{
			text:    "a/b.c.d(person.Person)",
			pattern: `^a\/b\.c\.d\(person\.Person\)$`,
		},
		{
			text:    "a/b.c.d(map[string]*person.Person)",
			pattern: `^a\/b\.c\.d\(map\[string\]\*person\.Person\)$`,
		},
		{
			text:    "a/b.c.d(string,func()string)",
			pattern: `^a\/b\.c\.d\(string\,func\(\)string\)$`,
		},
		{
			text:    "a/b.c.d(string,func()string)",
			pattern: `^a\/b\.c\.d\(string\,func\(\)string\)$`,
		},
		{
			text:    "a/b.c.d(string,func()string)func(int)string",
			pattern: `^a\/b\.c\.d\(string\,func\(\)string\)func\(int\)string$`,
		},
		{
			text:    "*.b()",
			pattern: `^[a-zA-Z0-9_\\/]+\.b\(\)$`,
		},
		{
			text:    "*.*()",
			pattern: `^[a-zA-Z0-9_\\/]+\.[a-zA-Z0-9_]+\(\)$`,
		},
		{
			text:    "a.*()",
			pattern: `^a\.[a-zA-Z0-9_]+\(\)$`,
		},
		{
			text:    "a/test/*.*()",
			pattern: `^a\/test\/[a-zA-Z0-9_\\/]+\.[a-zA-Z0-9_]+\(\)$`,
		},
		{
			text:    "a/test/*.*.*()",
			pattern: `^a\/test\/[a-zA-Z0-9_\\/]+\.[a-zA-Z0-9_*]+\.[a-zA-Z0-9_]+\(\)$`,
		},
		{
			text:    "a.b(*)",
			pattern: `^a\.b\([a-zA-Z0-9_*.\[\]{}/]+\)$`,
		},
		{
			text:    "a.b(...)",
			pattern: `^a\.b\(.*\)$`,
		},
		{
			text:    "a.b(string,...)",
			pattern: `^a\.b\(string\,.*\)$`,
		},

		{
			text:    "a.b()(...)",
			pattern: `^a\.b\(\).*$`,
		},
		{
			text:    "a.b()(string,...)",
			pattern: `^a\.b\(\)\(string\,.*\)$`,
		},
		{
			text:    "a.b()*",
			pattern: `^a\.b\(\)[a-zA-Z0-9_*.\[\]{}/]+$`,
		},
		{
			text:    "a.b()*string",
			pattern: `^a\.b\(\)\*string$`,
		},
		{
			text:    "a/b.c.d(string,func()string)func()int",
			pattern: `^a\/b\.c\.d\(string\,func\(\)string\)func\(\)int$`,
		},
	}
	asssert := assert.New(t)
	for _, c := range cases {
		regExp := NormalizeExpression(c.text)
		asssert.Equal(c.pattern, regExp.String())
	}
}
