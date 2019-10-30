package internal

import "go/ast"

// NewIdent create an ident
func NewIdent(name string) *ast.Ident {
	return ast.NewIdent(name)
}

// NewIdentObj create an ident
func NewIdentObj(name string) *ast.Ident {
	ident := NewIdent(name)
	ident.Obj = &ast.Object{
		Name: name,
	}
	return ident
}

// NewIdentObjVar create an ident
func NewIdentObjVar(name string) *ast.Ident {
	ident := NewIdentObj(name)
	ident.Obj.Kind = ast.Var
	return ident
}
