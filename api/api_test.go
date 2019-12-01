package api

import "testing"

func Test_Api(t *testing.T) {
	returning := func() Returning { return nil }
	around := func() Around { return nil }
	before := func() Before { return nil }
	New().WithReturning(returning, "*.*").
		WithAround(around, "").
		WithBefore(before, "").
		Exclude(".")
}
