package imports

import (
	"go/ast"
	"strings"
)

// GetImports return the list of imports
func GetImports(file *ast.File) map[string]string {
	imports := make(map[string]string)
	for _, im := range file.Imports {
		value := im.Path.Value[1 : len(im.Path.Value)-1]
		if im.Name != nil {
			imports[value] = im.Name.Name
			continue
		}
		parts := strings.Split(value, "/")
		imports[value] = parts[len(parts)-1]
	}
	return imports
}
