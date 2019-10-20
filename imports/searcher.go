package imports

import (
	"go/ast"
	"strings"
)

func GetImports(file *ast.File) map[string]string {
	imports := make(map[string]string)
	for _, im := range file.Imports {
		if im.Name != nil {
			imports[im.Path.Value] = im.Name.Name
			continue
		}
		parts := strings.Split(im.Path.Value[1:len(im.Path.Value)-1], "/")
		imports[im.Path.Value[1:len(im.Path.Value)-1]] = parts[len(parts)-1]
	}
	return imports
}
