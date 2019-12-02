package main

import (
	"fmt"
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond/api/context"
	"github.com/wesovilabs/beyond/testdata/advice"
	"github.com/wesovilabs/beyond/testdata/model"
	"github.com/wesovilabs/beyond/testdata/storage"
)

func main() {
	storage.SetUpDatabase()
	fmt.Println("-----------------------------------------------")
	johnDoe := &model.Person{
		FirstName: "John",
		LastName:  "Doe",
	}
	if err := storage.InsertPerson(johnDoe,nil); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("John Doe with uid %s is created\n", johnDoe.ID)
	fmt.Println("-----------------------------------------------")
	janeDoe := &model.Person{
		FirstName: "Jane",
		LastName:  "Doe",
	}
	if err := storage.InsertPerson(janeDoe,nil); err != nil {
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

func Beyond() *api.Beyond {
	return api.New().
		WithBefore(advice.NewComplexBefore(&advice.Attribute{}), `*.*Person(...)...`).
		WithBefore(advice.NewTracingAdvice, `*.*Person(...)...`).
		WithAround(advice.NewEmptyAround, `*.*(...)...`).
		WithReturning(newEmptyReturning,`*.*(...)...`).
		WithReturning(newEmptyReturning,`*.*(...)...`).
		WithAround(advice.NewComplexAround("test",advice.Attribute{},nil, struct{}{},""),`*.*(...)...`).
		WithReturning(newEmptyReturning,`*.*.(...)...`).
		WithAround(advice.NewEmptyAround,`*.*Person.*(...)...`).
		WithAround(advice.NewEmptyAround,`*.*Person.*(struct{},...string)...`).
		Ignore("*.*Person.Other(...)...")


}

type EmptyReturning struct{

}

func (r *EmptyReturning) Returning(ctx *context.BeyondContext){

}

func newEmptyReturning() api.Returning{
	return &EmptyReturning{}
}
