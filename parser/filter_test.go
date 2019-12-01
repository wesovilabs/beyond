package parser

import "testing"

func Test_applyPkgFilters(t *testing.T) {
	if res := applyPkgFilters(nil); res != nil {
		t.Fatal("Expected nil package")
	}
}
