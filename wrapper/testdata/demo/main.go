package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main(){
	file, _ := parser.ParseFile(&token.FileSet{}, "/Users/ivan/Workspace/Wesovilabs/goa/testdata/generated/main.go", nil, parser.ParseComments)
	fmt.Println(file)
}

