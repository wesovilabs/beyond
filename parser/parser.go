package parser

import (
	"fmt"
	"github.com/wesovilabs/goa/logger"
	"go/ast"
	"path/filepath"
	"strings"
)

const vendorDir = "vendor"

// GoaParser struct
type GoaParser struct {
	project string
	path    goPath
	vendor  bool
}

// New instance a new GoaParser
func New(path string, project string, vendor bool) *GoaParser {
	return &GoaParser{
		project: project,
		path:    goPath(path),
		vendor:  vendor,
	}
}

func (pp *GoaParser) goPaths() []goPath {
	if pp.vendor {
		vendorPath := filepath.Join(string(pp.path), vendorDir)
		return []goPath{pp.path, goPath(vendorPath)}
	}
	return []goPath{pp.path}
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
func (pp *GoaParser) Parse(project, rootPath string) map[string]*Package {
	pendingPaths := []string{rootPath}
	excludePaths := map[string]string{}
	packages := make(map[string]*Package)
	for {
		if len(pendingPaths) == 0 {
			return packages
		}
		path := pendingPaths[0]
		path = strings.TrimPrefix(path, pp.project)
		pendingPaths = pendingPaths[1:]
		if _, ok := excludePaths[path]; !ok {
			excludePaths[path] = path
			for _, gp := range pp.goPaths() {
				absPath := gp.AbsPath(path)
				logger.Infof("[path] %s", absPath)
				pkg, pkgImports := NewGoaPackage(absPath)
				if pkg == nil {
					continue
				}
				for _, pkg := range pkgImports {
					if _, ok := excludePaths[pkg]; !ok {
						pendingPaths = append(pendingPaths, pkg)
					}
				}
				packages[path] = &Package{
					node: pkg,
					path: fmt.Sprintf("%s%s", project, path),
				}
			}
		}

	}
}
