package internal

import "github.com/wesovilabs/goa/parser"

var packages = parser.
	New("testdata", "testdata", false).
	Parse("testdata", "")
