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
	CheckError(err)

	defer closeFile(f)

	fileSet := token.NewFileSet()
	cfg := printer.Config{
		Mode:     printer.UseSpaces,
		Indent:   0,
		Tabwidth: 8,
	}

	logger.Infof("[generated] %s", path)

	errPrint := cfg.Fprint(f, fileSet, node)
	CheckError(errPrint)
}
