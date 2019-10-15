package main

import (
	"fmt"
	"github.com/wesovilabs/goa/goa"
	"go/parser"
	"go/token"
	"os"
)

const rootDir = ".goa"

func main() {
	showBanner()
	fmt.Println(">  ....")
	os.Mkdir(rootDir, os.ModePerm)
	fileSet := token.NewFileSet()

	file, err := parser.ParseFile(fileSet, os.Args[1], nil, parser.ParseComments)
	if err != nil {
		fmt.Errorf("error %w", err)
	}
	goa.Goa().Execute(file)
	fmt.Println("> Code was generated successfully")
}

func showBanner() {
	fmt.Println(goa.Banner)
}
