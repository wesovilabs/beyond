package inspector

import (
	"github.com/wesovilabs/goa/inspector/internal"
	"go/ast"
)

// Function struct with required info to efine a function
type Function struct {
	path   string
	decl   *ast.FuncDecl
	parent *ast.File
}

func newFunction(fileDecl *ast.File, funcDecl *ast.FuncDecl) *Function {
	path := internal.BuildPath(fileDecl, funcDecl)
	return &Function{
		parent: fileDecl,
		decl:   funcDecl,
		path:   path,
	}
}

// Name returns the function name
func (f *Function) Name() string {
	return f.decl.Name.Name
}

// Pkg returns the package name
func (f *Function) Pkg() string {
	return f.parent.Name.Name
}

// Path return the expression path
func (f *Function) Path() string {
	return f.path
}

// ImportSpecs returns the list of imports
func (f *Function) ImportSpecs() []*ast.ImportSpec {
	return f.parent.Imports
}

// AddImport add imports to the file
func (f *Function) AddImport(importSpec *ast.ImportSpec) {
	f.parent.Imports = append(f.parent.Imports, importSpec)
}

// AddDeclsBefore adds decls at the top of the parent
func (f *Function) AddDeclsBefore(decls []ast.Decl) {
	f.parent.Decls = append(decls, f.parent.Decls...)
}

// FileDecls return the lis tof ast.decl of the file
func (f *Function) FileDecls() []ast.Decl {
	return f.parent.Decls
}

// AddStatementsAtBegin add a list of satements to the function
func (f *Function) AddStatementsAtBegin(statements []ast.Stmt) {
	f.decl.Body.List = append(statements, f.decl.Body.List...)
}

// ParamsList return the list of params that belong to the function
func (f *Function) ParamsList() []*ast.Field {
	return f.decl.Type.Params.List
}
