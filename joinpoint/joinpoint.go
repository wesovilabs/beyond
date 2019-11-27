package joinpoint

import (
	"fmt"
	"github.com/wesovilabs/goa/advice"
	"go/ast"
	"go/token"
	"regexp"
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
	advices map[string]*advice.Advice
}

// GetRecv return the fieldList
func (jp *JoinPoint) GetRecv() *ast.FieldList {
	return jp.decl.Recv
}

// Name returns the joinPoint name
func (jp *JoinPoint) Name() string {
	return jp.decl.Name.Name
}

// Advices returns the list of advices to be applied
func (jp *JoinPoint) Advices() map[string]*advice.Advice {
	return jp.advices
}

// RenameToInternal update the joinPoint name
func (jp *JoinPoint) RenameToInternal() {
	jp.decl.Name = ast.NewIdent(fmt.Sprintf("%sInternal", jp.decl.Name))
}

// Pkg returns the package name
func (jp *JoinPoint) Pkg() string {
	return jp.pkg
}

// Path return the expression path
func (jp *JoinPoint) Path() string {
	return jp.path
}

// PkgPath return the pakcage path
func (jp *JoinPoint) PkgPath() string {
	return jp.pkgPath
}

// Parent return the parent node
func (jp *JoinPoint) Parent() *ast.File {
	return jp.parent
}

// ImportSpecs returns the list of imports
func (jp *JoinPoint) ImportSpecs() []*ast.ImportSpec {
	return jp.parent.Imports
}

// AddImportSpec add imports to the file
func (jp *JoinPoint) AddImportSpec(importSpec *ast.ImportSpec) {
	jp.AddImportSpecs([]ast.Spec{importSpec})
}

// AddImportSpecs adds decls at the top of the parent
func (jp *JoinPoint) AddImportSpecs(decls []ast.Spec) {
	jp.parent.Decls = append([]ast.Decl{&ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: decls,
	}}, jp.parent.Decls...)

	for _, d := range decls {
		i := d.(*ast.ImportSpec)

		if i.Path.Value != `""` {
			jp.parent.Imports = append(jp.parent.Imports, i)
		}
	}
}

// FileDecls return the lis tof ast.decl of the file
func (jp *JoinPoint) FileDecls() []ast.Decl {
	return jp.parent.Decls
}

// AddStatementsAtBegin add a list of satements to the joinPoint
func (jp *JoinPoint) AddStatementsAtBegin(statements []ast.Stmt) {
	jp.decl.Body.List = append(statements, jp.decl.Body.List...)
}

// ParamsList return the list of params that belong to the joinPoint
func (jp *JoinPoint) ParamsList() []*ast.Field {
	return jp.decl.Type.Params.List
}

// ResultsList return the list of params that are returned by the joinPoint
func (jp *JoinPoint) ResultsList() []*ast.Field {
	if jp.decl.Type.Results != nil {
		return jp.decl.Type.Results.List
	}

	return []*ast.Field{}
}

func (jp *JoinPoint) findMatches(advices *advice.Advices) {
	matches := make(map[string]*advice.Advice)
	for index, adv := range advices.List() {
		if adv.Pkg() == jp.PkgPath() && adv.Name() == jp.Name() {
			jp.advices = make(map[string]*advice.Advice)
			return
		}
		if adv.Match(jp.Path()) {
			matches[fmt.Sprintf("advice%v", index)] = adv
		}
	}
	jp.advices = matches
}

func (jp *JoinPoint) canBeIntercepted(ignoredPaths []*regexp.Regexp) bool {
	if jp.pkg == "main" && (jp.Name() == "main" || jp.Name() == "Goa") {
		return false
	}
	for _, ignorePath := range ignoredPaths {

		if ignorePath.MatchString(jp.Path()) {
			return false
		}
	}

	return true
}
