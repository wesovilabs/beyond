package parser

import (
	"go/ast"
	"strings"
)

type pkgFilter func(k string) bool

var excludeTestPackages = func(k string) bool {
	return !strings.HasSuffix(k, "_test")
}

// it returns the first package that match
func applyPkgFilters(packages map[string]*ast.Package, filters ...pkgFilter) *ast.Package {
	for name, pkg := range packages {
		for _, filter := range filters {
			if filter(name) {
				return pkg
			}
		}
	}

	return nil
}
