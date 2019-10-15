package expression

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewExpression(t *testing.T) {
	cases := []struct {
		text    string
		pkg     string
		name    string
		pattern string
		in, out args
	}{
		{
			text:    "*.*(*)*",
			pkg:     "*",
			name:    "*",
			pattern: "(.*).(.*)((.*))(.*)",
			in:      args{{name: "", kind: "", isPointer: true}},
			out:     args{{name: "", kind: "", isPointer: true}},
		},
		{
			text:    "*.*(string)*",
			pkg:     "*",
			name:    "*",
			pattern: "(.*).(.*)(string)(.*)",
			in:      args{{name: "", kind: "string", isPointer: false}},
			out:     args{{name: "", kind: "", isPointer: true}},
		},
		{
			text:    "*.*(int,string)*",
			pkg:     "*",
			name:    "*",
			pattern: "(.*).(.*)(int\\,string)(.*)",
			in: args{
				{name: "", kind: "int", isPointer: false},
				{name: "", kind: "string", isPointer: false},
			},
			out: args{{name: "", kind: "", isPointer: true}},
		},
	}
	for _, c := range cases {
		exp, err := NewExpression(c.text)
		assert.Nil(t, err)
		assert.NotNil(t, exp)
		assert.Equal(t, c.pkg, exp.pkg)
		assert.Equal(t, c.name, exp.Name)
		assert.Equal(t, c.pattern, exp.pattern)
		assert.EqualValues(t, c.in, exp.in)
		assert.EqualValues(t, c.out, exp.out)

	}

}
