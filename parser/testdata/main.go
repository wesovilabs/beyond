package main

import (
	"fmt"
	"github.com/wesovilabs/goa/parser/testdata/a"
	"github.com/wesovilabs/goa/parser/testdata/b"
)

func main(){
	person:=&a.Person{}
	address:=&b.Address{}
	fmt.Println(address)
	fmt.Println(person)
}
