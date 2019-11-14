package model

import (
	"fmt"
	"github.com/wesovilabs/goa/testdata/advice"
)

type Person struct {
	ID        string
	FirstName string
	LastName  string
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func (p *Person) Apply([]advice.Attribute)func(string,int){
	return nil
}
