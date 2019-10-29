package internal

import "go/ast"

func NewIdent(name string) *ast.Ident {
	return ast.NewIdent(name)
}

func NewIdentObj(name string) *ast.Ident {
	ident := NewIdent(name)
	ident.Obj = &ast.Object{
		Name: name,
	}
	return ident
}
func NewIdentObjVar(name string) *ast.Ident {
	ident := NewIdentObj(name)
	ident.Obj.Kind = ast.Var
	return ident
}
