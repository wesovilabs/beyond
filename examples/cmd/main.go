//go:generate go run github.com/wesovilabs/goa --goPath ../ --project github.com/wesovilabs/goa/examples --verbose
package main

import (
	"fmt"
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/examples/model"
	"github.com/wesovilabs/goa/examples/storage"
)

func main() {
	storage.SetUpDatabase()
	fmt.Println("-----------------------------------------------")
	johnDoe := &model.Person{
		FirstName: "John",
		LastName:  "Doe",
	}
	if err := storage.InsertPerson(johnDoe); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("John Doe with uid %s is created\n", johnDoe.ID)
	fmt.Println("-----------------------------------------------")
	janeDoe := &model.Person{
		FirstName: "Jane",
		LastName:  "Doe",
	}
	if err := storage.InsertPerson(janeDoe); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Jane Doe with uid %s is created\n", janeDoe.ID)
	fmt.Println("-----------------------------------------------")
	people, err := storage.ListPeople()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, person := range people {
			fmt.Printf("%s\n", person.FullName())
		}
	}
	fmt.Println("-----------------------------------------------")
}

func Goa() *api.Goa {
	return api.New().
		WithBefore(api.NewTracingAdvice, `*.*Person(...)...`)
}
