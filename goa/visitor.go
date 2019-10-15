package goa

type Visitor struct {
}

/**
func (v Visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}
	switch n := node.(type) {
	case *ast.File:
		inspector := inspector.NewInspector(n)
		functions := inspector.SearchFunctions()
		if functions != nil {
			Goa().WithFunctions(functions)
		}
	}
	return v
}
**/
