package parser

import "path/filepath"

type goPath string

func (gp *goPath) string() string {
	return string(*gp)
}

func (gp *goPath) AbsPath(path string) string {
	return filepath.Join(gp.string(), path)
}
