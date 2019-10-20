package function

import (
	"go/ast"
)

type Functions struct {
	functions []*Function
}

func (f *Functions) List() []*Function {
	return f.functions
}

func (a *Functions) withFunction(function *Function) {
	a.functions = append(a.functions, function)
}

// Function struct with required info to efine a function
type Function struct {
	path   string
	decl   *ast.FuncDecl
	parent *ast.File
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

// Path return the expression path
func (f *Function) Parent() *ast.File {
	return f.parent
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
