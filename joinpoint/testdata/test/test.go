package test

import "github.com/wesovilabs/goa/joinpoint/testdata/test/model"

func a(name, firstName string) {

}

func b(string) {

}

func c(string, int) {

}

func d(name, title string, age int) {

}

func e(params ...string) {

}

func f(string, ...string) {

}

func g(...*string) {

}

func h(...func(...string)){

}

func i(...*model.TestArgument){

}

type Element struct {

}

func (e *Element) j(string)map[string]interface{}{
	return nil
}

func (e Element) k(){
}



