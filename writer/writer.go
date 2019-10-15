package writer

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

// Node persists a node into a file in the provided path
func Node(node ast.Node, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	fset := token.NewFileSet()
	cfg := printer.Config{
		Mode:     printer.UseSpaces,
		Indent:   0,
		Tabwidth: 8,
	}
	return cfg.Fprint(f, fset, node)
}
