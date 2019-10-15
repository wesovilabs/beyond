package inspector

/**
if funcDecl.Doc != nil {
	for _, comment := range funcDecl.Doc.List {
		exprStr := strings.TrimLeft(comment.Text, "/")
		exprStr = strings.TrimLeft(exprStr, " ")
		if beeRegexp.MatchString(exprStr) {
			exprStr = beeRegexp.FindStringSubmatch(exprStr)[0]

			exprStr = strings.TrimLeft(exprStr, "goa:")
			exprStr = strings.TrimLeft(exprStr, " ")
			/**
			if expression, err := buildExpression(exprStr); err == nil {
				expressions[path] = []*Expression{expression}
			}
		}
	}
}
**/
