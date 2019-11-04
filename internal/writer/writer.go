package writer

import (
	"github.com/wesovilabs/goa/logger"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

// SaveNode persists a node into a file in the provided path
func SaveNode(node ast.Node, path string) error {
	f, err := os.Create(path)
	if err != nil {
		logger.Errorf("Errorf while creating file: '%v'", err)
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			logger.Errorf("error while closing file: '%v'", err)
		}
	}()

	fileSet := token.NewFileSet()
	cfg := printer.Config{
		Mode:     printer.UseSpaces,
		Indent:   0,
		Tabwidth: 8,
	}

	logger.Infof("[output] %s", path)

	return cfg.Fprint(f, fileSet, node)
}
