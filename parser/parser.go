package parser

import (
	"errors"
	"github.com/wesovilabs/goa/logger"
	"go/ast"
	"path/filepath"
	"strings"
)

var ErrTooManyPackages = errors.New("more than one package found in a directory")

const vendorDir = "vendor"

// goPath is project root path
type GoaParser struct {
	project string
	path    goPath
	vendor  bool
	//packages map[string]*internal.GoaPackage
}

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

func (pp *GoaParser) Parse(rootPath string) map[string]*ast.Package {
	pendingPaths := []string{rootPath}
	excludePaths := map[string]string{}
	packages := make(map[string]*ast.Package)
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
				logger.Infof("[path] %s", path)
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
				packages[path] = pkg
			}
		}

	}
}
