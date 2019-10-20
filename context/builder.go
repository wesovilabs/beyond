package context

type builder struct {
	pkgName  string
	funcName string
	input    Input
	output   Output
}

var defaultBuilder = builder{
	pkgName: "",
	input:   Input{},
	output:  Output{},
}

// Builder builder type
type Builder func(*builder)
