package inspector

import (
	"github.com/wesovilabs/goa/inspector/aspect"
	"github.com/wesovilabs/goa/logger"
	"go/ast"
)

// AspectInspector struct
type AspectInspector struct {
	Node *ast.FuncDecl
}

// TakeAspects returns the aspect
func (i *AspectInspector) TakeAspects(pkg string) []*aspect.Aspect {
	output := make([]*aspect.Aspect, 0)
	for _, stmt := range i.Node.Body.List {
		if returnStmt, ok := stmt.(*ast.ReturnStmt); ok {
			if callExpr, ok := returnStmt.Results[0].(*ast.CallExpr); ok {
				aspects := make([]*aspect.Aspect, 0)
				aspects = i.takeAspects(pkg, callExpr, aspects)
				output = append(output, aspects...)
			}
		}
	}

	return output
}

// takeAspectFromCallExpr function used to find aspects by processing the ast
func takeAspectFromCallExpr(pkg string, expr *ast.CallExpr) *aspect.Aspect {
	found := false
	aspect := &aspect.Aspect{}
	for _, a := range expr.Args {
		switch e := a.(type) {
		case *ast.BasicLit:

			if _, err := aspect.WithPattern(e.Value[1 : len(e.Value)-1]); err == nil {
				found = true
				continue
			}
			logger.Errorf("Invalid expression %s\n", e.Value)
			return nil
		case *ast.Ident:
			aspect.WithName(e.Obj.Name).WithPkg(pkg)
		case *ast.SelectorExpr:
			aspect.WithName(e.Sel.Name)
			if i, ok := e.X.(*ast.Ident); ok {
				aspect.WithPkg(i.Name)
			}
		}
	}
	if found {
		logger.Infof("Registering aspect %s.%s", aspect.Pkg(), aspect.Name())
		return aspect
	}
	return nil
}

func (i *AspectInspector) takeAspects(pkgName string, expr *ast.CallExpr, aspects []*aspect.Aspect) []*aspect.Aspect {
	aspect := takeAspectFromCallExpr(pkgName, expr)
	if aspect != nil {
		aspects = append(aspects, aspect)
	}
	if expr.Fun != nil {
		if selectorExpr, ok := expr.Fun.(*ast.SelectorExpr); ok {
			if callExpr, ok := selectorExpr.X.(*ast.CallExpr); ok {
				aspects = i.takeAspects(pkgName, callExpr, aspects)
			}
		}
	}
	return aspects
}
