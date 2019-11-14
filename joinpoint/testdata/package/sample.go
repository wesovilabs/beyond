package _package

import "github.com/wesovilabs/goa/joinpoint/testdata/other"

func sample(_ Person) string {
	return "hey there!"
}


func sample2(other *other.Other) func(map[string]interface{}) {
	return func(map[string]interface{}){

	}
}
