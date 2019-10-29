package aspect

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/aspect/internal"
	goaParser "github.com/wesovilabs/goa/parser"
	"testing"
)

const testdataPath = "testdata"
const rootPkg = "github.com/wesovilabs/goa/aspect/testdata/a"

func TestGetDefinitions(t *testing.T) {
	cases := []struct {
		directory   string
		definitions *Definitions
	}{
		{
			directory: "a",
			definitions: &Definitions{
				items: []*Definition{
					{
						pkg:    "github.com/wesovilabs/goa/aspect/testdata/a/a2",
						name:   "AroundA2",
						kind:   around,
						regExp: internal.NormalizeExpression("*.*(*)..."),
					},
					{
						pkg:    "github.com/wesovilabs/goa/aspect/testdata/a/a1/a11",
						name:   "AroundA11",
						kind:   around,
						regExp: internal.NormalizeExpression("*.*(*)..."),
					},
					{
						pkg:    "github.com/wesovilabs/goa/aspect/testdata/a",
						name:   "NewTracingReturning",
						kind:   returning,
						regExp: internal.NormalizeExpression("*.*(*)..."),
					},
					{
						pkg:    "github.com/wesovilabs/goa/aspect/testdata/a",
						name:   "NewTracingBefore",
						kind:   before,
						regExp: internal.NormalizeExpression("*.*(*)..."),
					},
					{
						pkg:    "github.com/wesovilabs/goa/aspect/testdata/a",
						name:   "Around",
						kind:   around,
						regExp: internal.NormalizeExpression("*.*(*)..."),
					}, {
						pkg:    "github.com/wesovilabs/goa/aspect/testdata/a/a1",
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
			New(fmt.Sprintf("%s/%s", testdataPath, c.directory), "testdata", false).
			Parse("testdata", "")
		assert.NotNil(packages)
		defs := GetDefinitions(rootPkg, packages)
		assert.Len(c.definitions.items, len(defs.items))
		for index, definition := range c.definitions.List() {
			if !(assert.Equal(definition.pkg, defs.items[index].Pkg()) &&
				assert.Equal(definition.kind, defs.items[index].kind) &&
				assert.Equal(definition.regExp.String(), defs.items[index].regExp.String()) &&
				assert.Equal(definition.name, defs.items[index].Name())) {
				assert.FailNow("error")
			}
			if definition.kind == around || definition.kind == before {
				assert.True(definition.HasBefore())
			}
			if definition.kind == around || definition.kind == returning {
				assert.True(definition.HasReturning())
			}
		}
	}
}
