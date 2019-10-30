package matcher

import (
	"fmt"
	"github.com/wesovilabs/goa/aspect"
	"github.com/wesovilabs/goa/function"
)

// Matches struct
type Matches []*Match

// Match struct
type Match struct {
	Function    *function.Function
	Definitions map[string]*aspect.Definition
}

// FindMatches return the list of existing matches
func FindMatches(functions *function.Functions, definitions *aspect.Definitions) Matches {
	matches := Matches{}
	for _, f := range functions.List() {
		aspects := make(map[string]*aspect.Definition)
		for index, d := range definitions.List() {
			if d.Match(f.Path()) {
				aspects[fmt.Sprintf("aspect%v", index)] = d
			}
		}
		if len(aspects) > 0 {
			matches = append(matches, &Match{
				Function:    f,
				Definitions: aspects,
			})
		}
	}
	return matches
}
