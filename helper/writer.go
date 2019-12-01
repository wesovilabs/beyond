package helper

import (
	"github.com/wesovilabs/beyond/logger"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

// Save persists a node into a file in the provided path
func Save(node *ast.File, path string) {
	f, err := os.Create(path)
	checkError(err)

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

	logger.Infof("[generated] %s", path)

	errPrint:=cfg.Fprint(f, fileSet, node)
	checkError(errPrint)
}
