package model

import "fmt"

type Person struct {
	ID        string
	FirstName string
	LastName  string
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}
