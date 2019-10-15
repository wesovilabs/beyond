package context

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInput(t *testing.T) {
	cases := []struct {
		args  []*Arg
		len   int
		empty bool
	}{
		{
			args: []*Arg{
				NewArg("name", "John"),
			},
			len:   1,
			empty: false,
		},
		{
			args: []*Arg{
				NewArg("name", "John"),
				NewArg("age", 20),
			},
			len:   2,
			empty: false,
		},
		{
			args:  []*Arg{},
			len:   0,
			empty: true,
		},
	}
	for _, c := range cases {
		input := Input(c.args)
		assert.Len(t, input, c.len)
		assert.Equal(t, input.isEmpty(), c.empty)
	}
}
