package advice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/advice/internal"
	goaParser "github.com/wesovilabs/goa/parser"
	"testing"
)

const testdataPath = "testdata"
const rootPkg = "github.com/wesovilabs/goa/advice/testdata/a"

func TestGetAdvices(t *testing.T) {
	cases := []struct {
		directory   string
		definitions *Advices
	}{
		{
			directory: "a",
			definitions: &Advices{
				items: []*Advice{
					{
						pkg:    "github.com/wesovilabs/goa/advice/testdata/a/a2",
						name:   "AroundA2",
						kind:   around,
						regExp: internal.NormalizeExpression("*.*(*)..."),
					},

					{
						pkg:    "github.com/wesovilabs/goa/advice/testdata/a/a1/a11",
						name:   "AroundA11",
						kind:   around,
						regExp: internal.NormalizeExpression("*.*(*)..."),
					},
					{
						pkg:    "github.com/wesovilabs/goa/advice/testdata/a",
						name:   "NewTracingReturning",
						kind:   returning,
						regExp: internal.NormalizeExpression("*.*(*)..."),
					},
					{
						pkg:    "github.com/wesovilabs/goa/advice/testdata/a",
						name:   "NewTracingBefore",
						kind:   before,
						regExp: internal.NormalizeExpression("*.*(*)..."),
					},
					{
						pkg:    "github.com/wesovilabs/goa/advice/testdata/a",
						name:   "Around",
						kind:   around,
						regExp: internal.NormalizeExpression("*.*(*)..."),
					}, {
						pkg:    "github.com/wesovilabs/goa/advice/testdata/a/a1",
						name:   "AroundA1",
						kind:   around,
						regExp: internal.NormalizeExpression("*.*(*)..."),
					},
				},
			},
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		packages := goaParser.
			New(fmt.Sprintf("%s/%s", testdataPath, c.directory), fmt.Sprintf("github.com/wesovilabs/goa/advice/testdata/%s", c.directory)).
			Parse("")
		assert.NotNil(packages)
		defs := GetAdvices(rootPkg, packages)
		assert.Len(c.definitions.items, len(defs.items))
		for index, definition := range c.definitions.List() {
			assert.Equal(definition.pkg, defs.items[index].Pkg())
			assert.Equal(definition.kind, defs.items[index].kind)
			assert.Equal(definition.regExp.String(), defs.items[index].regExp.String())
			assert.Equal(definition.name, defs.items[index].Name())

			if definition.kind == around || definition.kind == before {
				assert.True(definition.HasBefore())
			}
			if definition.kind == around || definition.kind == returning {
				assert.True(definition.HasReturning())
			}
		}
	}
}
