package model

import (
	"fmt"
	testAdvice "github.com/wesovilabs/beyond/testdata/advice"
)

type Person struct {
	ID        string
	FirstName string
	LastName  string
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func (p *Person) Apply([]testAdvice.Attribute)func(string,int){
	return nil
}

func (p *Person) Other()(name string){
	name="hey"
	return
}
