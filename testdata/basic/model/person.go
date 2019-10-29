package model

type Person struct {
	firstName string
	age       int
}

func (p *Person) FirstName() string{
	return p.firstName
}

func (p *Person) Age() int{
	return p.age
}

func NewPerson(firstName string, age int) *Person {
	return &Person{
		firstName: firstName,
		age:       age,
	}
}

