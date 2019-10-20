package imports

type Imports struct {
	imports []*Import
}

type Import struct {
	name string
	path string
}
func (i *Imports) List() []*Import{
	return i.imports
}

func (i *Imports) WithImport(_import *Import) {
	i.imports = append(i.imports,_import)
}
