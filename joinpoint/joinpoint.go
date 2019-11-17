package joinpoint

import (
	"fmt"
	"go/ast"
	"go/token"
)

// JoinPoints struct
type JoinPoints struct {
	joinPoints []*JoinPoint
}

// List return the list of joinPoints
func (f *JoinPoints) List() []*JoinPoint {
	return f.joinPoints
}

// AddJoinPoint add a new joinPoint
func (f *JoinPoints) AddJoinPoint(joinPoint *JoinPoint) {
	f.joinPoints = append(f.joinPoints, joinPoint)
}

// JoinPoint struct with required info to efine a joinPoint
type JoinPoint struct {
	path    string
	decl    *ast.FuncDecl
	parent  *ast.File
	pkg     string
	pkgPath string
}

// GetRecv return the fieldList
func (f *JoinPoint) GetRecv() *ast.FieldList {
	return f.decl.Recv
}

// Name returns the joinPoint name
func (f *JoinPoint) Name() string {
	return f.decl.Name.Name
}

// RenameToInternal update the joinPoint name
func (f *JoinPoint) RenameToInternal() {
	f.decl.Name = ast.NewIdent(fmt.Sprintf("%sInternal", f.decl.Name))
}

// Pkg returns the package name
func (f *JoinPoint) Pkg() string {
	return f.pkg
}

// Path return the expression path
func (f *JoinPoint) Path() string {
	return f.path
}

// PkgPath return the pakcage path
func (f *JoinPoint) PkgPath() string {
	return f.pkgPath
}

// Parent return the parent node
func (f *JoinPoint) Parent() *ast.File {
	return f.parent
}

// ImportSpecs returns the list of imports
func (f *JoinPoint) ImportSpecs() []*ast.ImportSpec {
	return f.parent.Imports
}

// AddImportSpec add imports to the file
func (f *JoinPoint) AddImportSpec(importSpec *ast.ImportSpec) {
	f.AddImportSpecs([]ast.Spec{importSpec})
}

// AddImportSpecs adds decls at the top of the parent
func (f *JoinPoint) AddImportSpecs(decls []ast.Spec) {
	f.parent.Decls = append([]ast.Decl{&ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: decls,
	}}, f.parent.Decls...)

	for _, d := range decls {
		f.parent.Imports = append(f.parent.Imports, d.(*ast.ImportSpec))
	}
}

// FileDecls return the lis tof ast.decl of the file
func (f *JoinPoint) FileDecls() []ast.Decl {
	return f.parent.Decls
}

// AddStatementsAtBegin add a list of satements to the joinPoint
func (f *JoinPoint) AddStatementsAtBegin(statements []ast.Stmt) {
	f.decl.Body.List = append(statements, f.decl.Body.List...)
}

// ParamsList return the list of params that belong to the joinPoint
func (f *JoinPoint) ParamsList() []*ast.Field {
	return f.decl.Type.Params.List
}

// ResultsList return the list of params that are returned by the joinPoint
func (f *JoinPoint) ResultsList() []*ast.Field {
	if f.decl.Type.Results != nil {
		return f.decl.Type.Results.List
	}

	return []*ast.Field{}
}
