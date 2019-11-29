//go:generate go run github.com/wesovilabs/beyond --goPath ../ --project github.com/wesovilabs/beyond/testdata --verbose
package main

import (
	"fmt"
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond/api/advice"
	"github.com/wesovilabs/beyond/api/context"
	testAdvice "github.com/wesovilabs/beyond/testdata/advice"
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
		WithBefore(testAdvice.NewComplexBefore(&testAdvice.Attribute{}), `*.*Person(...)...`).
		WithBefore(advice.NewTracingAdvice, `*.*Person(...)...`).
		WithAround(testAdvice.NewEmptyAround, `*.*(...)...`).
		// since it's a non public function, they won't be used
		WithReturning(newEmptyReturning,`*.*(...)...`).
		WithReturning(newEmptyReturning,`*.*(...)...`).
		WithAround(testAdvice.NewComplexAround("test",testAdvice.Attribute{},nil),`*.*(...)...`)
}

type EmptyReturning struct{

}

func (r *EmptyReturning) Returning(ctx *context.BeyondContext){

}

func newEmptyReturning() api.Returning{
	return &EmptyReturning{}
}
