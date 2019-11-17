package parser

import (
	"fmt"
	"github.com/wesovilabs/goa/logger"
	"go/ast"
	"strings"
)

// GoaParser struct
type GoaParser struct {
	project string
	goPath  goPath
}

// New instance a new GoaParser
func New(gopath string, project string) *GoaParser {
	return &GoaParser{
		project: project,
		goPath:  goPath(gopath),
	}
}

func (pp *GoaParser) goPaths() []goPath {
	return []goPath{pp.goPath}
}

// Package struct
type Package struct {
	node *ast.Package
	path string
}

// Node return the node instance
func (p *Package) Node() *ast.Package {
	return p.node
}

// Path return the package path
func (p *Package) Path() string {
	return p.path
}

// Parse parse the input
func (pp *GoaParser) Parse(path string) map[string]*Package {
	pendingPaths := []string{path}
	excludePaths := map[string]string{}
	packages := make(map[string]*Package)

	for {
		if len(pendingPaths) == 0 {
			return packages
		}

		path := pendingPaths[0]
		path = strings.TrimPrefix(path, pp.project)

		if len(path) > 0 && path[0] == '/' {
			path = path[1:]
		}

		pendingPaths = pendingPaths[1:]

		if _, ok := excludePaths[path]; !ok {
			excludePaths[path] = path

			for _, gp := range pp.goPaths() {
				absPath := gp.AbsPath(path)
				pkg, pkgImports := NewGoaPackage(absPath)

				if pkg == nil {
					continue
				}

				for _, pkg := range pkgImports {
					if _, ok := excludePaths[pkg]; !ok {
						pendingPaths = append(pendingPaths, pkg)
					}
				}

				logger.Infof("[path] %s", fmt.Sprintf("%s/%s", pp.project, path))

				packages[path] = &Package{
					node: pkg,
					path: path,
				}
			}
		}
	}
}
