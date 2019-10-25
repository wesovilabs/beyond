package aspect

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/aspect/internal"
	"go/parser"
	"go/token"
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
						pkgPath: "github.com/wesovilabs/goa/aspect/testdata/a/a2",
						name:    "AroundA2",
						kind:    around,
						regExp:  internal.NormalizeExpression("*.*(*)..."),
					},
					{
						pkgPath: "github.com/wesovilabs/goa/aspect/testdata/a/a1/a11",
						name:    "AroundA11",
						kind:    around,
						regExp:  internal.NormalizeExpression("*.*(*)..."),
					},
					{
						pkgPath: "github.com/wesovilabs/goa/aspect/testdata/a",
						name:    "NewTracingReturning",
						kind:    returning,
						regExp:  internal.NormalizeExpression("*.*(*)..."),
					},
					{
						pkgPath: "github.com/wesovilabs/goa/aspect/testdata/a",
						name:    "NewTracingBefore",
						kind:    before,
						regExp:  internal.NormalizeExpression("*.*(*)..."),
					},
					{
						pkgPath: "github.com/wesovilabs/goa/aspect/testdata/a",
						name:    "Around",
						kind:    around,
						regExp:  internal.NormalizeExpression("*.*(*)..."),
					}, {
						pkgPath: "github.com/wesovilabs/goa/aspect/testdata/a/a1",
						name:    "AroundA1",
						kind:    around,
						regExp:  internal.NormalizeExpression("*.*(*)..."),
					},
				},
			},
		},
	}
	assert := assert.New(t)
	for _, c := range cases {
		packages, err := parser.ParseDir(&token.FileSet{}, fmt.Sprintf("%s/%s", testdataPath, c.directory), nil, parser.ParseComments)
		assert.Nil(err)
		assert.NotNil(packages)
		defs := GetDefinitions(rootPkg, packages)
		assert.Len(c.definitions.items, len(defs.items))
		for index, definition := range c.definitions.List() {
			if !(assert.Equal(definition.pkgPath, defs.items[index].Pkg()) &&
				assert.Equal(definition.kind, defs.items[index].kind) &&
				assert.Equal(definition.regExp.String(), defs.items[index].regExp.String()) &&
				assert.Equal(definition.name, defs.items[index].Name())) {
				assert.FailNow("error")
			}
		}
	}
}
