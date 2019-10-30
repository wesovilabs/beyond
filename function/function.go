package function

import (
	"fmt"
	"go/ast"
	"go/token"
)

// Functions struct
type Functions struct {
	functions []*Function
}

// List return the list of functions
func (f *Functions) List() []*Function {
	return f.functions
}

// AddFunction add a new function
func (f *Functions) AddFunction(function *Function) {
	f.functions = append(f.functions, function)
}

// Function struct with required info to efine a function
type Function struct {
	path   string
	decl   *ast.FuncDecl
	parent *ast.File
	pkg    string
}

// Name returns the function name
func (f *Function) Name() string {
	return f.decl.Name.Name
}

// RenameToInternal update the function name
func (f *Function) RenameToInternal() {
	f.decl.Name = ast.NewIdent(fmt.Sprintf("%sInternal", f.decl.Name))
}

// Pkg returns the package name
func (f *Function) Pkg() string {
	return f.pkg
}

// Path return the expression path
func (f *Function) Path() string {
	return f.path
}

// Parent return the parent node
func (f *Function) Parent() *ast.File {
	return f.parent
}

// ImportSpecs returns the list of imports
func (f *Function) ImportSpecs() []*ast.ImportSpec {
	return f.parent.Imports
}

// AddImportSpec add imports to the file
func (f *Function) AddImportSpec(importSpec *ast.ImportSpec) {
	f.AddImportSpecs([]ast.Spec{importSpec})

}

// AddImportSpecs adds decls at the top of the parent
func (f *Function) AddImportSpecs(decls []ast.Spec) {
	f.parent.Decls = append([]ast.Decl{&ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: decls,
	}}, f.parent.Decls...)

	for _, d := range decls {
		f.parent.Imports = append(f.parent.Imports, d.(*ast.ImportSpec))
	}

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

// ResultsList return the list of params that are returned by the function
func (f *Function) ResultsList() []*ast.Field {
	if f.decl.Type.Results != nil {
		return f.decl.Type.Results.List
	}
	return []*ast.Field{}
}
